package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/strebul/gogy/component"
	"github.com/strebul/gogy/model"
	"strings"
	"time"
)

var GogIdCmd = &cobra.Command{
	Use:   "id [string]",
	Short: "Searching log by id",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := component.LoadConfig(gogConfigFile)
		if err != nil {
			return err
		}

		ids := strings.Join(args, "\",\"")

		query := fmt.Sprintf(`_id:("%s")`, ids)

		request := model.Request{
			Query:     query,
			TimeStart: time.Now().Add(-time.Duration(gogDuration) * time.Hour),
			TimeEnd:   time.Now(),
			Size:      10,
			Order:     "desc",
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
			//			fmt.Println(log.Source["exception"])
		}

		return nil
	},
}

var gogIdDuration int
var gogIdConfigFile string

func init() {
	GogIdCmd.Flags().IntVarP(&gogIdDuration, "duration", "d", 24, "")
	GogIdCmd.Flags().StringVarP(&gogIdConfigFile, "config", "c", component.DefaultConfig, "")
}
