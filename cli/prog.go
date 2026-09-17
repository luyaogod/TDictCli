package cli

import (
	"fmt"

	"tdict/db"
	"tdict/output"

	"github.com/spf13/cobra"
)

var (
	progKW    string
	progLang  string
	progLimit int
)

var progCmd = &cobra.Command{
	Use:     "prog [程序编号]",
	Aliases: []string{"program", "job"},
	Short:   "查询程序/作业字典 (程序做什么、被哪些作业使用)",
	Long: `查询 T100 的程序与作业登记(azzi900 程式基本資料 / azzi910 作業基本資料):
程序编号对应的中文作业名、程序类别、归属模块、是否客制、引用主程序,
以及**哪些作业用了这个程序**——作业通过 gzzz_t.gzzz002 挂到程序上,
一个程序可以被多个作业使用(作业名称取它所挂程序的名称,与 azzi910 一致)。

读代码时的第一个问题"这个程序做什么"就用它;也可以拿作业编号去查它挂的是哪个程序。
无参数时列出全部程序(--kw 按程序编号/中文名称搜索,从业务词找程序)。
程序名称有简体(--lang zh_CN,默认)与繁体(--lang zh_TW)两份,搜索用字要对应;
一个程序被很多作业使用时(如共用维护程序可达数百个),作业表默认只列前 --limit 行。
数据族为"程序与作业"(--help 末尾会显示本地是否已同步;缺就 tdict db sync)。`,
	Example: `  tdict prog aapi011          # 程序做什么 + 被哪些作业使用
  tdict prog --kw 对帐        # 按中文作业名找程序
  tdict prog aapit100 --json
  tdict prog aapi011 --conn 正式区`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return runProgList()
		}
		return runProgDetail(args[0])
	},
}

// runProgDetail 打印一个程序编号的详情:程序信息 + 使用它的作业。
// 编号只登记为作业时,跟进它挂的程序再打印一遍(一眼看到"谁在用它")。
func runProgDetail(code string) error {
	info, err := GetDB().QueryProgInfo(code, progLang)
	if err != nil {
		if db.IsMissingTable(err) {
			fmt.Println(missingHint("程序与作业 (gzza_t/gzzz_t 等表)"))
			return nil
		}
		return err
	}
	if info == nil {
		fmt.Printf("未找到程序/作业 '%s'。\n", code)
		fmt.Println("提示: 用业务词搜索试试 —— tdict prog --kw <关键字>")
		return nil
	}

	if IsJSON() {
		jobs, err := GetDB().QueryProgJobs(code, progLang)
		if err != nil && !db.IsMissingTable(err) {
			return err
		}
		return output.PrintJSON(map[string]any{"程序": info, "作业": jobs})
	}

	if info.IsProg {
		printProgInfo(info)
	} else {
		fmt.Printf("=== %s ===\n该编号不是程序登记,是作业:挂的程序 = %s\n", code, orDash(info.JobProg))
		if info.JobProg != "" && info.JobProg != code {
			// 跟到它挂的程序:程序信息 + 作业清单(本作业会出现在其中)
			pi, err := GetDB().QueryProgInfo(info.JobProg, progLang)
			if err != nil && !db.IsMissingTable(err) {
				return err
			}
			if pi != nil && pi.IsProg {
				fmt.Println()
				printProgInfo(pi)
			}
		}
		return nil
	}

	jobs, err := GetDB().QueryProgJobs(code, progLang)
	if err != nil {
		if db.IsMissingTable(err) {
			fmt.Println(missingHint("程序与作业 (gzzz_t 等表)"))
			return nil
		}
		return err
	}
	fmt.Printf("\n使用它的作业 (%d):\n", len(jobs))
	if len(jobs) == 0 {
		fmt.Println("  (无;该程序未被作业登记,或是被作其他用途引用)")
		return nil
	}
	headers := []string{"作业编号", "作业名称", "归属模块", "应用参数组", "参数组说明", "默认单据性质"}
	shown := jobs
	if progLimit > 0 && len(jobs) > progLimit {
		shown = jobs[:progLimit]
	}
	var rows [][]string
	for _, j := range shown {
		rows = append(rows, []string{j.JobCode, j.JobName, j.Module, j.ParamGrp, j.ParamDesc, j.DocType})
	}
	output.PrintTable(headers, rows)
	if len(shown) < len(jobs) {
		fmt.Printf("\n… 还有 %d 个未显示(共 %d 个);--limit 0 显示全部,或 --json 导出\n",
			len(jobs)-len(shown), len(jobs))
	}
	return nil
}

// printProgInfo 打印程序登记信息。
func printProgInfo(p *db.ProgInfo) {
	fmt.Printf("=== %s ===\n", p.Code)
	if p.Name != "" {
		fmt.Printf("程序名称: %s\n", p.Name)
	}
	if p.ShortName != "" {
		fmt.Printf("程序简称: %s\n", p.ShortName)
	}
	parts := make([]string, 0, 5)
	if p.Category != "" {
		parts = append(parts, "程序类别: "+p.Category+categoryLabel(p.Category))
	}
	if p.Module != "" {
		parts = append(parts, "归属模块: "+p.Module)
	}
	if p.Cust != "" {
		parts = append(parts, "客制: "+p.Cust)
	}
	if p.Status != "" {
		parts = append(parts, "状态码: "+p.Status)
	}
	for _, s := range parts {
		fmt.Println(s)
	}
	if p.RefMain != "" {		fmt.Printf("引用主程序: %s\n", p.RefMain)
	}
	if p.RunCmd != "" {
		fmt.Printf("系统运行指令: %s\n", p.RunCmd)
	}
}

// runProgList 列表模式:列出/搜索程序(附各自的作业数)。
func runProgList() error {
	list, err := GetDB().QueryProgList(progLang, progKW)
	if err != nil {
		if db.IsMissingTable(err) {
			fmt.Println(missingHint("程序与作业 (gzza_t 等表)"))
			return nil
		}
		return err
	}
	if len(list) == 0 {
		if progKW != "" {
			fmt.Printf("没有匹配 '%s' 的程序。\n提示: 程序名称分简体(--lang zh_CN,默认)与繁体(--lang zh_TW)两份,用字可能是异体(如 对账/对帐/對帳);换个写法、换 --lang 或用更短的词再试。\n", progKW)
		} else {
			fmt.Println(emptyHint("程序与作业"))
		}
		return nil
	}

	headers := []string{"程序编号", "程序名称", "程序类别", "归属模块", "客制", "作业数"}
	var rows [][]string
	for _, p := range list {
		rows = append(rows, []string{p.Code, p.Name, p.Category, p.Module, p.Cust, fmt.Sprintf("%d", p.JobCount)})
	}
	if IsJSON() {
		return output.PrintJSON(list)
	}
	if IsCSV() {
		return output.PrintCSVFromMaps(headers, rows)
	}
	output.PrintTable(headers, rows)
	fmt.Printf("\n共 %d 个程序", len(list))
	if progKW == "" {
		fmt.Print(";用 --kw <业务词> 按程序编号/中文名称过滤,如 tdict prog --kw 对帐")
	}
	fmt.Println()
	return nil
}

// categoryLabel 程序类别码的中文补充(取 T100 程序编号第 4 码的类别约定)。
func categoryLabel(code string) string {
	switch code {
	case "i":
		return "(基本资料维护)"
	case "m":
		return "(主档维护)"
	case "t":
		return "(交易处理)"
	case "s":
		return "(参数设定)"
	case "p":
		return "(批次处理)"
	case "q":
		return "(查询)"
	case "r":
		return "(报表)"
	default:
		return ""
	}
}

func orDash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

func init() {
	progCmd.Flags().StringVar(&progKW, "kw", "", "按程序编号/中文名称搜索 (无编号时的列表模式)")
	progCmd.Flags().StringVar(&progLang, "lang", "zh_CN", "程序名称语言别 (gzzal002, 默认 zh_CN)")
	progCmd.Flags().IntVar(&progLimit, "limit", 20, "作业最多显示几行 (0 = 全部;大程序可被数百个作业使用)")
	rootCmd.AddCommand(progCmd)
}
