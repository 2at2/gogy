package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/strebul/gogy/component"
	"github.com/strebul/gogy/model"
	"gopkg.in/yaml.v2"
	"strings"
	"time"
)

var GogCmd = &cobra.Command{
	Use:   "query [arguments to search]",
	Short: "Searching logs by query",
	Run: func(cmd *cobra.Command, args []string) {
		executeGog(args)
	},
}

var level string
var size int
var duration int
var scriptId string
var sessionId string
var message string
var configFile string

func init() {
	GogCmd.Flags().StringVarP(&level, "log-level", "l", "", "~debug or warning,error")
	GogCmd.Flags().IntVarP(&size, "size", "s", 100, "")
	GogCmd.Flags().IntVarP(&duration, "duration", "d", 24, "")
	GogCmd.Flags().StringVarP(&scriptId, "script-id", "", "", "")
	GogCmd.Flags().StringVarP(&sessionId, "session-id", "", "", "")
	GogCmd.Flags().StringVarP(&message, "message", "m", "", "")
	GogCmd.Flags().StringVarP(&configFile, "config", "c", "", "")
}

func executeGog(args []string) {
	var config component.Reader

	if bytes, err := component.LoadConfig(configFile); err == nil {
		yaml.Unmarshal(bytes, &config)
	} else {
		panic(err)
	}

	// Build query
	var query string

	if len(scriptId) > 0 {
		query += fmt.Sprintf(`script-id: "%s"`, scriptId)
	}
	if len(sessionId) > 0 {
		query += fmt.Sprintf(`sessionId: "%s"`, sessionId)
	}
	if len(message) > 0 {
		query += fmt.Sprintf(`message: "%s"`, message)
	}
	if len(level) > 0 {
		logLevel := convertLevelToQuery(level)
		if len(query) > 0 {
			query += " AND (" + logLevel + ")"
		} else {
			query = logLevel
		}
	}
	if len(args) > 0 {
		if len(query) > 0 {
			query += " AND"
		}
		query += " (" + strings.Join(args, " AND ") + ")"
	}
	if len(query) == 0 {
		query = "*"
	}

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
}

func convertLevelToQuery(str string) string {
	var query string

	if len(str) > 0 {
		var include []string
		var exclude []string
		for _, val := range strings.Split(str, ",") {
			if strings.HasPrefix(val, "~") {
				exclude = append(exclude, val[1:len(val)])
			} else {
				include = append(include, val)
			}
		}

		if cap(include) > 0 {
			query += "log-level: (\"" + strings.Join(include, "\",\"") + "\")"
		}

		if cap(exclude) > 0 {
			if len(query) > 0 {
				query += " AND"
			}
			query += " NOT log-level: (\"" + strings.Join(exclude, "\",\"") + "\")"
		}
	}

	return strings.Trim(query, " ")
}
