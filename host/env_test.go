package host

import (
	"os"
	"path/filepath"
	"testing"
)

func writeCfg(t *testing.T, body string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatalf("写入临时配置失败: %v", err)
	}
	return p
}

// 新键 hosts:正常解析,activeEnv 生效。
func TestLoadHostsNewKey(t *testing.T) {
	p := writeCfg(t, `{"hosts":{"activeEnv":"b","sshs":[{"name":"a","host":"1.1.1.1"},{"name":"b","host":"2.2.2.2"}]}}`)
	h, err := LoadHosts(p)
	if err != nil {
		t.Fatalf("LoadHosts: %v", err)
	}
	if h.ActiveEnv != "b" {
		t.Fatalf("activeEnv = %q, want b", h.ActiveEnv)
	}
	if len(h.SSHs) != 2 {
		t.Fatalf("sshs = %d, want 2", len(h.SSHs))
	}
	if got := h.ByName(""); got == nil || got.Name != "b" {
		t.Fatalf("ByName(\"\") 应回退 activeEnv b, got %+v", got)
	}
}

// 旧键 debug:向后兼容,activeEnv 缺省取首条,db 正常解析。
func TestLoadHostsLegacyDebugKey(t *testing.T) {
	p := writeCfg(t, `{"debug":{"sshs":[{"name":"legacy","host":"3.3.3.3","db":{"type":"oracle","host":"3.3.3.3","service":"t35prd"}}]}}`)
	h, err := LoadHosts(p)
	if err != nil {
		t.Fatalf("LoadHosts 旧键: %v", err)
	}
	if len(h.SSHs) != 1 || h.SSHs[0].Name != "legacy" {
		t.Fatalf("sshs = %+v", h.SSHs)
	}
	if h.ActiveEnv != "legacy" {
		t.Fatalf("activeEnv 应缺省取首条, got %q", h.ActiveEnv)
	}
	if h.SSHs[0].DB == nil || h.SSHs[0].DB.Type != "oracle" || h.SSHs[0].DB.Svc() != "t35prd" {
		t.Fatalf("db 未正确解析: %+v", h.SSHs[0].DB)
	}
}

// 两个键都存在时,新键 hosts 优先。
func TestLoadHostsNewKeyWins(t *testing.T) {
	p := writeCfg(t, `{"hosts":{"sshs":[{"name":"new"}]},"debug":{"sshs":[{"name":"old"}]}}`)
	h, err := LoadHosts(p)
	if err != nil {
		t.Fatalf("LoadHosts: %v", err)
	}
	if h.SSHs[0].Name != "new" {
		t.Fatalf("应优先 hosts 键, got %q", h.SSHs[0].Name)
	}
}

// 缺少配置节 / 无 sshs 均报错。
func TestLoadHostsErrors(t *testing.T) {
	if _, err := LoadHosts(writeCfg(t, `{"query":{"source":"local"}}`)); err == nil {
		t.Fatal("缺少 hosts/debug 配置节应报错")
	}
	if _, err := LoadHosts(writeCfg(t, `{"hosts":{"sshs":[]}}`)); err == nil {
		t.Fatal("sshs 为空应报错")
	}
}
