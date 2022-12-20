package cmds

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	pb "github.com/upsilonproject/upsilon-cli/gen/amqpproto"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

func runPing(cmd *cobra.Command, args []string) {
	req := pb.PingRequest{}

	consumer, handler := amqp.ConsumeForever("PingResponse", func(d amqp.Delivery) {
		d.Message.Ack(true)

		res := pb.PingResponse{}
		
		amqp.Decode(d.Message.Body, &res)

		log.Infof("Ping reply: %+v", res.Hostname)
	});

	consumer.Wait()

	amqp.PublishPb(&req)

	handler.Wait()
}

var CmdPing = &cobra.Command{
	Use:   "ping",
	Short: "Ping everything",
	Run: runPing,
}


