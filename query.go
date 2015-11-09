package main

import (
    "gogy/component"
    "time"
    "gogy/model"
    "flag"
    "strings"
    "fmt"
)

func main() {
    // Command options
    configFile      := flag.String("config", "", "Config file")
    level           := flag.String("log-level", "", "Log level")
    size            := flag.Int("size", 100, "Size")
    lastHours       := flag.Int("lastHours", 72, "Logs of the last hourse")
    scriptId        := flag.String("scriptId", "", "Script id")
    sessionId       := flag.String("sessionId", "", "Session id")
    message         := flag.String("message", "", "Message")
    flag.Parse()

    // Build query
    var query string

    if len(*scriptId) > 0 {
        query += fmt.Sprintf(`script-id: "%s"`, *scriptId)
    }

    if len(*sessionId) > 0 {
        query += fmt.Sprintf(`sessionId: "%s"`, *sessionId)
    }

    if len(*message) > 0 {
        query += fmt.Sprintf(`message: "%s"`, *message)
    }

    if len(*level) > 0 {
        logLevel := convertLevelToQuery(*level)
        if (len(query) > 0) {
            query += " AND (" + logLevel + ")"
        } else {
            query = logLevel
        }
    }

    if args := flag.Args(); cap(args) > 0 {
        if (len(query) > 0) {
            query += " AND"
        }
        query += " (" + strings.Join(flag.Args(), " AND ") + ")"
    }

    if len(query) == 0 {
        query = "*"
    }

    config := component.Config{Filename: *configFile}
    config.Init()

    endTime := time.Now()
    startTime := endTime.Add(- time.Duration(*lastHours) * time.Hour);

    request := model.Request{
        Query: query,
        TimeStart: startTime.UnixNano() / int64(time.Millisecond),
        TimeEnd: endTime.UnixNano() / int64(time.Millisecond),
        Size: *size,
    }

    client := component.Client{
        Host: config.Logstash.Host,
        Login: config.Logstash.Login,
        Password: config.Logstash.Password,
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

        if (cap(include) > 0) {
            query += "log-level: (\"" + strings.Join(include, "\",\"") + "\")"
        }

        if (cap(exclude) > 0) {
            if (len(query) > 0) {
                query += " AND"
            }
            query += " NOT log-level: (\"" + strings.Join(exclude, "\",\"") + "\")"
        }
    }

    return strings.Trim(query, " ")
}
