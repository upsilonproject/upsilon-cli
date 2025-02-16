package cmds

import (
	"fmt"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"

	. "github.com/upsilonproject/upsilon-cli/internal/runtimeconfig"
	"github.com/upsilonproject/upsilon-gocommon/pkg/amqp"
)

var RootCmd = &cobra.Command{
	Use:   "upsilon",
	Short: "The upsilon command.",
}

var listenCmd = &cobra.Command {
	Use: "listen",
	Aliases: []string{ "watch", "w" },
	Short: "Listen to messages",
}

var cmdReport = &cobra.Command {
	Use: "report",
	Short: "Reports from the custodian.",
	Aliases: []string{ "rpt" },
}

func init() {
	cobra.OnInitialize(initConfig)

	listenCmd.AddCommand(CmdListenNodeHeartbeats)
	listenCmd.AddCommand(CmdMsgTail)

	RootCmd.AddCommand(ConfigCmd)
	RootCmd.AddCommand(CmdPing)
	RootCmd.AddCommand(CmdStatus)
	RootCmd.AddCommand(listenCmd)

	CmdAmqp.AddCommand(CmdAmqpConnections)
	CmdAmqp.AddCommand(CmdAmqpInstall)
	RootCmd.AddCommand(CmdAmqp)

	cmdReport.AddCommand(CmdServicesReport)

	RootCmd.AddCommand(CmdRequest)
	CmdRequest.AddCommand(cmdReport)
	CmdRequest.AddCommand(CmdRequestUpdate)
	CmdRequest.AddCommand(CmdGitPull)
	CmdRequest.AddCommand(CmdRequestExecution)

	//RootCmd.AddCommand(cmds.CmdDrone)
	//cmds.CmdDrone.AddCommand(cmds.CmdLatestVersion)


	RootCmd.PersistentFlags().StringP("format", "f", "table", "output format")
	RootCmd.PersistentFlags().StringP("logLevel", "l", "info", "log level")
}

func initFlagValues() {
	RuntimeConfig.OutputFormat, _ = RootCmd.PersistentFlags().GetString("format")

	RuntimeConfig.LogLevel, _ = RootCmd.PersistentFlags().GetString("logLevel")
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
