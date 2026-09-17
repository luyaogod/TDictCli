package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

// mkSkills 在 dir 下建一个技能源树:每个技能一个目录 + SKILL.md。
func mkSkills(t *testing.T, dir string, names ...string) {
	t.Helper()
	for _, n := range names {
		p := filepath.Join(dir, n)
		if err := os.MkdirAll(p, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(p, "SKILL.md"), []byte("---\nname: "+n+"\n---\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func TestListSkills(t *testing.T) {
	src := t.TempDir()
	mkSkills(t, src, "beta", "alpha")
	got, err := listSkills(src)
	if err != nil {
		t.Fatalf("listSkills: %v", err)
	}
	if len(got) != 2 || got[0] != "alpha" || got[1] != "beta" {
		t.Fatalf("listSkills = %v, want [alpha beta](按名字排序)", got)
	}

	// 技能目录缺 SKILL.md 必须报错
	if err := os.MkdirAll(filepath.Join(src, "broken"), 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := listSkills(src); err == nil {
		t.Fatal("技能目录缺 SKILL.md 时应报错")
	}

	// 空源目录必须报错
	if _, err := listSkills(t.TempDir()); err == nil {
		t.Fatal("skills 源目录为空时应报错")
	}
}

func TestInstallSkillsTreeCopies(t *testing.T) {
	src := t.TempDir()
	mkSkills(t, src, "alpha", "beta")
	dst := filepath.Join(t.TempDir(), "skills")

	copied, err := installSkillsTree(src, dst, false)
	if err != nil {
		t.Fatalf("installSkillsTree: %v", err)
	}
	want := []string{"alpha/SKILL.md", "beta/SKILL.md"}
	if len(copied) != len(want) {
		t.Fatalf("copied = %v, want %v", copied, want)
	}
	for i, w := range want {
		if filepath.ToSlash(copied[i]) != w {
			t.Fatalf("copied[%d] = %q, want %q", i, copied[i], w)
		}
	}

	// 装完的目录必须是「技能名/SKILL.md」形态,不是扁平 .md
	for _, n := range []string{"alpha", "beta"} {
		if _, err := os.Stat(filepath.Join(dst, n, "SKILL.md")); err != nil {
			t.Fatalf("%s/SKILL.md 不存在: %v", n, err)
		}
	}
}

func TestInstallSkillsTreeConflict(t *testing.T) {
	src := t.TempDir()
	mkSkills(t, src, "alpha")
	dst := t.TempDir()
	mkSkills(t, dst, "alpha") // 目标已有同名技能

	if _, err := installSkillsTree(src, dst, false); err == nil {
		t.Fatal("目标已存在同名技能目录时,不加 --force 应拒绝")
	}
	if _, err := installSkillsTree(src, dst, true); err != nil {
		t.Fatalf("--force 应覆盖: %v", err)
	}
}

func TestInstallSkillsTreeRejectsSameDir(t *testing.T) {
	src := t.TempDir()
	mkSkills(t, src, "alpha")
	if _, err := installSkillsTree(src, src, true); err == nil {
		t.Fatal("目标与源相同应报错(而不是自我覆盖)")
	}
}

// install 不依赖本地字典库:必须覆盖 root 的 SQLite 打开钩子,
// 否则在没有 erp_data.db 的目录(便携版/新项目)里会报"数据库文件未找到"。
func TestInstallSkipsSQLiteHook(t *testing.T) {
	for _, c := range []*cobra.Command{installCmd, installSkillsCmd, installPathCmd} {
		if c.PersistentPreRunE == nil {
			t.Fatalf("%s 应定义 PersistentPreRunE 以覆盖 root 的 SQLite 钩子", c.Name())
		}
		if err := c.PersistentPreRunE(c, nil); err != nil {
			t.Fatalf("%s 的 PersistentPreRunE 应返回 nil: %v", c.Name(), err)
		}
	}
}
