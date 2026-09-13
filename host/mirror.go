package host

// 本地源码镜像引擎:把某环境(T100 服务器)的 4gl/4fd 源码拉到本地镜像目录。
//
// 白名单原则(与调试读码一致):只收各模块 4gl/4fd 两个目录树
//   find erp com -type f \( -path '*/4gl/*' -o -path '*/4fd/*' \)
// per(界面源)/42m/42r/42f/多语言/设计器辅助等一概不拉 —— AI 只读源代码与
// 前端字段描述,本地目录与服务器同构(erp/、com/),AI 用本地文件工具直接读。
//
// 传输:服务器侧 tar 打包(绝对路径,exec 通道不经登录 profile) → SFTP 流式
// 下载 → 本地标准库解压(gzip+tar),全程不进内存。
// 增量:服务器保留 marker 文件(.tdict-mirror-<env>.mark),pull 默认只打包
// find -newer 的变更文件;--full 全量打包并在本地整目录替换(含删除残留)。

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"
)

// MirrorStats 一次镜像拉取的统计结果(CLI 打印用)。
type MirrorStats struct {
	Env     string // 环境名
	TopDir  string // 服务器 TOP(打包根)
	Full    bool   // 是否全量
	Files   int    // 归档内文件数(无变更时为 0)
	Bytes   int64  // 下载字节数(压缩后)
	Elapsed string // 耗时
	Note    string // 附加说明(如 "无变更")
}

// localMarkName 本地镜像目录内的"完整基线"标记文件:存在表示该目录曾成功
// 建立/更新过完整镜像,增量更新以此为前提;缺失时强制全量(见 MirrorPull)。
const localMarkName = ".tdict-mirror.ok"

// mirrorTag 环境名 → 服务器 marker 文件名 ASCII 尾缀。
// 环境名可能是中文/含特殊字符(不能直接作文件名),且同一服务器可挂多个环境,
// 统一取 sha1 前 10 位十六进制:唯一、安全、与语言无关。
func mirrorTag(name string) string {
	sum := sha1.Sum([]byte(name))
	return fmt.Sprintf("%x", sum)[:10]
}

// MirrorProgress 拉取过程的进度(Web 进度条用;CLI 传 nil 回调即无进度)。
type MirrorProgress struct {
	Phase   string // connect|probe|pack|download|done
	Message string // 人类可读说明
	Bytes   int64  // 已处理字节:pack=服务器归档当前大小,download=已读压缩字节
	Total   int64  // 总字节(仅 download 阶段已知;0=未知)
	Files   int    // 文件数(已知时)
}

// ProgressFunc 进度回调(可能被多次调用,实现需自行加锁)。
type ProgressFunc func(MirrorProgress)

// MirrorPull 拉取/更新指定环境(NamedSsh)的 4gl+4fd 源码镜像到 mirrorDir/<环境名>/。
// full=true 全量重建;否则增量(服务器 marker 记录基线)。
// 增量只在本地已有完整基线(envDir 内存在 .tdict-mirror.ok 标记)时允许;
// 本地首次/目录被清/旧版无标记 → 自动转全量,防止"服务器 marker 存在而本地
// 无镜像"时只拉到变更文件造成镜像永久残缺。成功结束后写入/刷新本地基线标记。
// 镜像目录本身(mirrorDir)由调用方保证已配置;本函数只负责连服务器与落盘。
func MirrorPull(e *NamedSsh, mirrorDir string, full bool) (*MirrorStats, error) {
	return MirrorPullProgress(e, mirrorDir, full, nil)
}

// MirrorPullProgress 同 MirrorPull,并在各阶段回调 onProgress(供 Web 显示拉取进度)。
func MirrorPullProgress(e *NamedSsh, mirrorDir string, full bool, onProgress ProgressFunc) (*MirrorStats, error) {
	report := func(p MirrorProgress) {
		if onProgress != nil {
			onProgress(p)
		}
	}
	start := time.Now()
	st := &MirrorStats{Env: e.Name, Full: full}
	if e.Name == "" {
		return nil, fmt.Errorf("环境名为空")
	}

	report(MirrorProgress{Phase: "connect", Message: "连接 SSH " + e.SSHConfig.Addr()})
	conn, err := Dial(e.SSHConfig)
	if err != nil {
		return nil, fmt.Errorf("连接 %s 失败: %w", e.SSHConfig.Addr(), err)
	}
	defer conn.Close()

	report(MirrorProgress{Phase: "probe", Message: "探测 T100 目录(zone " + e.Zone + ")"})
	top, err := mirrorTopDir(conn, e)
	if err != nil {
		return nil, err
	}
	st.TopDir = top

	tag := mirrorTag(e.Name)
	base := top + "/.tdict-mirror-" + tag

	// 本地无完整基线(首次拉取/目录被清/旧版本遗留)时,增量没有可叠加的对象,
	// 自动按全量拉取,避免只拿到服务器 marker 之后的变更文件导致镜像残缺
	envDir := filepath.Join(mirrorDir, e.Name)
	needFull := full
	if !needFull {
		if _, err := os.Stat(filepath.Join(envDir, localMarkName)); err != nil {
			needFull = true
			logfMirror("本地镜像缺少完整基线标记(%s),本次自动按全量拉取以保证完整", filepath.Join(envDir, localMarkName))
		}
	}
	st.Full = needFull

	// 服务器打包(输出 TDICT_MIRROR_OK <文件数> <字节> / TDICT_MIRROR_NONE)
	report(MirrorProgress{Phase: "pack", Message: "服务器打包中…"})
	n, size, err := mirrorPackProgress(conn, base, top, needFull, report)
	if err != nil {
		return nil, fmt.Errorf("服务器打包失败: %w", err)
	}
	if n == 0 {
		st.Elapsed = time.Since(start).Round(100 * time.Millisecond).String()
		st.Note = "无变更(服务器 4gl/4fd 自上次 pull 后未修改;--full 可强制全量重建)"
		// 服务器无变更但本地基线标记缺失时,依然补一个基线标记,
		// 否则下次 pull 仍会误判"无基线"而反复尝试全量
		if !needFull {
			if err := writeLocalMark(envDir); err != nil {
				logfMirror("写入本地基线标记失败(忽略): %v", err)
			}
		}
		report(MirrorProgress{Phase: "done", Message: st.Note})
		return st, nil
	}
	st.Bytes = size

	// 下载并解压(流式):countingReader 按压缩字节回报进度
	if err := os.MkdirAll(mirrorDir, 0o755); err != nil {
		return nil, fmt.Errorf("创建镜像根目录失败: %w", err)
	}
	sc, err := conn.SFTP()
	if err != nil {
		return nil, fmt.Errorf("打开 SFTP 失败: %w", err)
	}
	defer sc.Close()
	rf, err := sc.Open(base + ".tar.gz")
	if err != nil {
		return nil, fmt.Errorf("打开服务器归档失败: %w", err)
	}
	defer rf.Close()

	report(MirrorProgress{Phase: "download", Message: "下载并解压…", Total: size})
	cr := &countingReader{r: rf, total: size, report: report}
	var files int
	if needFull {
		files, err = extractFull(cr, envDir)
	} else {
		files, err = extractIncremental(cr, envDir)
	}
	if err != nil {
		return nil, fmt.Errorf("本地解压失败: %w", err)
	}
	st.Files = files
	if err := writeLocalMark(envDir); err != nil {
		return nil, fmt.Errorf("写入本地基线标记失败: %w", err)
	}

	// 清理服务器归档(marker 保留,供下次增量)
	if out, err := conn.Output(`rm -f "`+base+`.tar.gz"`, time.Minute); err != nil && !strings.Contains(out, "No such file") {
		// 清理失败不阻断:下次全量/增量都会重建归档
		logfMirror("清理服务器归档失败(忽略): %v", err)
	}
	st.Elapsed = time.Since(start).Round(100 * time.Millisecond).String()
	report(MirrorProgress{Phase: "done", Message: fmt.Sprintf("完成:%d 个文件", files), Bytes: size, Total: size, Files: files})
	return st, nil
}

// countingReader 统计已读字节并节流回调进度(下载/解压阶段共用;压缩字节即下载量)。
type countingReader struct {
	r      io.Reader
	n      int64
	total  int64
	last   time.Time
	report ProgressFunc
}

func (c *countingReader) Read(p []byte) (int, error) {
	n, err := c.r.Read(p)
	if n > 0 {
		c.n += int64(n)
	}
	now := time.Now()
	// 节流:每 200ms 一次;流结束时(EOF)必报一次最终字节数
	if (n > 0 && now.Sub(c.last) >= 200*time.Millisecond) || err == io.EOF {
		c.last = now
		c.report(MirrorProgress{Phase: "download", Message: "下载并解压…", Bytes: c.n, Total: c.total})
	}
	return n, err
}

// mirrorPackProgress 在服务器打包期间轮询归档大小回报进度(pack 阶段无法预知总量)。
func mirrorPackProgress(conn *SSHConn, base, top string, full bool, report ProgressFunc) (int, int64, error) {
	var stopped atomic.Bool
	done := make(chan struct{})
	defer func() {
		stopped.Store(true)
		close(done)
	}()
	go func() {
		t := time.NewTicker(2 * time.Second)
		defer t.Stop()
		cmd := `stat -c %s ` + shdq(base+".tar.gz") + ` 2>/dev/null || echo 0`
		for {
			select {
			case <-done:
				return
			case <-t.C:
				out, err := conn.Output(cmd, 8*time.Second)
				if err != nil {
					continue
				}
				var sz int64
				if _, err := fmt.Sscanf(strings.TrimSpace(out), "%d", &sz); err == nil && sz > 0 && !stopped.Load() {
					report(MirrorProgress{Phase: "pack", Message: "服务器打包中…", Bytes: sz})
				}
			}
		}
	}()
	return mirrorPack(conn, base, top, full)
}

// writeLocalMark 在镜像环境目录写入完整基线标记(带时间)。
func writeLocalMark(envDir string) error {
	if err := os.MkdirAll(envDir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(envDir, localMarkName),
		[]byte(time.Now().UTC().Format(time.RFC3339)+"\n"), 0o644)
}

// MirrorEnvDir 返回某环境的本地镜像目录(mirrorDir/<环境名>)。
func MirrorEnvDir(mirrorDir, envName string) string {
	if mirrorDir == "" || envName == "" {
		return ""
	}
	return filepath.Join(mirrorDir, envName)
}

// MirrorReady 本地是否已有完整镜像基线(.tdict-mirror.ok 标记;增量更新的前提)。
// 无基线时 MirrorPull 会自动转全量。
func MirrorReady(mirrorDir, envName string) bool {
	dir := MirrorEnvDir(mirrorDir, envName)
	if dir == "" {
		return false
	}
	_, err := os.Stat(filepath.Join(dir, localMarkName))
	return err == nil
}

func logfMirror(format string, args ...any) {
	// 轻量日志:镜像失败清理之类不致命信息走 stderr(无全局 logger 依赖)
	fmt.Fprintf(os.Stderr, "[mirror] "+format+"\n", args...)
}

// mirrorTopDir 确定打包根:仅 SSH 登录探测(Runtime.TOP)——T100 路径无静态配置,
// 一律按登录区域动态获取;探测失败即报错。
func mirrorTopDir(conn *SSHConn, e *NamedSsh) (string, error) {
	if conn == nil || e.Zone == "" {
		return "", fmt.Errorf("环境 %q 缺少登录区域(zone),无法探测 T100 目录", e.Name)
	}
	env, err := ProbeTEnv(conn, e.Zone)
	if err != nil {
		return "", fmt.Errorf("环境 %s(zone %s)登录探测 T100 目录失败: %w", e.Name, e.Zone, err)
	}
	return strings.TrimSuffix(env.TOP, "/"), nil
}

// shdq 把值包成双引号 shell 字面量(路径含特殊字符安全)。
func shdq(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	s = strings.ReplaceAll(s, "$", `\$`)
	s = strings.ReplaceAll(s, "`", "\\`")
	return `"` + s + `"`
}

// mirrorPack 在服务器 TOP 下打包 4gl/4fd 白名单文件。返回 (文件数, 归档字节);
// 无变更时返回 (0, 0) 且不产生归档。marker 保留在 TOP 下供下次增量。
func mirrorPack(conn *SSHConn, base, top string, full bool) (int, int64, error) {
	m := shdq(base + ".mark")
	o := shdq(base + ".tar.gz")
	l := shdq(base + ".list")
	t := shdq(top)
	// 注意:exec 通道不经过登录 profile,但这里只依赖绝对路径与标准工具
	findExpr := `\( -path '*/4gl/*' -o -path '*/4fd/*' \)`
	fullFlag := "0"
	if full {
		fullFlag = "1"
	}
	script := `set -e
cd ` + t + ` || exit 9
R=""
for d in erp com; do
  [ -d "$d" ] && R="$R $d"
done
if [ -z "$R" ]; then
  echo TDICT_MIRROR_NONE
  exit 0
fi
if [ "` + fullFlag + `" = "1" ] || [ ! -f ` + m + ` ]; then
  find $R -type f ` + findExpr + ` > ` + l + `
else
  find $R -type f ` + findExpr + ` -newer ` + m + ` > ` + l + `
fi
if [ -s "` + l + `" ]; then
  n=$(wc -l < "` + l + `")
  tar -czf ` + o + ` -T "` + l + `"
  touch ` + m + `
  s=$(wc -c < "` + o + `")
  rm -f "` + l + `"
  echo "TDICT_MIRROR_OK $n $s"
else
  rm -f "` + l + `"
  echo TDICT_MIRROR_NONE
fi`
	out, err := conn.Output(script, 20*time.Minute)
	if err != nil {
		return 0, 0, fmt.Errorf("%s(输出: %s)", err, strings.TrimSpace(out))
	}
	// 取最后一行标记(前面可能有 tar 警告)
	mark := ""
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "TDICT_MIRROR_OK") || line == "TDICT_MIRROR_NONE" {
			mark = line
		}
	}
	switch {
	case strings.HasPrefix(mark, "TDICT_MIRROR_OK"):
		var n int
		var s int64
		if _, err := fmt.Sscanf(mark, "TDICT_MIRROR_OK %d %d", &n, &s); err != nil {
			return 0, 0, fmt.Errorf("解析打包标记失败: %q", mark)
		}
		return n, s, nil
	case mark == "TDICT_MIRROR_NONE":
		return 0, 0, nil
	default:
		return 0, 0, fmt.Errorf("服务器脚本未输出预期标记(输出: %s)", strings.TrimSpace(out))
	}
}

// extractIncremental 增量解压:直接覆盖写入 <mirrorDir>/<环境名>/(文件级幂等)。
func extractIncremental(r io.Reader, envDir string) (int, error) {
	if err := os.MkdirAll(envDir, 0o755); err != nil {
		return 0, err
	}
	return extractTarGz(r, envDir)
}

// extractFull 全量解压:先写 .staging 临时目录,再整目录替换(旧目录改名 .old 后删除)。
func extractFull(r io.Reader, envDir string) (int, error) {
	ts := time.Now().Format("20060102-150405.000")
	staging := filepath.Join(filepath.Dir(envDir), ".staging-"+ts)
	if err := os.MkdirAll(staging, 0o755); err != nil {
		return 0, err
	}
	n, err := extractTarGz(r, staging)
	if err != nil {
		os.RemoveAll(staging)
		return 0, err
	}
	old := envDir + ".old-" + ts
	if _, err := os.Stat(envDir); err == nil {
		if err := os.Rename(envDir, old); err != nil {
			os.RemoveAll(staging)
			return 0, fmt.Errorf("备份旧镜像目录失败: %w", err)
		}
	}
	if err := os.Rename(staging, envDir); err != nil {
		// 回滚:把旧目录放回去
		if _, serr := os.Stat(old); serr == nil {
			os.Rename(old, envDir)
		}
		os.RemoveAll(staging)
		return 0, fmt.Errorf("启用新镜像目录失败: %w", err)
	}
	if _, err := os.Stat(old); err == nil {
		if err := os.RemoveAll(old); err != nil {
			logfMirror("删除旧镜像备份失败(可手动清理 %s): %v", old, err)
		}
	}
	return n, nil
}

// extractTarGz 从 tar.gz 流解压到 dest(标准库;不做任何 shell)。返回文件数。
// tar 头路径做穿越防护;回写 mtime 便于与服务器比对;文件 0644。
func extractTarGz(r io.Reader, dest string) (int, error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return 0, fmt.Errorf("解压 gzip: %w", err)
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	files := 0
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return files, err
		}
		name := filepath.Clean(filepath.FromSlash(hdr.Name))
		if name == "." || filepath.IsAbs(name) || name == ".." || strings.HasPrefix(name, ".."+string(filepath.Separator)) {
			return files, fmt.Errorf("归档含非法路径: %q", hdr.Name)
		}
		target := filepath.Join(dest, name)
		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0o755); err != nil {
				return files, err
			}
		case tar.TypeReg, tar.TypeRegA:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return files, err
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
			if err != nil {
				return files, err
			}
			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return files, err
			}
			f.Close()
			_ = os.Chtimes(target, hdr.ModTime, hdr.ModTime)
			files++
		default:
			// 符号链接/设备等不落盘(4gl/4fd 白名单内不会出现)
		}
	}
	return files, nil
}
