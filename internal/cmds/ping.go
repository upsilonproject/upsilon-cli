package cmds

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	pb "github.com/upsilonproject/upsilon-cli/gen/amqpproto"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
	"time"
)

func runPing(cmd *cobra.Command, args []string) {
	req := pb.PingRequest{}

	go amqp.Consume("PingResponse", func(d amqp.Delivery) {
		d.Message.Ack(true)

		res := pb.PingResponse{}
		
		amqp.Decode(d.Message.Body, &res)

		log.Infof("Ping reply: %+v", res.Hostname)
	});

	// The AMQP Server seems to need a moment to create the consumer.
	time.Sleep(time.Second * 1)

	amqp.PublishPb(&req)

	time.Sleep(time.Second * 120)
}

var CmdPing = &cobra.Command{
	Use:   "ping",
	Short: "Ping everything",
	Run: runPing,
}


