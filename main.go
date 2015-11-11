package main

import (
	"github.com/spf13/cobra"
	"github.com/strebul/gogy/command"
)

func main() {
	var RootCmd = &cobra.Command{Use: "gg"}

	RootCmd.AddCommand(command.GogCmd)
	RootCmd.AddCommand(command.GogIdCmd)
	RootCmd.Execute()
}
