package main

import (
	"github.com/spf13/cobra"
	"gogy/command"
)

var env string

func main() {
	var rootCmd = &cobra.Command{Use: "gg"}

	rootCmd.AddCommand(command.QueryCmd)
	rootCmd.Execute()

}
