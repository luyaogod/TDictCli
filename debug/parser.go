package debug

import (
	"regexp"
	"strings"
)

// ---------- 协议锚点正则(全部来自 172.16.1.09 真机实测输出) ----------

var (
	// ANSI 转义:CSI 序列 + 单字符转义(= > 等 keypad 模式)
	ansiRe = regexp.MustCompile("\x1b\\[[0-9;?]*[a-zA-Z]|\x1b[=>]")

	// 断点命中:`Breakpoint 1, asf_bsft001_wf.bsft001_wf_construct() at asf_bsft001_wf.4gl:4452`
	reBreakHit = regexp.MustCompile(`^Breakpoint (\d+), (.+?) at (.+?):(\d+)\s*$`)
	// 人工中断(SIGINT 后;终端 ^C 回显可能与字样合并为一行,用子串匹配)
	reInterrupt = regexp.MustCompile(`INTERRUPT`)
	// 步进/位置头:`cl_ap_code_fuzzyquery() at lib_cl_ap_code.4gl:4593`(函数变化时出现)
	reStopHeader = regexp.MustCompile(`^([A-Za-z_][\w.]*)\(\) at (.+?):(\d+)\s*$`)
	// 调用栈帧:`#0 asf_bsft001_wf.bsft001_wf_construct() at asf_bsft001_wf.4gl:4452`
	reFrame = regexp.MustCompile(`^#(\d+)\s+(.+?)\s+at\s+(.+?):(\d+)\s*$`)
	// (fgldb) 提示符:裸提示符(后随 ANSI 残留 = 或 >),命令回显行不算
	rePrompt = regexp.MustCompile(`^\(fgldb\)\s*[=>]?\s*$`)
	// shell 提示符:`<t35prd:/u1/t35prd> `
	reShellPrompt = regexp.MustCompile(`<[A-Za-z0-9_.@-]+:[^>]*>\s*$`)
	// 断点设置成功:`Breakpoint 1 at 0x00000000: file asf_bsft001_wf.4gl, line 4452.`
	reBPSet = regexp.MustCompile(`^Breakpoint (\d+) at \S+: file (.+), line (\d+)\.`)
	// print 结果:`$1 = "CNJ-ICC-250800000012"`(记录/数组为多行)
	rePrintVal = regexp.MustCompile(`^\$(\d+) = (.*)$`)
	// 变量不存在:`No symbol "lp_str" in current context.`
	reNoSymbol = regexp.MustCompile(`No symbol "(.+)" in current context`)
	// 恢复运行
	reContinuing = regexp.MustCompile(`^Continuing\.`)
	// 停站上下文中的源码行:`   4452   LET ...` / `-> 4452   LET ...`
	reSource = regexp.MustCompile(`^(->)?\s*(\d+)\s{2,}(.*)$`)
	// 步进的紧凑停站(源码不可用时):`89	 in lib_cl_ap.4gl`
	reCompactStep = regexp.MustCompile(`^(\d+)\s+in\s+(\S+)\s*$`)
	// info breakpoints 表行:`1   breakpoint     keep y   0x00000000 in asf_bsft001_wf.xx at asf_xx.4gl:4452`
	reBPInfo = regexp.MustCompile(`^(\d+)\s+breakpoint\s+keep\s+(\w)\s+\S+\s+in\s+(\S+)\s+at\s+(.+):(\d+)`)
	// info locals 变量行:`lp_str = "CNJ-ICC-250800000012"`(记录/数组值可续行)
	reLocalVar = regexp.MustCompile(`^([A-Za-z_][\w.]*)\s*=\s*(.*)$`)
	// info variables 声明行(fgldb 实测为 `Globals:` 头 + `名字 类型` 两列,非 name=value)
	reGlobalDecl = regexp.MustCompile(`^([A-Za-z_][\w.]*)\s+([A-Za-z_][\w.\[\], ]*)$`)
	// info functions 函数行:`bsft001_wf_construct ()`
	reFuncList = regexp.MustCompile(`^([A-Za-z_][\w.]*)\s*\(\s*\)\s*$`)
	// info line 位置行:`Line 801 of "asf_bsft001_wf.4gl" starts at address 0x0 <main+0>`
	reInfoLine = regexp.MustCompile(`^Line (\d+) of "([^"]+)"`)
	// until 目标行不存在:`No line 999999 in file asf_bsft001_wf.`
	reNoLine = regexp.MustCompile(`^No line \d+ in file`)
	// until 目标文件不存在:`No source file named bsft001_wf.4gl.`
	reNoSourceFile = regexp.MustCompile(`^No source file named`)
	// 程序退出:`Program exited normally.` / `Program exited with code N.`(作业窗口被关闭或正常结束)
	reProgramExited = regexp.MustCompile(`^Program exited`)
)

// fdbErrRes 已知的调试器错误行(出现在响应中时转成命令错误)
var fdbErrRes = []*regexp.Regexp{
	reNoSymbol,
	regexp.MustCompile(`^The program is not being run`),
	regexp.MustCompile(`^No stack\.`),
	regexp.MustCompile(`^Cannot execute this command`),
	regexp.MustCompile(`^Program not being run`),
	regexp.MustCompile(`^invalid argument`),
	reNoLine,
	reNoSourceFile,
}

// matchFdbErr 返回匹配到的错误行文本
func matchFdbErr(ln string) string {
	for _, re := range fdbErrRes {
		if re.MatchString(ln) {
			return ln
		}
	}
	return ""
}

// isBarePrompt 判断是否为裸 (fgldb) 提示符(命令回显行不算)
func isBarePrompt(ln string) bool { return rePrompt.MatchString(ln) }

// StripANSI 剥离转义序列与残留控制字符(保留 \t)
func StripANSI(s string) string {
	s = ansiRe.ReplaceAllString(s, "")
	return strings.Map(func(r rune) rune {
		if r < 0x20 && r != '\t' {
			return -1
		}
		return r
	}, s)
}

// LineParser 把 PTY 字节流组装成完整行(处理跨包断行;\r 视为行结束)
type LineParser struct {
	cur strings.Builder
}

// Feed 喂入原始字节,返回已完成的行(已剥离 ANSI/控制字符,已 trim 尾部 \r)
func (p *LineParser) Feed(data []byte) []string {
	var out []string
	for _, b := range data {
		switch b {
		case '\n':
			out = append(out, p.flush())
		case '\r':
			// \r\n 或单独 \r 都视为行结束
			out = append(out, p.flush())
		default:
			p.cur.WriteByte(b)
		}
	}
	return out
}

// PartialStr 返回当前未完成半行的内容(不出行)
func (p *LineParser) PartialStr() string {
	if p.cur.Len() == 0 {
		return ""
	}
	return StripANSI(p.cur.String())
}

// FlushPartial 强制出行半行(用于提示符检测:提示符后面没有换行)
func (p *LineParser) FlushPartial() string {
	if p.cur.Len() == 0 {
		return ""
	}
	return p.flush()
}

// Rest 返回未完成的半行(供会话结束时输出)
func (p *LineParser) Rest() string {
	return p.FlushPartial()
}

func (p *LineParser) flush() string {
	s := StripANSI(p.cur.String())
	p.cur.Reset()
	return strings.TrimRight(s, " ")
}
