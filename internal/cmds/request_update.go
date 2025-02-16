package cmds

import (
	"github.com/spf13/cobra"
	pb "github.com/upsilonproject/upsilon-gocommon/pkg/amqpproto"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

func runRequestUpdate(cmd *cobra.Command, args []string) {
	req := pb.UpdateRequest{}

	amqp.PublishPb(&req)
}

var CmdRequestUpdate = &cobra.Command{
	Use:   "update",
	Short: "Send update request",
	Run: runRequestUpdate,
}


