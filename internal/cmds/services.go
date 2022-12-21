package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	pb "github.com/upsilonproject/upsilon-cli/gen/amqpproto"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
	"github.com/upsilonproject/upsilon-cli/internal/output"
)

func runReport(cmd *cobra.Command, args []string) {
	req := pb.ReportRequest{}

	consumer, handler := amqp.ConsumeSingle("ReportResponse", func(d amqp.Delivery) {
		d.Message.Ack(true)

		res := &pb.ReportResponse{}
		
		amqp.Decode(d.Message.Body, res)

		tbl := rpt2tbl(res)

		fmt.Printf(output.Format(tbl))
	});

	consumer.Wait()

	amqp.PublishPb(&req)

	handler.Wait()
}

func rpt2tbl(res *pb.ReportResponse) *output.DataTable { 
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

func init() {
	CmdServices.AddCommand(cmdServicesReport)
}

var cmdServicesReport = &cobra.Command{
	Use: 	"report",
	Short: "report",
	Run: runReport,
}

var CmdServices = &cobra.Command{
	Use:   "services",
	Short: "Services Commands",
}


