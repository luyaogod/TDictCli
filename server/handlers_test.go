package server

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func writeTempConfig(t *testing.T, body string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatalf("写入临时配置失败: %v", err)
	}
	return p
}

func readRoot(t *testing.T, p string) map[string]any {
	t.Helper()
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	var root map[string]any
	if err := json.Unmarshal(b, &root); err != nil {
		t.Fatal(err)
	}
	return root
}

// GET /api/config 返回 hosts 节;PUT 整块替换并保留其余顶层键。
func TestConfigGetPut(t *testing.T) {
	p := writeTempConfig(t, `{"hosts":{"activeEnv":"a","sshs":[{"name":"a","host":"1.2.3.4","port":22,"user":"u","password":"p"}]},"query":{"source":"local"}}`)
	s := New(p, "", nil)

	rec := httptest.NewRecorder()
	s.hConfigGet(rec, httptest.NewRequest("GET", "/api/config", nil))
	if rec.Code != 200 {
		t.Fatalf("GET code=%d", rec.Code)
	}
	var got configResp
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if !got.Exists || len(got.SSHs) != 1 || got.SSHs[0].Host != "1.2.3.4" {
		t.Fatalf("GET 结果不符: %+v", got)
	}

	body := `{"activeEnv":"b","sshs":[
      {"name":"a","host":"1.2.3.4","port":22,"user":"u","password":"p"},
      {"name":"b","host":"5.6.7.8","port":22,"user":"u","password":"p","db":{"type":"oracle","host":"5.6.7.8","port":1521,"service":"t35prd","accounts":[{"account":"ds","password":"ds"}]}}]}`
	rec2 := httptest.NewRecorder()
	s.hConfigPut(rec2, httptest.NewRequest("PUT", "/api/config", bytes.NewBufferString(body)))
	if rec2.Code != 200 {
		t.Fatalf("PUT code=%d body=%s", rec2.Code, rec2.Body.String())
	}
	root := readRoot(t, p)
	if _, ok := root["query"]; !ok {
		t.Fatal("其余顶层键 query 应保留")
	}
	h, _ := root["hosts"].(map[string]any)
	if h == nil || h["activeEnv"] != "b" {
		t.Fatalf("hosts.activeEnv 应为 b: %v", root["hosts"])
	}
	if sshs, _ := h["sshs"].([]any); len(sshs) != 2 {
		t.Fatalf("hosts.sshs 应为 2 条: %v", h["sshs"])
	}
}

// 旧键 debug 在保存时迁移为 hosts 并删除。
func TestConfigPutMigratesLegacyKey(t *testing.T) {
	p := writeTempConfig(t, `{"debug":{"activeEnv":"a","sshs":[{"name":"a","host":"1.2.3.4","port":22,"user":"u","password":"p"}]}}`)
	s := New(p, "", nil)
	body := `{"activeEnv":"a","sshs":[{"name":"a","host":"1.2.3.4","port":22,"user":"u","password":"p"}]}`
	rec := httptest.NewRecorder()
	s.hConfigPut(rec, httptest.NewRequest("PUT", "/api/config", bytes.NewBufferString(body)))
	if rec.Code != 200 {
		t.Fatalf("PUT code=%d", rec.Code)
	}
	root := readRoot(t, p)
	if _, ok := root["debug"]; ok {
		t.Fatal("旧键 debug 应被删除")
	}
	if _, ok := root["hosts"]; !ok {
		t.Fatal("应写入 hosts 键")
	}
}

// 环境缺少主机 / 空列表均拒绝保存。
func TestConfigPutValidation(t *testing.T) {
	p := writeTempConfig(t, `{"hosts":{"sshs":[{"name":"a","host":"1.2.3.4"}]}}`)
	s := New(p, "", nil)

	rec := httptest.NewRecorder()
	s.hConfigPut(rec, httptest.NewRequest("PUT", "/api/config", bytes.NewBufferString(`{"sshs":[]}`)))
	if !bytes.Contains(rec.Body.Bytes(), []byte(`"ok":false`)) {
		t.Fatalf("空列表应被拒绝: %s", rec.Body.String())
	}

	rec2 := httptest.NewRecorder()
	s.hConfigPut(rec2, httptest.NewRequest("PUT", "/api/config", bytes.NewBufferString(`{"sshs":[{"name":"x","host":""}]}`)))
	if !bytes.Contains(rec2.Body.Bytes(), []byte(`"ok":false`)) {
		t.Fatalf("缺主机应被拒绝: %s", rec2.Body.String())
	}
}

// config.json 不存在时首次保存创建文件。
func TestConfigPutCreatesFile(t *testing.T) {
	p := filepath.Join(t.TempDir(), "config.json")
	s := New(p, "", nil)
	body := `{"activeEnv":"a","sshs":[{"name":"a","host":"1.2.3.4","port":22,"user":"u","password":"p"}]}`
	rec := httptest.NewRecorder()
	s.hConfigPut(rec, httptest.NewRequest("PUT", "/api/config", bytes.NewBufferString(body)))
	if rec.Code != 200 {
		t.Fatalf("PUT code=%d body=%s", rec.Code, rec.Body.String())
	}
	if _, err := os.Stat(p); err != nil {
		t.Fatalf("应创建配置文件: %v", err)
	}
}
