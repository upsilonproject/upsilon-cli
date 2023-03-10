package cmds

import (
	"time"
	"github.com/spf13/cobra"
	"github.com/upsilonproject/upsilon-cli/internal/output"
	term "github.com/buger/goterm"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
	pb "github.com/upsilonproject/upsilon-cli/gen/amqpproto"
)

type SeenNodes map[string]*pb.Heartbeat;

func updateNodes(seenNodes *SeenNodes) {
	_, handler := amqp.ConsumeForever("Heartbeat", func(d amqp.Delivery) {
		d.Message.Ack(true)

		hb := &pb.Heartbeat{}

		amqp.Decode(d.Message.Body, hb)

		if _, found := (*seenNodes)[hb.Hostname]; !found {
			(*seenNodes)[hb.Hostname] = hb
		} else {
			(*seenNodes)[hb.Hostname].Version = hb.Version
		}
	});

	handler.Wait()
}

func nodeList(cmd *cobra.Command, args []string) {
	seenNodes := make(SeenNodes, 0)

	go updateNodes(&seenNodes)

	for {
		term.Clear()
		term.MoveCursor(1, 1)

		tbl := getNodeTable();

		for _, hb := range seenNodes {
			row := make(output.TableRow)
			row["identifier"] = hb.Hostname
			row["type"] = hb.Type
			row["version"] = hb.Version
			tbl.Append(&row)
		}

		term.Println(output.Format(tbl))
		term.Flush()
		time.Sleep(time.Second)
	}
}

func getNodeTable() *output.DataTable {
	return output.NewDataTable([]interface{} {
		"identifier",
		"type",
		"version",
	})
}

var CmdListenNodeHeartbeats = &cobra.Command{
	Use: "heartbeats",
	Run: nodeList,
}
