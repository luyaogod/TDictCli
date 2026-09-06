package debug

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
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

var mirrorTagRe = regexp.MustCompile(`[^A-Za-z0-9]`)

// MirrorPull 拉取/更新指定环境(NamedSsh)的 4gl+4fd 源码镜像到 mirrorDir/<环境名>/。
// full=true 全量重建;否则增量(服务器 marker 记录基线)。镜像目录本身(mirrorDir)
// 由调用方保证已配置;本函数只负责连服务器与落盘。
func MirrorPull(e *NamedSsh, mirrorDir string, full bool) (*MirrorStats, error) {
	start := time.Now()
	st := &MirrorStats{Env: e.Name, Full: full}
	if e.Name == "" {
		return nil, fmt.Errorf("环境名为空")
	}

	conn, err := Dial(e.SSHConfig)
	if err != nil {
		return nil, fmt.Errorf("连接 %s 失败: %w", e.SSHConfig.Addr(), err)
	}
	defer conn.Close()

	top, err := mirrorTopDir(conn, e)
	if err != nil {
		return nil, err
	}
	st.TopDir = top

	tag := mirrorTagRe.ReplaceAllString(e.Name, "")
	if tag == "" {
		tag = "env"
	}
	if len(tag) > 24 {
		tag = tag[:24]
	}
	base := top + "/.tdict-mirror-" + tag

	// 服务器打包(输出 TDICT_MIRROR_OK <文件数> <字节> / TDICT_MIRROR_NONE)
	n, size, err := mirrorPack(conn, base, top, full)
	if err != nil {
		return nil, fmt.Errorf("服务器打包失败: %w", err)
	}
	if n == 0 {
		st.Elapsed = time.Since(start).Round(100 * time.Millisecond).String()
		st.Note = "无变更(服务器 4gl/4fd 自上次 pull 后未修改;--full 可强制全量重建)"
		return st, nil
	}
	st.Bytes = size

	// 下载并解压(流式)
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

	envDir := filepath.Join(mirrorDir, e.Name)
	var files int
	if full {
		files, err = extractFull(rf, envDir)
	} else {
		files, err = extractIncremental(rf, envDir)
	}
	if err != nil {
		return nil, fmt.Errorf("本地解压失败: %w", err)
	}
	st.Files = files

	// 清理服务器归档(marker 保留,供下次增量)
	if out, err := conn.Output(`rm -f "`+base+`.tar.gz"`, time.Minute); err != nil && !strings.Contains(out, "No such file") {
		// 清理失败不阻断:下次全量/增量都会重建归档
		logfMirror("清理服务器归档失败(忽略): %v", err)
	}
	st.Elapsed = time.Since(start).Round(100 * time.Millisecond).String()
	return st, nil
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
	env, err := probeTEnv(conn, e.Zone)
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
