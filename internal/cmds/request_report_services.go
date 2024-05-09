package cmds

import (
	"github.com/spf13/cobra"
	pb "github.com/upsilonproject/upsilon-cli/gen/amqpproto"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
	"github.com/upsilonproject/upsilon-cli/internal/output"
	text "github.com/jedib0t/go-pretty/text"
)

func runReport(cmd *cobra.Command, args []string) {
	includeGood, _ := cmd.PersistentFlags().GetBool("includeGood")

	req := pb.ReportRequest{
		IncludeGood: includeGood,
	}

	consumer, handler := amqp.ConsumeSingle("ReportResponse", func(d amqp.Delivery) {
		d.Message.Ack(true)

		res := pb.ReportResponse{}

		amqp.Decode(d.Message.Body, &res)

		tbl := reportToDataTable(&res)

		output.Prepare(
			tbl,
			"All services are " + text.FgGreen.Sprintf("GOOD") + ".",
		)

		if output.IsPrettyTable() {
			output.PrettyTableRelativeTimestamps(5)
			output.PrettyTableAddKarma(4)
			output.PrettyTableSortBy("karma", "node", "service")
			output.PrettyTableHeaders(getHeaders())
		}

		output.PrintPrepared()
	});

	consumer.Wait()

	amqp.PublishPb(&req)

	handler.Wait()
}

func reportToDataTable(res *pb.ReportResponse) *output.DataTable { 
	var headers []interface{}

	for _, col := range res.Columns {
		headers = append(headers, col.Header)
	}

	tbl := output.NewDataTable(headers)

	for _, row := range res.Rows {
		cells := &output.TableRow{}

		for _, hdr := range headers {
			(*cells)[hdr.(string)] = row.Cells[hdr.(string)]
		}

		tbl.Append(cells)
	}

	return tbl
}

func getSortField() string {
	sort, _ := CmdServicesReport.PersistentFlags().GetString("sortAsc")

	if sort == "" {
		sort, _ = CmdServicesReport.PersistentFlags().GetString("sortDec")
	}

	return sort
}

func getHeaders() []string {
	headers, _ := CmdServicesReport.PersistentFlags().GetStringArray("headers")

	return headers
}

func init() {
	CmdServicesReport.PersistentFlags().StringP("sortAsc", "s", "", "Sort Ascending")
	CmdServicesReport.PersistentFlags().StringP("sortDec", "S", "", "Sort Descending")
	CmdServicesReport.PersistentFlags().StringArrayP("headers", "H", []string {}, "Headers")
	CmdServicesReport.PersistentFlags().BoolP("includeGood", "g", false, "Include good services")
	CmdServicesReport.Run = runReport
}

var CmdServicesReport = &cobra.Command{
	Use:   "services",
	Aliases: []string { "svc" },
	Short: "Services Commands",
}


