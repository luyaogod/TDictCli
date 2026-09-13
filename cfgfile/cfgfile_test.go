package cfgfile

import (
	"os"
	"path/filepath"
	"testing"
)

// 仅有旧键 debug 时:Hosts 将其提升为 hosts 并删除 debug,其余顶层键保留。
func TestHostsMigratesLegacyDebugKey(t *testing.T) {
	root := map[string]any{
		"debug": map[string]any{"activeEnv": "e1"},
		"query": map[string]any{"source": "local"},
	}
	h, err := Hosts(root)
	if err != nil {
		t.Fatalf("Hosts: %v", err)
	}
	if h["activeEnv"] != "e1" {
		t.Fatalf("activeEnv = %v, want e1", h["activeEnv"])
	}
	if _, ok := root["debug"]; ok {
		t.Fatal("旧 debug 键应被删除")
	}
	if _, ok := root["hosts"]; !ok {
		t.Fatal("应提升为 hosts 键")
	}
	if _, ok := root["query"]; !ok {
		t.Fatal("其余顶层键应保留")
	}
}

// 已存在 hosts 键时:直接返回,不改动 debug。
func TestHostsPrefersNewKey(t *testing.T) {
	root := map[string]any{
		"hosts": map[string]any{"activeEnv": "new"},
		"debug": map[string]any{"activeEnv": "old"},
	}
	h, err := Hosts(root)
	if err != nil {
		t.Fatalf("Hosts: %v", err)
	}
	if h["activeEnv"] != "new" {
		t.Fatalf("应优先 hosts, got %v", h["activeEnv"])
	}
	if _, ok := root["debug"]; !ok {
		t.Fatal("已存在 hosts 时不应改动 debug")
	}
}

func TestHostsMissing(t *testing.T) {
	if _, err := Hosts(map[string]any{}); err == nil {
		t.Fatal("缺少配置节应报错")
	}
}

// Edit 写路径:旧 debug 键在落盘时迁移为 hosts。
func TestEditMigratesAndPersists(t *testing.T) {
	p := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(p, []byte(`{"debug":{"activeEnv":"e1"},"query":{"source":"local"}}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := Edit(p, nil, func(root map[string]any) error {
		h, err := Hosts(root)
		if err != nil {
			return err
		}
		h["activeEnv"] = "e2"
		return nil
	}); err != nil {
		t.Fatalf("Edit: %v", err)
	}
	root, err := Open(p)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := root["debug"]; ok {
		t.Fatal("落盘后不应再有 debug 键")
	}
	if _, ok := root["query"]; !ok {
		t.Fatal("其余顶层键应保留")
	}
	h, _ := root["hosts"].(map[string]any)
	if h == nil || h["activeEnv"] != "e2" {
		t.Fatalf("hosts = %v, want activeEnv e2", root["hosts"])
	}
}
