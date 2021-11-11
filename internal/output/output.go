package output

import (
	"encoding/json"
	prettytable "github.com/tatsushid/go-prettytable"
	. "github.com/upsilonproject/upsilon-cli/internal/runtimeconfig"

	log "github.com/sirupsen/logrus"
)

type TableRow = map[string]string
type DataTable struct {
	Headers []string
	Rows map[int]TableRow
}

func (dt *DataTable) Append(row TableRow) {
	idx := len(dt.Rows)
	dt.Rows[idx] = row
}

func Format(rows DataTable) string {
	output := "format: "

	switch RuntimeConfig.OutputFormat {
	case "table":
		return formatOutputTable(rows)
	case "json":
		return formatOutputJson(rows)
	}

	return output
}

func formatOutputJson(rows DataTable) string {
	ret, _ := json.Marshal(rows)

	return string(ret)
}

func formatOutputTable(dataTable DataTable) string {
	var columns = make([]prettytable.Column, 0)

	for _, header := range dataTable.Headers {
		columns = append(columns, prettytable.Column{
			Header: header,
		})
	}

	prettyTable, err := prettytable.NewTable(columns...)

	if err != nil {
		log.Warnf("%v", err)
	}

	prettyTable.Separator = " | "

	for _, row := range dataTable.Rows {
		for _, cel := range row {
			prettyTable.AddRow(cel)
		}
	}

	return prettyTable.String()
}
