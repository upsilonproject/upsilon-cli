package cmds

import (
	"github.com/spf13/cobra"
)

var CmdRequest = &cobra.Command {
	Use: "request",
	Aliases: []string {"req", "get"},
	Short: "Send requests to the AMQP network",
}

