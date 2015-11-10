package command

import (
    "github.com/spf13/cobra"
    "fmt"
    "strings"
    "gogy/component"
    "time"
    "gogy/model"
)

var QueryCmd = &cobra.Command{
    Use:   "query [arguments to search]",
    Short: "Searching logs by query",
    Run: func(cmd *cobra.Command, args []string) {
        execute(args)
    },
}

var level string
var size int
var duration int
var scriptId string
var sessionId string
var message string
var env string

func init() {
    QueryCmd.Flags().StringVarP(&level, "log-level", "l", "", "~debug or warning,error")
    QueryCmd.Flags().IntVarP(&size, "size", "s", 100, "")
    QueryCmd.Flags().IntVarP(&duration, "duration", "d", 24, "")
    QueryCmd.Flags().StringVarP(&scriptId, "script-id", "", "", "")
    QueryCmd.Flags().StringVarP(&sessionId, "session-id", "", "", "")
    QueryCmd.Flags().StringVarP(&message, "message", "m", "", "")
    QueryCmd.Flags().StringVarP(&env, "env", "", "live", "dev, rel, live")
}

func execute(args []string) {
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

    config := component.Config{Env: env}
    config.InitConfigFile("params")

    fmt.Println(config.Source)

    endTime := time.Now()
    startTime := endTime.Add(-(time.Duration(duration) * time.Hour))

    request := model.Request{
        Query:     query,
        TimeStart: startTime.UnixNano() / int64(time.Millisecond),
        TimeEnd:   endTime.UnixNano() / int64(time.Millisecond),
        Size:      size,
    }

    client := component.Client{
        Host:     config.Source["logstash.host"].(string),
        Login:    config.Source["logstash.login"].(string),
        Password: config.Source["logstash.password"].(string),
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
