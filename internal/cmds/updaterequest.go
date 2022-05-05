package cmds

import (
	"github.com/spf13/cobra"
	pb "github.com/upsilonproject/upsilon-cli/gen/amqpproto"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

func runUpdateRequest(cmd *cobra.Command, args []string) {
	req := pb.UpdateRequest{}

	amqp.PublishPb(&req)
}

var CmdUpdateRequest = &cobra.Command{
	Use:   "update-request",
	Short: "Send update request",
	Run: runUpdateRequest,
}


