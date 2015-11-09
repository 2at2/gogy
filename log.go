package main

import (
    "gogy/component"
    "time"
    "gogy/model"
    "flag"
    "fmt"
    "os"
)

func main() {
    configFile  := flag.String("config", "", "Config file")
    lastHours   := flag.Int("lastHours", 72, "Logs of the last hourse")
    size        := flag.Int("size", 10, "Size")
    flag.Parse()

    id := flag.Arg(0)

    if len(id) == 0 {
        fmt.Println("Argument `id` missed")
        os.Exit(0)
    }

    query := fmt.Sprintf(`_id:"%s"`, id)

    endTime := time.Now()
    startTime := endTime.Add(- time.Duration(*lastHours) * time.Hour);

    request := model.Request{
        Query: query,
        TimeStart: startTime.UnixNano() / int64(time.Millisecond),
        TimeEnd: endTime.UnixNano() / int64(time.Millisecond),
        Size: *size,
    }

    config := component.Config{Filename: *configFile}
    config.Init()

    client := component.Client{
        Host: config.Logstash.Host,
        Login: config.Logstash.Login,
        Password: config.Logstash.Password,
    }

    list := client.FindLogs(request)

    decorator := component.Decorator{}
    decorator.DecorateRequest(request)

    for _, log := range list {
        decorator.DecorateDetails(log)
    }
}
