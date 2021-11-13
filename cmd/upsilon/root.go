package main

import (
	"fmt"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"

	. "github.com/upsilonproject/upsilon-cli/internal/runtimeconfig"
	"github.com/upsilonproject/upsilon-cli/internal/cmds"
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

	rootCmd.PersistentFlags().StringP("format", "f", "table", "output format")
}

func initFlagValues() {
	RuntimeConfig.OutputFormat, _ = rootCmd.PersistentFlags().GetString("format")
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
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.UnmarshalExact(&RuntimeConfig); err != nil {
		log.Warnf("Unmarshal config err: %v", err)
	}
}
