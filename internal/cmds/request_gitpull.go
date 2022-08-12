package cmds

import (
	"github.com/spf13/cobra"
	pb "github.com/upsilonproject/upsilon-cli/gen/amqpproto"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

func runGitPull(cmd *cobra.Command, args []string) {
	req := pb.GitPullRequest{}
	req.GitUrl = "ssh://git@upsilon:/opt/upsilon-config/"
	amqp.PublishPb(req)
}

var CmdGitPull = &cobra.Command{
	Use:   "gitpull",
	Short: "Trigger drones to pull git repos",
	Run: runGitPull,
}


