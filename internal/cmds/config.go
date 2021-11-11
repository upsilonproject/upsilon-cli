package cmds

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	. "github.com/upsilonproject/upsilon-cli/internal/runtimeconfig"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Get config",
	Run: func (cmd *cobra.Command, args []string) {
		log.Infof("%+v", RuntimeConfig)
	},
}

