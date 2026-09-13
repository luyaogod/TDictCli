package host

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestMirrorEnvDirAndReady(t *testing.T) {
	dir := t.TempDir()
	if got := MirrorEnvDir(dir, "envA"); got != filepath.Join(dir, "envA") {
		t.Fatalf("MirrorEnvDir = %q", got)
	}
	if MirrorEnvDir("", "envA") != "" || MirrorEnvDir(dir, "") != "" {
		t.Fatal("空参数应返回空串")
	}
	if MirrorReady(dir, "envA") {
		t.Fatal("无基线标记应为 false")
	}
	if err := writeLocalMark(filepath.Join(dir, "envA")); err != nil {
		t.Fatalf("writeLocalMark: %v", err)
	}
	if !MirrorReady(dir, "envA") {
		t.Fatal("有基线标记应为 true")
	}
	if MirrorReady("", "envA") {
		t.Fatal("空镜像根应为 false")
	}
}

// countingReader 统计压缩字节,并在 EOF 时回调一次进度(小数据只触发收尾回调)。
func TestCountingReaderProgress(t *testing.T) {
	var got []MirrorProgress
	cr := &countingReader{
		r:      strings.NewReader("hello world"),
		total:  11,
		report: func(p MirrorProgress) { got = append(got, p) },
	}
	buf := make([]byte, 4)
	for {
		if _, err := cr.Read(buf); err != nil {
			break
		}
	}
	if cr.n != 11 {
		t.Fatalf("已读字节 = %d, want 11", cr.n)
	}
	if len(got) == 0 {
		t.Fatal("应至少回调一次进度")
	}
	last := got[len(got)-1]
	if last.Phase != "download" || last.Bytes != 11 || last.Total != 11 {
		t.Fatalf("末次进度 = %+v", last)
	}
}
