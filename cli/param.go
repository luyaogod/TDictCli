package cli

// tdict sysp / docp:查询参数定义档(gzsz_t,由 azzzi990 系统参数定义作业与
// azzi991 单据别参数维护作业维护)。两个作业操作同一张定义表,区别在参数群
// (gzsz001):azzi990 视域 = 非 ooac_t(A 系统级 gzsa_t / E 企业级 ooaa_t /
// S 据点级 ooab_t),azzi991 视域 = 恒 ooac_t(单据别 D 级,子表 gzsy_t 绑定
// 单据性质)。参数编号形如 A-SYS-0100 / E-CIR-0001 / D-MFG-0076。
//
// 本命令只查"定义与说明"(名称/说明来自 gzszl_t 多语言表);参数当前值在
// 客户化值表(gzsa_t/ooaa_t/ooab_t/ooac_t),不在查询范围。语言策略同 msg:
// 默认显示 --lang(zh_CN)行,无该语言行时列出可用语言。

import (
	"fmt"
	"strings"

	"tdict/db"
	"tdict/output"

	"github.com/spf13/cobra"
)

var paramLang string

// docpObj 单据别参数的 JSON 展示对象。
type docpObj struct {
	Param db.ParamDefRow  `json:"参数"`
	Docs  []db.DocTypeRow `json:"单据性质"`
}

var syspCmd = &cobra.Command{
	Use:   "sysp <编号>[,<编号>...]",
	Short: "查询系统参数说明 (azzi990/gzsz_t: A/E/S 级参数定义)",
	Long: `查询 T100 系统参数定义 (gzsz_t,作业 azzi990 维护) — 系统/企业/据点级参数
(不含单据别参数 ooac_t,那是 azzi991/tdict docp 的范围)。
参数编号形如 A-SYS-0100 / E-CIR-0001 / S-BAS-0028 (<型态码>-<领域3码>-<4位流水>;
型态码 A=系统级 E=企业级 S=据点级)。
返回: 名称/说明 (gzszl_t 多语言)、参数群与级别、输入型态 (gzsz003: 1=Y/N 2=整数选项
3=范围设定 4=字符或SCC 5=日期)、领域、预设值 (gzsz008)、值域 (gzsz009)、SCC 选项、
校核/开窗引用、异常处理与修改频度等定义;不查运行时当前值(在客户化值表)。
数据来源: gzsz_t gzszl_t,需先执行 tdict db sync 同步(或用 --conn 远程直查)。`,
	Example: `  tdict sysp A-SYS-0100
  tdict sysp S-BAS-0028 --lang zh_TW
  tdict sysp "A-SYS-0100,E-CIR-0001" --json
  tdict sysp D-MFG-0076            # 单据别参数会提示改用 docp`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runParamQuery(args[0], false)
	},
}

var docpCmd = &cobra.Command{
	Use:   "docp <编号>[,<编号>...]",
	Short: "查询单据参数说明 (azzi991/gzsz_t: D 级单据别参数 + 单据性质)",
	Long: `查询 T100 单据别参数定义 (gzsz_t 的 ooac_t 群,作业 azzi991 维护) — 单据参数
编号形如 D-MFG-0076 / D-BAS-0058 (<型态码 D>-<领域3码>-<4位流水>)。
返回: 名称/说明 (gzszl_t 多语言)、输入型态/值域/预设值等定义,以及该参数绑定的
单据性质清单 (gzsy_t: 模块、单据性质、是否已抛转)。参数当前值在各单据别值表
ooac_t(每单据别一行,由值维护作业 aooi200 维护),不在查询范围。
数据来源: gzsz_t gzszl_t gzsy_t,需先执行 tdict db sync 同步(或用 --conn 远程直查)。`,
	Example: `  tdict docp D-MFG-0076
  tdict docp D-BAS-0058 --lang zh_TW
  tdict docp "D-MFG-0076,D-BAS-0058" --json
  tdict docp A-SYS-0100            # 系统参数会提示改用 sysp`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runParamQuery(args[0], true)
	},
}

func init() {
	rootCmd.AddCommand(syspCmd, docpCmd)
	syspCmd.Flags().StringVar(&paramLang, "lang", "zh_CN", "说明语言别 (gzszl003, 缺省 zh_CN)")
	docpCmd.Flags().StringVar(&paramLang, "lang", "zh_CN", "说明语言别 (gzszl003, 缺省 zh_CN)")
}

// runParamQuery 按视域查询并输出:docOnly=true 走 docp(ooac_t 群 + gzsy 子行)。
func runParamQuery(arg string, docOnly bool) error {
	codes := splitNames(arg)
	if len(codes) == 0 {
		return fmt.Errorf("未指定有效的参数编号")
	}
	other := "docp"
	if docOnly {
		other = "sysp"
	}

	var pickedSys []db.ParamDefRow
	var pickedDoc []docpObj
	for _, code := range codes {
		rows, err := GetDB().QueryParam(code)
		if err != nil {
			if db.IsMissingTable(err) {
				return fmt.Errorf("本地库尚未包含参数档 (gzsz_t/gzszl_t)。请先执行 tdict db sync,或用 --conn <环境名> 远程直查")
			}
			return err
		}
		if len(rows) == 0 {
			if !IsJSON() {
				fmt.Printf("未找到参数编号 '%s'。\n", code)
			}
			continue
		}
		// 视域过滤:sysp 排除 ooac_t;docp 只要 ooac_t
		var kept []db.ParamDefRow
		for i := range rows {
			if (rows[i].Group == "ooac_t") == docOnly {
				kept = append(kept, rows[i])
			}
		}
		if len(kept) == 0 {
			if IsJSON() {
				continue
			}
			fmt.Printf("编号 '%s' 是%s参数;请用 tdict %s 查询。\n", code, paramKind(rows[0].Group), other)
			continue
		}
		match, avail := pickLangRow(kept, paramLang, func(r db.ParamDefRow) string { return r.Lang })
		if match == nil {
			if IsJSON() {
				continue
			}
			fmt.Printf("%s\n", langMissMsg(code, paramLang, avail))
			continue
		}
		if !docOnly {
			pickedSys = append(pickedSys, *match)
			if !IsJSON() {
				printParam(match, nil)
			}
			continue
		}
		// docp:附带单据性质绑定(gzsy_t 缺表时静默降级为空)
		docs, err := GetDB().QueryDocTypes(code, match.Lang)
		if err != nil && !db.IsMissingTable(err) {
			return err
		}
		if err != nil {
			docs = nil
		}
		if IsJSON() {
			pickedDoc = append(pickedDoc, docpObj{Param: *match, Docs: docs})
		} else {
			printParam(match, docs)
		}
	}

	if IsJSON() {
		if docOnly {
			return output.PrintJSON(pickedDoc)
		}
		return output.PrintJSON(pickedSys)
	}
	return nil
}

// paramKind 参数群 → 级别说明(命令提示用)。
func paramKind(group string) string {
	switch group {
	case "ooac_t":
		return "单据别"
	case "gzsa_t":
		return "系统级"
	case "ooaa_t":
		return "企业级"
	case "ooab_t":
		return "据点级"
	}
	return group
}

// paramLevelLabel 参数群 → 中文级别标签(展示用)。
func paramLevelLabel(group string) string {
	switch group {
	case "gzsa_t":
		return "系统级参数"
	case "ooaa_t":
		return "企业级参数"
	case "ooab_t":
		return "据点级参数"
	case "ooac_t":
		return "单据别参数"
	}
	return ""
}

func printParam(m *db.ParamDefRow, docs []db.DocTypeRow) {
	fmt.Printf("=== %s (%s) ===\n", m.Code, m.Lang)
	if m.Name != "" {
		fmt.Printf("名称:     %s\n", m.Name)
	}
	if m.Desc != "" {
		fmt.Printf("说明:     %s\n", m.Desc)
	}
	level := paramLevelLabel(m.Group)
	if level != "" && m.GroupName != "" {
		level += "(" + m.GroupName + ")"
	}
	line := "群:       " + m.Group
	if level != "" {
		line += "  " + level
	}
	fmt.Println(line)
	fmt.Printf("型态:     %s\n", typeWithCode(paramTypeLabel(m.TypeCode), m.TypeCode))
	if m.AreaCode != "" {
		fmt.Printf("领域:     %s\n", m.AreaCode)
	}
	if m.Default != "" {
		fmt.Printf("预设值:   %s\n", m.Default)
	}
	if m.RangeVal != "" {
		fmt.Printf("值域:     %s\n", m.RangeVal)
	}
	if m.DateFmt != "" {
		fmt.Printf("日期格式: %s\n", m.DateFmt)
	}
	if m.SccCode != "" {
		fmt.Printf("SCC选项:  清单 %s(型态 4 字符时的可选项来源)\n", m.SccCode)
	}
	if m.RVCode != "" {
		fmt.Printf("校核带值: %s\n", m.RVCode)
	}
	if m.RQCode != "" {
		fmt.Printf("开窗程序: %s\n", m.RQCode)
	}
	if m.Except != "" {
		fmt.Printf("异常处理: %s\n", typeWithCode(paramExceptLabel(m.Except), m.Except))
	}
	if m.Freq != "" {
		fmt.Printf("修改频度: %s (%s)\n", paramFreqLabel(m.Freq), m.Freq)
	}
	if m.Live != "" {
		fmt.Printf("即时抓取: %s\n", m.Live)
	}
	if m.ValueProg != "" {
		fmt.Printf("值维护作业: %s\n", m.ValueProg)
	}
	if m.Status != "" {
		fmt.Printf("状态:     %s\n", statusLabel(m.Status))
	}
	for _, text := range [][2]string{{"备注一", m.Desc2}, {"备注二", m.Desc3}} {
		if text[1] == "" {
			continue
		}
		fmt.Printf("\n%s (gzszl006/007):\n", text[0])
		for _, l := range strings.Split(text[1], "\n") {
			fmt.Println("  " + l)
		}
	}
	if len(docs) > 0 {
		fmt.Printf("\n单据性质绑定 (gzsy_t, %d):\n", len(docs))
		headers := []string{"单据性质", "名称", "模块", "已抛转"}
		var rows [][]string
		for _, d := range docs {
			rows = append(rows, []string{d.DocType, d.DocName, d.Module, d.Generate})
		}
		output.PrintTable(headers, rows)
	}
	fmt.Println()
}

// paramTypeLabel 输入型态 (SCC'89'): 1=Y/N 2=整数选项 3=范围设定 4=字符或SCC 5=日期。
func paramTypeLabel(code string) string {
	switch code {
	case "1":
		return "Y/N"
	case "2":
		return "整数选项"
	case "3":
		return "范围设定"
	case "4":
		return "字符或SCC"
	case "5":
		return "日期"
	default:
		return code
	}
}

// paramExceptLabel 取参异常处理 (SCC'160'): 1=改抓设定时的预设值 2=显示错误回传NULL 3=不回传错误直接NULL。
func paramExceptLabel(code string) string {
	switch code {
	case "1":
		return "改抓设定时的预设值"
	case "2":
		return "显示错误并回传NULL"
	case "3":
		return "不回传错误直接回传NULL"
	default:
		return code
	}
}

// paramFreqLabel 修改频度 (SCC'258'): Y=未限制 N=配置后不可修改。
func paramFreqLabel(code string) string {
	switch code {
	case "Y":
		return "未限制"
	case "N":
		return "配置后不可修改"
	default:
		return code
	}
}
