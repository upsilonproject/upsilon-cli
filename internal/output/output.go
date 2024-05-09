package output

import (
	"encoding/json"
	table "github.com/jedib0t/go-pretty/table"
	text "github.com/jedib0t/go-pretty/text"
	. "github.com/upsilonproject/upsilon-cli/internal/runtimeconfig"

	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type TableRow = map[string]string

type DataTable struct {
	Headers []interface{}
	Rows map[int]*TableRow
}

var emptyMessage string
var preparedTable table.Writer
var preparedRowCount int
var headers = []string {}
var columnConfigs = []table.ColumnConfig {}

func (dt *DataTable) Append(row *TableRow) {
	idx := len(dt.Rows)
	dt.Rows[idx] = row
}

func Prepare(rows *DataTable, msg string) {
	emptyMessage = msg

	if IsPrettyTable() {
		preparedRowCount = len(rows.Rows)
		preparedTable = FormatOutputPrettyTable(rows)
	}
}

func PrintPrepared() {
	if IsPrettyTable() {
		if preparedRowCount == 0 {
			fmt.Println(emptyMessage)
		} else {
			fmt.Printf(preparedTable.Render() + "\n")
		}
		return
	}

	log.Fatalf("Prepared format is unknown.")
}

func Format(rows *DataTable) string {
	switch RuntimeConfig.OutputFormat {
	case "table":
		tbl := FormatOutputPrettyTable(rows)

		return tbl.Render() + "\n"
	case "json":
		return formatOutputJson(rows)
	default:
		log.Errorf("Unsupported format")
		return ""
	}

}

func IsPrettyTable() bool {
	if RuntimeConfig.OutputFormat == "table" {
		return true
	} else {
		return false
	}
}

func formatOutputJson(rows *DataTable) string {
	ret, _ := json.Marshal(rows)

	return string(ret)
}

func PrettyTableRelativeTimestamps(column int) {
	timestampTransformer := text.Transformer(func(val interface{}) string {
		cellTime, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s", val))

		if err != nil {
			log.Errorf("%v", err)
			return "?"
		}

		diff := time.Now().Sub(cellTime)

		if diff > (time.Hour * 24) {
			return fmt.Sprintf("%v", val)
		} else {
			return diff.Truncate(time.Second).String()
		}
	})

	columnConfigs = append(columnConfigs, table.ColumnConfig {
		Number: column,
		Transformer: timestampTransformer,
	})
}

func PrettyTableAddKarma(column int) {
	karmaTransformer := text.Transformer(func(val interface{}) string {
		if val == "GOOD" {
			return text.FgGreen.Sprint(val)
		}

		if val == "BAD" {
			return text.FgRed.Sprint(val)
		}

		return fmt.Sprint(val)
	})

	columnConfigs = append(columnConfigs, table.ColumnConfig {
		Number: column,
		Transformer: karmaTransformer,
	})
}

func PrettyTableSortBy(columns ...string) {
	sortings := []table.SortBy{}

	for _, column := range columns {
		sortings = append(sortings, table.SortBy{Name: column, Mode: table.Asc})
	}

	preparedTable.SortBy(sortings)
}

func PrettyTableHeaders(newHeaders []string) {
	headers = newHeaders

	preparedTable.SetColumnConfigs(columnConfigs)
}

func FormatOutputPrettyTable(dataTable *DataTable) table.Writer {
	if dataTable == nil {
		log.Warnf("Cannot format a nil DataTable")
		return nil
	}

	if len(dataTable.Headers) == 0 {
		log.Warnf("%+v %v", dataTable, len(dataTable.Rows))
		log.Warnf("Cannot format a DataTable that does not have any headers.")
		return nil
	}

	tbl := table.NewWriter()
	tbl.AppendHeader(dataTable.Headers)
	tbl.SetStyle(table.StyleLight)
	tbl.Style().Color.Header = text.Colors{text.Bold}
	tbl.Style().Options.DrawBorder = false
	tbl.SortBy([]table.SortBy{
		{Name: dataTable.Headers[0].(string), Mode: table.Asc},
	})

	for i, _ := range dataTable.Rows {
		row := dataTable.Rows[i] // because range is nondeterministic
		var cells []interface{} 

		for _, hdr := range dataTable.Headers {
			cells = append(cells, (*row)[hdr.(string)])
		}

		tbl.AppendRow(cells)
	}

	return tbl
}

func NewDataTable(headers []interface{}) *DataTable {
	tbl := &DataTable {
		Headers: headers,
	}

	tbl.Rows = make(map[int]*TableRow)

	return tbl
}
