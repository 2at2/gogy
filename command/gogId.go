package command

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/strebul/gogy/component"
	"github.com/strebul/gogy/model"
	"time"
)

var GogIdCmd = &cobra.Command{
	Use:   "id [string]",
	Short: "Searching log by id",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := component.LoadConfig(configFile)
		if err != nil {
			return err
		}

		if len(id) == 0 {
			return errors.New("Set id")
		}

		query := fmt.Sprintf(`_id:"%s"`, id)

		request := model.Request{
			Query:     query,
			TimeStart: time.Now().Add(-time.Duration(duration) * time.Hour),
			TimeEnd:   time.Now(),
			Size:      10,
		}

		client := component.Client{
			Host:     config.ReadString("logstash.host"),
			Login:    config.ReadString("logstash.login"),
			Password: config.ReadString("logstash.password"),
		}

		list := client.FindLogs(request)

		decorator := component.Decorator{}

		decorator.DecorateRequest(request)

		for _, log := range list {
			decorator.DecorateDetails(log)
		}

		return nil
	},
}

var id string

func init() {
	GogIdCmd.Flags().StringVarP(&id, "id", "i", "", "")
	GogIdCmd.Flags().IntVarP(&duration, "duration", "d", 24, "")
}
