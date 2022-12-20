package cmds

import (
	"github.com/spf13/cobra"
	pb "github.com/upsilonproject/upsilon-cli/gen/amqpproto"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

func runGitPull(cmd *cobra.Command, args []string) {
	alias := cmd.Flags().Lookup("alias").Value.String()

	req := pb.GitPullRequest{}
	req.GitUrlAlias = alias
	amqp.PublishPb(req)
}

func init() {
	CmdGitPull.Flags().StringP("alias", "", "fabric-config", "Git URL Alias")
}

var CmdGitPull = &cobra.Command{
	Use:   "gitpull",
	Short: "Trigger drones to pull git repos",
	Run: runGitPull,
}


