package pastabank

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/strebul/gogy/component"
	"github.com/strebul/gogy/component/worker"
	"github.com/strebul/gogy/model"
	"github.com/strebul/gogy/model/log"
)

func Report(cl component.Client, duration int) {
	object := `Psp.Module.Notification.Components.Adapter.PastabankAdapter`

	var request model.Request
	var logs []log.Log

	c := color.New(color.FgGreen, color.Bold)
	c.Println("Pastabank")

	// Received mails
	request = worker.Request(object, "Created mail with id :id", "", duration)
	logs = cl.FindLogs(request)

	fmt.Printf(" • %-20s %s", "Received:", color.CyanString(fmt.Sprint(cap(logs))))
	fmt.Println()

	// Processed mails
	request = worker.Request(object, "Found mail with id :id", "", duration)
	logs = cl.FindLogs(request)

	fmt.Printf(" • %-20s %s", "Processed:", color.CyanString(fmt.Sprint(cap(logs))))
	fmt.Println()

	// Files
	request = worker.Request(object, "Starts parsing attach :file", "", duration)
	logs = cl.FindLogs(request)

	fmt.Printf(" • %-20s %s", "Files:", color.CyanString(fmt.Sprint(cap(logs))))
	fmt.Println()

	// Processed chb
	request = worker.Request(object, "Processed :c1 and stored :c2 chargebacks", "", duration)
	logs = cl.FindLogs(request)

	var chargeback int
	for _, log := range logs {
		chargeback += int(log.Source["c1"].(float64))
	}

	fmt.Printf(" • %-20s %s", "Chargebacks:", color.RedString(fmt.Sprint(chargeback)))

	// Similar chb
	request = worker.Request(object, "Found similar chargeback with id :id", "", duration)
	logs = cl.FindLogs(request)

	fmt.Printf(" / %s", color.BlackString(fmt.Sprint(cap(logs))))
	fmt.Println()

	// Processed alerts
	request = worker.Request(object, "Processed :c1 and stored :c2 alerts", "", duration)
	logs = cl.FindLogs(request)

	var alerts int
	for _, log := range logs {
		alerts += int(log.Source["c1"].(float64))
	}

	fmt.Printf(" • %-20s %s", "Alerts:", color.RedString(fmt.Sprint(alerts)))

	// Similar alerts
	request = worker.Request(object, "Found similar alert with id :id", "", duration)
	logs = cl.FindLogs(request)

	fmt.Printf(" / %s", color.BlackString(fmt.Sprint(cap(logs))))
	fmt.Println()

	// Total
	request = worker.Request(object, "Send :count rows", ` NOT count: "0"`, duration)
	logs = cl.FindLogs(request)

	var sendCount int
	for _, log := range logs {
		sendCount += int(log.Source["count"].(float64))
	}

	fmt.Printf(" • %-20s %s", "Send:", color.CyanString(fmt.Sprint(sendCount)))
	fmt.Println()

	// Instances
	request = worker.Request(object, "Start pasta bank adapter", "", 1)
	logs = cl.FindLogs(request)

	instances := map[string]string{}

	for _, log := range logs {
		if _, ok := instances[log.ScriptId]; !ok {
			instances[log.ScriptId] = ""
		}
	}

	fmt.Printf(" • %-20s %s", "Instances (last h):", color.CyanString(fmt.Sprint(len(instances))))
	fmt.Println()
}
