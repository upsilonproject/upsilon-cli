package cmds

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"

	//"google.golang.org/protobuf/proto"

	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

func msgTail(cmd *cobra.Command, args []string) {
	amqp.Consume("*", func(d amqp.Delivery) {
		d.Message.Ack(true)

		msgType := d.Message.Headers["Upsilon-Msg-Type"]

		log.Infof("Delivery %v %+v", msgType, string(d.Message.Body));
		//len(d.Message.Body), string(d.Message.Body))
	})
}

var CmdMsgTail = &cobra.Command{
	Use:   "msgtail",
	Short: "Message Tail",
	Run: msgTail,
}

