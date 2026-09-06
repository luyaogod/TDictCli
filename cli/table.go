package cli

import (
	"fmt"
	"strings"

	"tdict/db"
	"tdict/output"

	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:     "rt <table_name>",
	Aliases: []string{"table"},
	Short:   "查询数据表字典",
	Long: `查询一张或多张数据表的字典:这张表在系统里做什么(表说明/所属模块/表类型),
以及字段(中文含义/数据类型/长度/主键/必填)、键值与索引。
读代码、看 SQL、查界面字段含义时用它。
支持逗号分隔多个表名;输出简体中文。`,
	Example: `  tdict rt dzea_t
  tdict rt "dzea_t,dzeb_t,dzed_t"
  tdict rt dzea_t --json
  tdict rt dzea_t --conn 主机正式区   # 切到某环境的远程库直查(--conn local 回本地)`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tables := splitNames(args[0])
		if len(tables) == 0 {
			return fmt.Errorf("未指定有效的表名")
		}

		var dicts []*db.TableDict
		for _, t := range tables {
			d, err := queryTableDict(t)
			if err != nil {
				return err
			}
			dicts = append(dicts, d)
		}

		if IsJSON() {
			return output.PrintJSON(dicts)
		}

		if IsCSV() {
			return printTableCSV(dicts)
		}

		for _, d := range dicts {
			printTableDict(d)
		}
		return nil
	},
}

// queryTableDict assembles a table's complete dictionary (meta + fields + keys + indexes).
func queryTableDict(name string) (*db.TableDict, error) {
	d := &db.TableDict{TableName: name}

	meta, err := GetDB().QueryTableMeta(name)
	if err != nil {
		return nil, err
	}
	if meta != nil {
		d.TableName = meta.TableName
		d.TableDesc = meta.TableDesc
		d.Module = meta.Module
		d.TableType = meta.TableType
	}

	fields, err := GetDB().QueryTable(name)
	if err != nil {
		return nil, err
	}
	d.Fields = fields

	// 键值/索引依赖 dzed_t/dzec_t，部分库可能未含这两张表
	keys, err := GetDB().QueryKeys(name)
	if err != nil {
		if db.IsMissingTable(err) {
			keys = nil
		} else {
			return nil, err
		}
	}
	d.Keys = keys

	indexes, err := GetDB().QueryIndexes(name)
	if err != nil {
		if db.IsMissingTable(err) {
			indexes = nil
		} else {
			return nil, err
		}
	}
	d.Indexes = indexes

	return d, nil
}

// printTableDict prints a table's dictionary as plain text sections.
func printTableDict(d *db.TableDict) {
	if d.TableDesc == "" && d.Module == "" && d.TableType == "" && len(d.Fields) == 0 {
		fmt.Printf("未找到表 '%s' 或该表无字段定义。\n\n", d.TableName)
		return
	}
	fmt.Printf("=== %s ===\n", d.TableName)
	if d.TableDesc != "" {
		fmt.Printf("表说明: %s\n", d.TableDesc)
	}
	var metaParts []string
	if d.Module != "" {
		metaParts = append(metaParts, "模块: "+d.Module)
	}
	if d.TableType != "" {
		metaParts = append(metaParts, "类型: "+d.TableType)
	}
	if len(metaParts) > 0 {
		fmt.Printf("%s\n", strings.Join(metaParts, "    "))
	}

	fmt.Printf("\n字段 (%d):\n", len(d.Fields))
	if len(d.Fields) == 0 {
		fmt.Println("  (无)")
	} else {
		headers := []string{"序号", "字段名", "字段说明", "数据类型", "长度", "主键", "必填", "备注"}
		var rows [][]string
		for _, f := range d.Fields {
			rows = append(rows, []string{f.Seq, f.FieldName, f.FieldDesc,
				f.DataType, f.Length, f.IsPK, f.Required, f.Remark})
		}
		output.PrintTable(headers, rows)
	}

	fmt.Printf("\n键值 (%d):\n", len(d.Keys))
	if len(d.Keys) == 0 {
		fmt.Println("  (无)")
	} else {
		headers := []string{"键名", "类型", "键值字段", "外键表", "外键字段"}
		var rows [][]string
		for _, k := range d.Keys {
			rows = append(rows, []string{k.KeyName, keyTypeLabel(k.KeyType), k.KeyField, k.RefTable, k.RefField})
		}
		output.PrintTable(headers, rows)
	}

	fmt.Printf("\n索引 (%d):\n", len(d.Indexes))
	if len(d.Indexes) == 0 {
		fmt.Println("  (无)")
	} else {
		headers := []string{"索引名", "类型", "索引字段"}
		var rows [][]string
		for _, ix := range d.Indexes {
			rows = append(rows, []string{ix.IndexName, ix.IndexType, ix.IndexField})
		}
		output.PrintTable(headers, rows)
	}
	fmt.Println()
}

// printTableCSV flattens the field definitions across tables into CSV rows.
func printTableCSV(dicts []*db.TableDict) error {
	headers := []string{"表名", "序号", "字段名", "字段说明", "数据类型", "长度", "主键", "必填", "备注"}
	var rows [][]string
	for _, d := range dicts {
		for _, f := range d.Fields {
			rows = append(rows, []string{d.TableName, f.Seq, f.FieldName, f.FieldDesc,
				f.DataType, f.Length, f.IsPK, f.Required, f.Remark})
		}
	}
	return output.PrintCSVFromMaps(headers, rows)
}

// keyTypeLabel maps dzed_t key type codes to Chinese labels.
func keyTypeLabel(code string) string {
	switch code {
	case "P":
		return "主键"
	case "F":
		return "外键"
	case "U":
		return "唯一"
	default:
		return code
	}
}

func init() {
	rootCmd.AddCommand(tableCmd)
}

func splitNames(s string) []string {
	parts := strings.Split(s, ",")
	var result []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}
