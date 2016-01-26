package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/strebul/gogy/component"
	"github.com/strebul/gogy/model"
	"github.com/strebul/gogy/model/log"
	"strings"
	"time"
)

var gogLevel string
var gogSize int
var gogDuration int
var gogScriptId string
var gogSessionId string
var gogMessage string
var gogObject string

var GogCmd = &cobra.Command{
	Use:   "query [arguments to search]",
	Short: "Searching logs by query",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := component.LoadConfig(ConfigFile)
		if err != nil {
			return err
		}

		query := buildQuery(args)

		request := model.Request{
			Query:     query,
			TimeStart: time.Now().Add(-time.Duration(gogDuration) * time.Hour),
			TimeEnd:   time.Now(),
			Size:      gogSize,
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
		decorator.DecorateList(list, true)

		return nil
	},
}

func init() {
	GogCmd.Flags().StringVarP(&gogLevel, "log-level", "l", "", "~debug or warning,error")
	GogCmd.Flags().IntVarP(&gogSize, "size", "s", 100, "")
	GogCmd.Flags().IntVarP(&gogDuration, "duration", "d", 24, "")
	GogCmd.Flags().StringVarP(&gogScriptId, "script-id", "", "", "")
	GogCmd.Flags().StringVarP(&gogSessionId, "session-id", "", "", "")
	GogCmd.Flags().StringVarP(&gogMessage, "message", "m", "", "")
	GogCmd.Flags().StringVarP(&gogObject, "object", "o", "", "")
}

func buildQuery(args []string) string {
	var query string

	if len(gogScriptId) > 0 {
		if len(query) > 0 {
			query += " AND "
		}
		query += fmt.Sprintf(`script-id: "%s"`, gogScriptId)
	}
	if len(gogSessionId) > 0 {
		if len(query) > 0 {
			query += " AND "
		}
		query += fmt.Sprintf(`sessionId: "%s"`, gogSessionId)
	}
	if len(gogMessage) > 0 {
		if len(query) > 0 {
			query += " AND "
		}
		query += fmt.Sprintf(`message: "%s"`, gogMessage)
	}
	if len(gogObject) > 0 {
		if len(query) > 0 {
			query += " AND "
		}
		object := strings.Replace(gogObject, "\\", ".", -1)
		if strings.HasPrefix(object, ".") {
			object = object[1:len(object)]
		}
		query += fmt.Sprintf(`object: "%s"`, object)
	}

	if len(gogLevel) > 0 {
		logLevel := levelCondition(gogLevel)
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

	return strings.Trim(query, " ")
}

func levelCondition(str string) string {
	var query string

	var levels []string

	levels = append(
		levels,
		log.DEBUG_LONG_STRING,
		log.INFO_LONG_STRING,
		log.NOTICE_LONG_STRING,
		log.WARNING_LONG_STRING,
		log.ERROR_LONG_STRING,
		log.CRITICAL_LONG_STRING,
		log.ALERT_LONG_STRING,
		log.EMERGENCY_LONG_STRING,
	)

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

		if cap(include) == 1 {
			var list []string
			level := include[0]

			s := false
			for _, val := range levels {
				if level == val {
					s = true
				}
				if s {
					list = append(list, val)
				}
			}

			include = list
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

	return query
}
