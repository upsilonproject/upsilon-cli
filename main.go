package main

import (
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
	"github.com/upsilonproject/upsilon-cli/internal/cmds"
	log "github.com/sirupsen/logrus"
)

func Execute() {
	if err := cmds.RootCmd.Execute(); err != nil {
		log.Fatalf("Root command fatal error: %v", err);
	}
}

func main() {
	amqp.ConnectionIdentifier = "upsilon-cli"
	Execute()
}
