package main

import (
	"fmt"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"

	. "github.com/upsilonproject/upsilon-cli/internal/runtimeconfig"
	"github.com/upsilonproject/upsilon-cli/internal/cmds"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

var rootCmd = &cobra.Command{
	Use:   "upsilon",
	Short: "The upsilon command.",
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(cmds.NodeCmd)
	cmds.NodeCmd.AddCommand(cmds.NodeListCmd)

	rootCmd.AddCommand(cmds.ConfigCmd)
	rootCmd.AddCommand(cmds.CmdMsgTail)
	rootCmd.AddCommand(cmds.CmdPing)
	rootCmd.AddCommand(cmds.CmdServices)

	cmds.CmdAmqp.AddCommand(cmds.CmdAmqpConnections)
	cmds.CmdAmqp.AddCommand(cmds.CmdAmqpInstall)
	rootCmd.AddCommand(cmds.CmdAmqp)

	rootCmd.AddCommand(cmds.CmdRequest)
	cmds.CmdRequest.AddCommand(cmds.CmdRequestUpdate)
	cmds.CmdRequest.AddCommand(cmds.CmdGitPull)
	cmds.CmdRequest.AddCommand(cmds.CmdRequestExecution)

	//rootCmd.AddCommand(cmds.CmdDrone)
	//cmds.CmdDrone.AddCommand(cmds.CmdLatestVersion)


	rootCmd.PersistentFlags().StringP("format", "f", "table", "output format")
	rootCmd.PersistentFlags().StringP("logLevel", "l", "info", "log level")
}

func initFlagValues() {
	RuntimeConfig.OutputFormat, _ = rootCmd.PersistentFlags().GetString("format")

	RuntimeConfig.LogLevel, _ = rootCmd.PersistentFlags().GetString("logLevel")
}

func initConfig() {
	initFlagValues()

	home, err := homedir.Dir()

	if err != nil {
		fmt.Println(err)
	}

	viper.AddConfigPath(home)
	viper.AddConfigPath("/etc/upsilon-cli")
	viper.SetConfigName("upsilon-cli")

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		log.Debugf("Using config file: %s", viper.ConfigFileUsed())
	}

	if err := viper.UnmarshalExact(&RuntimeConfig); err != nil {
		log.Warnf("Unmarshal config err: %v", err)
	}

	onConfigChanged() 
}

func onConfigChanged() {
	logLevel, _ := log.ParseLevel(RuntimeConfig.LogLevel)

	if logLevel != log.GetLevel() {
		log.Infof("Setting log level to: %v", logLevel)
		log.SetLevel(logLevel)
	}

	amqp.AmqpHost = RuntimeConfig.AmqpHost
	amqp.AmqpUser = RuntimeConfig.AmqpUser
	amqp.AmqpPass = RuntimeConfig.AmqpPass
	amqp.AmqpPort = RuntimeConfig.AmqpPort
}
