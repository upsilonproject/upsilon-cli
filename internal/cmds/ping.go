package cmds

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	pb "github.com/upsilonproject/upsilon-cli/gen/amqpproto"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
	"time"
)

func runPing(cmd *cobra.Command, args []string) {
	c, err := amqp.GetChannel()

	if err != nil {
		log.Warnf("Could not get chan: %s", err)
		return
	}

	req := pb.PingRequest{}

	go amqp.Consume(c, "PingResponse", func(d amqp.Delivery) {
		d.Message.Ack(true)

		res := pb.PingResponse{}
		
		amqp.Decode(d.Message.Body, &res)

		log.Infof("Ping reply: %+v", res.Hostname)
	});

	log.Infof("consuming")
	
	time.Sleep(time.Second * 1)

	amqp.PublishPb(c, req)

	time.Sleep(time.Second * 120)
}

var CmdPing = &cobra.Command{
	Use:   "ping",
	Short: "Ping everything",
	Run: runPing,
}


