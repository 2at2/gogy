package main

import (
	"github.com/spf13/cobra"
	"github.com/strebul/gogy/command"
)

var ConfFile string
var Verbose bool

func main() {
	var RootCmd = &cobra.Command{Use: "gg"}

	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().StringVarP(&ConfFile, "config", "c", "", "configuration file")

	RootCmd.AddCommand(command.FrequencyCode)
	RootCmd.AddCommand(command.GogCmd)
	RootCmd.AddCommand(command.GogIdCmd)
	RootCmd.Execute()
}
