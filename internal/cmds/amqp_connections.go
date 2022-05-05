package cmds

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

func cmdAmqpConnections(cmd *cobra.Command, args []string) {
	log.Infof("Connections?!")
}

var CmdAmqpConnections = &cobra.Command {
	Use: "connections",
	Short: "Connections",
	Run: cmdAmqpConnections,
}
