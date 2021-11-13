package cmds

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"

	//"google.golang.org/protobuf/proto"

	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

func msgTail(cmd *cobra.Command, args []string) {
	c, err := amqp.GetChannel()

	if err != nil {
		log.Warnf("Could not get chan: %s", err)
		return
	}

	amqp.Consume(c, "*", func(d amqp.Delivery) {
		d.Message.Ack(true)

		log.Infof("Delivery %+v", string(d.Message.Body));
		//len(d.Message.Body), string(d.Message.Body))
	})
}

var CmdMsgTail = &cobra.Command{
	Use:   "msgtail",
	Short: "Message Tail",
	Run: msgTail,
}

