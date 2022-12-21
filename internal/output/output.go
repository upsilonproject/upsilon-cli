package output

import (
	"encoding/json"
	table "github.com/jedib0t/go-pretty/table"
	text "github.com/jedib0t/go-pretty/text"
	. "github.com/upsilonproject/upsilon-cli/internal/runtimeconfig"

	log "github.com/sirupsen/logrus"
)

type TableRow = map[string]string
type DataTable struct {
	Headers []interface{}
	Rows map[int]*TableRow
}

func (dt *DataTable) Append(row *TableRow) {
	idx := len(dt.Rows)
	dt.Rows[idx] = row
}

func Format(rows *DataTable) string {
	switch RuntimeConfig.OutputFormat {
	case "table":
		return formatOutputTable(rows)
	case "json":
		return formatOutputJson(rows)
	default:
		log.Errorf("Unsupported format")
		return ""
	}

}

func formatOutputJson(rows *DataTable) string {
	ret, _ := json.Marshal(rows)

	return string(ret)
}

func formatOutputTable(dataTable *DataTable) string {
	karmaTransformer := text.Transformer(func(val interface{}) string {
		return text.FgRed.Sprint(val)
	})

	tbl := table.NewWriter()
	tbl.AppendHeader(dataTable.Headers)
	tbl.SetStyle(table.StyleLight)
//	tbl.Style().Color.Header = text.Colors{text.Bold}
	tbl.SetColumnConfigs([]table.ColumnConfig {
		{
			ColorsHeader: text.Colors{text.Bold},
			Transformer: karmaTransformer,
		},
	})
	tbl.Style().Options.DrawBorder = false

	for i, _ := range dataTable.Rows {
		row := dataTable.Rows[i] // because range is nondeterministic
		var cells []interface{} 

		for _, hdr := range dataTable.Headers {
			cells = append(cells, (*row)[hdr.(string)])
		}

		tbl.AppendRow(cells)
	}


	return tbl.Render() + "\n"
}

func NewDataTable(headers []interface{}) *DataTable {
	tbl := &DataTable {
		Headers: headers,
	}

	tbl.Rows = make(map[int]*TableRow)

	return tbl
}
