package command

import (
	"github.com/spf13/cobra"
	"github.com/strebul/gogy/component"
	"github.com/strebul/gogy/component/worker/pastabank"
)

var workerDuration int

var Worker = &cobra.Command{
	Use:   "worker [name of worker]",
	Short: "Worker analysis",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := component.LoadConfig(ConfigFile)
		if err != nil {
			return err
		}

		client := component.Client{
			Host:     config.ReadString("logstash.host"),
			Login:    config.ReadString("logstash.login"),
			Password: config.ReadString("logstash.password"),
		}

		pastabank.Report(client, workerDuration)

		return nil
	},
}

func init() {
	Worker.Flags().IntVarP(&workerDuration, "duration", "d", 24, "")
}
