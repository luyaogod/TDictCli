package debug

import "testing"

// hTopent 服务端 trim 语义(文本值 + 剔除两侧空白)——纯字符串逻辑验证
func TestTopentTrimSemantics(t *testing.T) {
	cases := map[string]string{
		"  99 ":       "99",
		"\tabc-99_X ": "abc-99_X",
		"   ":         "",
		"99":          "99",
	}
	for in, want := range cases {
		if got := trimTopent(in); got != want {
			t.Fatalf("trimTopent(%q) = %q, want %q", in, got, want)
		}
	}
}
