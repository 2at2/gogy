package main

import (
	"flag"
	"fmt"
	"gogy/component"
	"gogy/model"
	"os"
	"time"
)

func main() {
	configFile := flag.String("config", "", "Config file")
	lastHours := flag.Int("lastHours", 24, "Logs of the last hourse")
	size := flag.Int("size", 10, "Size")
	flag.Parse()

	id := flag.Arg(0)

	if len(id) == 0 {
		fmt.Println("Argument `id` missed")
		os.Exit(0)
	}

	query := fmt.Sprintf(`_id:"%s"`, id)

	endTime := time.Now()
	startTime := endTime.Add(-time.Duration(*lastHours) * time.Hour)

	request := model.Request{
		Query:     query,
		TimeStart: startTime.UnixNano() / int64(time.Millisecond),
		TimeEnd:   endTime.UnixNano() / int64(time.Millisecond),
		Size:      *size,
	}

	config := component.Config{}
	config.InitConfigFile(*configFile)

	client := component.Client{
		Host:     config.Source["logstash.host"].(string),
		Login:    config.Source["logstash.login"].(string),
		Password: config.Source["logstash.password"].(string),
	}

	list := client.FindLogs(request)

	decorator := component.Decorator{}
	decorator.DecorateRequest(request)

	for _, log := range list {
		decorator.DecorateDetails(log)
	}
}
