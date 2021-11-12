package cmds

import (
	"github.com/spf13/cobra"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
	amqp2 "github.com/streadway/amqp"
	log "github.com/sirupsen/logrus"
	"time"
)

func msgTail(cmd *cobra.Command, args []string) {
	c, err := amqp.GetChannel("upsilon-cli")

	if err != nil {
		log.Warnf("%v", err)
		return
	}

	q, err := c.QueueDeclare(
		"asdf",
		false, // durable
		false, // delete when unused
		true, // exclusive
		true, // nowait
		nil, // args
	)

	if err != nil {
		log.Warnf("%v", err)
		return
	}

	err = c.QueueBind(
		q.Name, 
		"*", // key
		"ex_upsilon",
		true, // nowait
		nil, // args
	)

	if err != nil {
		log.Warnf("%v", err)
		return
	}

	var done chan error;

		deliveries, err := c.Consume(
			q.Name, // name
			"tag",      // consumerTag,
			false,      // noAck
			false,      // exclusive
			false,      // noLocal
			false,      // noWait
			nil,        // arguments
		)

		if err != nil {
			log.Warnf("%v", err)
			return
		}

		go consumeDeliveries(deliveries, done)	

	for {
		time.Sleep(10 * time.Second)
	}
}

func consumeDeliveries(deliveries <-chan amqp2.Delivery, done chan error) {
	log.Infof("consuming")

	for d := range deliveries {
		log.Infof(
			"msg (%dB): %q",
			len(d.Body),
			d.Body,
		)
		d.Ack(false)
	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}

var CmdMsgTail = &cobra.Command{
	Use:   "msgtail",
	Short: "Message Tail",
	Run: msgTail,
}

