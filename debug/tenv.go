package debug

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// RuntimeEnv 登录后从服务器环境脚本动态获取的 T100 路径(与标准 debug 同源)。
// T100 无 $TOPDIR 变量:登录区域经站点 profile 的 case 表得到 ZONE 目录名,
// topenv 用 TOP=/u1/$ZONE 派生全部路径,ERP/COM 由 TOP 派生;模块变量由 topsys 扫描生成。
// 动态值优先于静态配置(TopDir/ModuleRoots 仅作探针失败兜底)。
type RuntimeEnv struct {
	TOP             string    // /u1/t35tst
	ERP             string    // $TOP/erp
	COM             string    // $TOP/com
	FGLDIR          string    // /u1/genero/fgl
	FGLResourcePath string    // $ERP:$COM
	FetchedAt       time.Time // 获取时间(缓存 TTL 判定)
}

// valid 动态环境是否可用
func (e *RuntimeEnv) valid() bool { return e != nil && e.TOP != "" && e.ERP != "" }

var reTEnvKV = regexp.MustCompile(`^(TDICT_TOP|TDICT_ERP|TDICT_COM|TDICT_FGLDIR|TDICT_FGLRESOURCEPATH)=(\S*)\s*$`)

// reZone 区域代码白名单(31/35/36/39/t/36k/1 等):允许数字/字母/下划线/连字符,防注入
var reZone = regexp.MustCompile(`^[A-Za-z0-9_-]{1,16}$`)

// tenvProbeScript 拼出「加载 zone 环境并回读关键变量」的 bash 脚本。
// 与登录 profile 链同源:chenv(区域分支)→ topenv(TOP/ERP/COM)→ topsys(模块变量);
// 候选路径覆盖 UAT(/u3/pub)与金仓(/u1)两站点,source 失败静默跳过,最终以 echo 值为准。
// zone 含非法字符时返回空串(防注入)。
func tenvProbeScript(zone string) string {
	if !reZone.MatchString(zone) {
		return ""
	}
	z := strings.ReplaceAll(zone, "'", "'\\''")
	return fmt.Sprintf(`bash -lc '
ZONE=%s; export ZONE
source /u3/pub/bin/chenv %s >/dev/null 2>&1
source /u3/pub/etc/chenv >/dev/null 2>&1
source /u1/etc/chenv >/dev/null 2>&1
source /u3/pub/etc/topenv >/dev/null 2>&1
source /u1/etc/topenv >/dev/null 2>&1
source /u1/etc/topsys >/dev/null 2>&1
echo TDICT_TOP=$TOP
echo TDICT_ERP=$ERP
echo TDICT_COM=$COM
echo TDICT_FGLDIR=$FGLDIR
echo TDICT_FGLRESOURCEPATH=$FGLRESOURCEPATH
'`, z, z)
}

// probeTEnv 在服务器上探测登录区域对应的 T100 环境变量
func probeTEnv(conn *SSHConn, zone string) (*RuntimeEnv, error) {
	script := tenvProbeScript(zone)
	if script == "" {
		return nil, fmt.Errorf("区域代码非法: %q", zone)
	}
	out, err := conn.Output(script, 30*time.Second)
	if err != nil && out == "" {
		return nil, fmt.Errorf("探测 T100 环境失败: %w", err)
	}
	env := &RuntimeEnv{FetchedAt: time.Now()}
	for _, ln := range strings.Split(out, "\n") {
		if m := reTEnvKV.FindStringSubmatch(strings.TrimSpace(ln)); m != nil {
			switch m[1] {
			case "TDICT_TOP":
				env.TOP = m[2]
			case "TDICT_ERP":
				env.ERP = m[2]
			case "TDICT_COM":
				env.COM = m[2]
			case "TDICT_FGLDIR":
				env.FGLDIR = m[2]
			case "TDICT_FGLRESOURCEPATH":
				env.FGLResourcePath = m[2]
			}
		}
	}
	if !env.valid() {
		return nil, fmt.Errorf("未探测到 T100 环境变量(TOP 为空,区域 %s 环境脚本未加载?)", zone)
	}
	return env, nil
}
