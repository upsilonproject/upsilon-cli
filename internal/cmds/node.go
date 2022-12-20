package cmds

import (
	"time"
	"github.com/spf13/cobra"
	"github.com/upsilonproject/upsilon-cli/internal/output"
	term "github.com/buger/goterm"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
	pb "github.com/upsilonproject/upsilon-cli/gen/amqpproto"
)

func updateNodes(tbl *output.DataTable) {
	//seenNodes := make(map[string]int, 0)

	_, handler := amqp.ConsumeForever("Heartbeat", func(d amqp.Delivery) {
		d.Message.Ack(true)

		hb := pb.Heartbeat{}

		amqp.Decode(d.Message.Body, &hb)

	//	if node, found := seenNodes[hb.Hostname]; !found {
	//		seenNodes[hb.Hostname] = hb.Version
	//	}

		row := make(output.TableRow)
		row["hostname"] = hb.Hostname
		row["version"] = hb.Version
		tbl.Append(&row)
	});

	handler.Wait()
}

func nodeList(cmd *cobra.Command, args []string) {
	tbl := getNodeTable();

	go updateNodes(tbl)

	for {
		term.Clear()
		term.MoveCursor(1, 1)
		term.Println(output.Format(tbl))
		term.Flush()
		time.Sleep(time.Second)
	}
}

func getNodeTable() *output.DataTable {
	return output.NewDataTable([]string {
		"identifier",
		"type",
		"version",
	})
}

var NodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Node details",
}

var NodeListCmd = &cobra.Command{
	Use: "list",
	Run: nodeList,
}
