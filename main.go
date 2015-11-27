package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/strebul/gogy/command"
)

func main() {
	var RootCmd = &cobra.Command{Use: "gg"}

	RootCmd.PersistentFlags().BoolVarP(&command.Verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().StringVarP(&command.ConfigFile, "config", "c", "gg.yml", "configuration file")

	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.SetDefault("config", "NAME HERE <EMAIL ADDRESS>")

	RootCmd.AddCommand(command.FrequencyCode)
	RootCmd.AddCommand(command.GogCmd)
	RootCmd.AddCommand(command.GogIdCmd)
	RootCmd.AddCommand(command.Worker)
	RootCmd.Execute()
}
