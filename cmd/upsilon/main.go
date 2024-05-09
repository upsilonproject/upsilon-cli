package main

import (
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
	log "github.com/sirupsen/logrus"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Root command fatal error: %v", err);
	}
}

func main() {
	amqp.ConnectionIdentifier = "upsilon-cli"
	Execute()
}
