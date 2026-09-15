package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateInstallArg(t *testing.T) {
	for _, ok := range [][]string{nil, {}, {"skills"}, {"SKILLS"}, {" skills "}} {
		if err := validateInstallArg(ok); err != nil {
			t.Errorf("args %v 应通过, got %v", ok, err)
		}
	}
	for _, bad := range [][]string{{"out"}, {"foo"}, {"skills", "extra"}} {
		if err := validateInstallArg(bad); err == nil {
			t.Errorf("args %v 应被拒绝", bad)
		}
	}
}

func TestSameDir(t *testing.T) {
	d := t.TempDir()
	if !sameDir(d, d+string(filepath.Separator)) {
		t.Error("同一目录(带尾分隔符)应判为相同")
	}
	if sameDir(d, filepath.Join(d, "sub")) {
		t.Error("不同目录不应判为相同")
	}
}

// copyTree 递归复制(含子目录),文件数与内容正确。
func TestCopyTree(t *testing.T) {
	src := t.TempDir()
	dst := filepath.Join(t.TempDir(), "skills")
	mk := func(rel, content string) {
		p := filepath.Join(src, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	mk("tdict.md", "A")
	mk("erp-read.md", "B")
	mk("nested/deep.md", "C")

	n, err := copyTree(src, dst)
	if err != nil {
		t.Fatalf("copyTree: %v", err)
	}
	if n != 3 {
		t.Fatalf("复制 %d 个文件, want 3", n)
	}
	for rel, want := range map[string]string{"tdict.md": "A", "erp-read.md": "B", "nested/deep.md": "C"} {
		b, err := os.ReadFile(filepath.Join(dst, rel))
		if err != nil {
			t.Fatalf("读取 %s: %v", rel, err)
		}
		if string(b) != want {
			t.Fatalf("%s = %q, want %q", rel, b, want)
		}
	}
}

// 目标目录已存在时应覆盖同名文件(不是嵌套 skills/skills)。
func TestCopyTreeOverwrites(t *testing.T) {
	src := t.TempDir()
	dst := t.TempDir()
	if err := os.WriteFile(filepath.Join(src, "a.md"), []byte("new"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dst, "a.md"), []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := copyTree(src, dst); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(filepath.Join(dst, "a.md"))
	if string(b) != "new" {
		t.Fatalf("a.md = %q, want new", b)
	}
}

// install 不依赖本地字典库:必须覆盖 root 的 SQLite 打开钩子,
// 否则在没有 erp_data.db 的目录(便携版/新项目)里会报"数据库文件未找到"。
func TestInstallSkipsSQLiteHook(t *testing.T) {
	if installCmd.PersistentPreRunE == nil {
		t.Fatal("installCmd 应定义 PersistentPreRunE 以覆盖 root 的 SQLite 钩子")
	}
	if err := installCmd.PersistentPreRunE(installCmd, nil); err != nil {
		t.Fatalf("PersistentPreRunE 应返回 nil: %v", err)
	}
}
