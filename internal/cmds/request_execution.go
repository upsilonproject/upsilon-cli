package cmds

import (
	"github.com/spf13/cobra"
	pb "github.com/upsilonproject/upsilon-cli/gen/amqpproto"
	log "github.com/sirupsen/logrus"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

func runReqExec(cmd *cobra.Command, args []string) {
	req := pb.ExecutionRequest{
		Hostname: cmd.Flags().Lookup("hostname").Value.String(),
		CommandName: cmd.Flags().Lookup("command").Value.String(),
	}

	consumer, handler := amqp.ConsumeSingle("ExecutionResult", func(d amqp.Delivery) {
		d.Message.Ack(true)

		execResult := pb.ExecutionResult{}

		amqp.Decode(d.Message.Body, &execResult)

		log.Infof("%v", execResult)
	})

	consumer.Wait()

	amqp.PublishPb(req)

	handler.Wait()
}

func init() {
	CmdRequestExecution.Flags().StringP("hostname", "n", "", "Drone Hostname")
	CmdRequestExecution.MarkFlagRequired("hostname")
	CmdRequestExecution.Flags().StringP("command", "c", "", "Command name")
	CmdRequestExecution.MarkFlagRequired("command")
}


var CmdRequestExecution = &cobra.Command{
	Use:   "execution",
	Aliases: []string {"exec"},
	Short: "Request drones execute a command",
	Run: runReqExec,
}


