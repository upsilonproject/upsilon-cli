package cmds

import (
	"github.com/spf13/cobra"
)

var CmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Status",
	Aliases: []string { "st" },
	Run: runReport,
}


