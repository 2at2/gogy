package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/strebul/gogy/component"
	"github.com/strebul/gogy/model"
	"gopkg.in/yaml.v2"
	"time"
)

var GogIdCmd = &cobra.Command{
	Use:   "query [arguments to search]",
	Short: "Searching logs by query",
	Run: func(cmd *cobra.Command, args []string) {
		var config component.Reader

		if bytes, err := component.LoadConfig(configFile); err == nil {
			yaml.Unmarshal(bytes, &config)
		} else {
			panic(err)
		}

		if len(id) == 0 {
			fmt.Println("Set id")
			return
		}

		query := fmt.Sprintf(`_id:"%d"`, id)

		endTime := time.Now()
		startTime := endTime.Add(-(time.Duration(duration) * time.Hour))

		request := model.Request{
			Query:     query,
			TimeStart: startTime.UnixNano() / int64(time.Millisecond),
			TimeEnd:   endTime.UnixNano() / int64(time.Millisecond),
			Size:      size,
		}

		client := component.Client{
			Host:     config.ReadString("logstash.host"),
			Login:    config.ReadString("logstash.login"),
			Password: config.ReadString("logstash.password"),
		}
		list := client.FindLogs(request)

		decorator := component.Decorator{}

		decorator.DecorateRequest(request)
		decorator.DecorateList(list, true)
	},
}

var id string
var size int
var duration int
var configFile string

func init() {
	GogIdCmd.Flags().StringVarP(&id, "id", "i", "", "")
	GogIdCmd.Flags().IntVarP(&size, "size", "s", 10, "")
	GogIdCmd.Flags().IntVarP(&duration, "duration", "d", 24, "")
	GogIdCmd.Flags().StringVarP(&configFile, "config", "c", "", "")
}