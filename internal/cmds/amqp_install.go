package cmds

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	commonamqp "github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

func cmdInstall(cmd *cobra.Command, args []string) {
	log.Infof("Installing")

	c, err := commonamqp.GetChannel("install")

	if err != nil {
		log.Fatal(err)
		return
	}

	err = c.ExchangeDeclare(
		"ex_upsilon",
		"topic",
		true, // durable
		false, // autodelete
		false, // internal 
		false, // nowait
		nil,
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Infof("Exchange declared")
}

var CmdAmqpInstall = &cobra.Command{
	Use:   "install",
	Short: "Install",
	Run: cmdInstall,
}
