package main

import (
	"fmt"
	"os"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	amqp.ConnectionIdentifier = "upsilon-cli"
	Execute()
}
