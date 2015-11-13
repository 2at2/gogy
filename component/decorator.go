package component

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/strebul/gogy/model"
	"github.com/strebul/gogy/model/log"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type Decorator struct {
}

func (o *Decorator) DecorateRequest(req model.Request) {
	fmt.Println()
	c := color.New(color.FgGreen, color.Bold)
	c.Println("Request")

	fmt.Printf(" • Query: %s", color.CyanString(req.Query))
	fmt.Println()
	fmt.Printf(" • Time start: %s", color.CyanString(fmt.Sprint(req.TimeStart.Format(time.Stamp))))
	fmt.Println()
	fmt.Printf(" • Time end: %s", color.CyanString(fmt.Sprint(req.TimeEnd.Format(time.Stamp))))
	fmt.Println()
	fmt.Printf(" • Size: %s", color.CyanString(fmt.Sprint(req.Size)))
	fmt.Println()
	fmt.Println()
}

func (o *Decorator) DecorateList(list []log.Log, placeholders bool) {
	for _, entity := range list {
		date := entity.Time.Format(time.Stamp)
		fmt.Printf("%s ", color.BlueString(date))

		fmt.Printf("%s ", color.GreenString(entity.Id))

		level := o.colorizeLevel(entity.Level)
		fmt.Printf("%s ", level)

		message := entity.Message
		if placeholders {
			message = o.replacePlaceholders(message, entity.Source)
		}

		fmt.Printf("%s ", message)

		fmt.Println()
	}
}

func (o *Decorator) DecorateDetails(entity log.Log) {
	fmt.Println()
	c := color.New(color.FgGreen, color.Bold)
	c.Println("Response")

	fmt.Printf(" • Id: %s", color.GreenString(entity.Id))
	fmt.Println()

	level := o.colorizeLevel(entity.Level)
	fmt.Printf(" • Level: %s", level)
	fmt.Println()

	date := entity.Time.Format(time.Stamp)
	fmt.Printf(" • Time: %s", color.WhiteString(date))
	fmt.Println()

	fmt.Printf(" • Host: %s", color.WhiteString(entity.Host))
	fmt.Println()

	if len(entity.SessionId) > 0 {
		fmt.Printf(" • Session id: %s", color.GreenString(entity.SessionId))
		fmt.Println()
	}

	message := o.replacePlaceholders(entity.Message, entity.Source)
	fmt.Printf(" • Message: %s", color.CyanString(message))
	fmt.Println()

	fmt.Printf(" • Script id: %s", color.YellowString(entity.ScriptId))
	fmt.Println()

	fmt.Printf(" • Object: %s", (entity.Object))
	fmt.Println()

	fmt.Println(" • Exception:")

	style1 := color.New(color.FgWhite, color.BgBlack)

	if entity.Exception != nil {
		style1.Printf("   • Message: %s", entity.Exception.Message)
		style1.Println()
		style1.Printf("   • Code: %d", entity.Exception.Code)
		style1.Println()

		for _, trace := range entity.Exception.Trace {
			style1.Printf("   • File: %s:%d", trace.File, trace.Line)
			style1.Println()
		}
	}

	fmt.Println()
}

func (obj *Decorator) replacePlaceholders(
	str string,
	placeholders map[string]interface{},
) string {
	r := regexp.MustCompile(":\\w+")
	for _, key := range r.FindAllString(str, -1) {
		name := strings.Replace(key, ":", "", -1)

		if value, ok := placeholders[name]; ok {
			switch reflect.TypeOf(value).String() {
			case "string":
				value = color.CyanString(fmt.Sprint(value))
				break
			case "int", "int64":
				value = color.BlueString(fmt.Sprintf("%d", value))
			case "float", "float64":
				value = color.BlueString(fmt.Sprintf("%f", value))
				break
			default:
				value = fmt.Sprint(value)
			}

			str = strings.Replace(str, key, value.(string), -1)
		}
	}

	return str
}

func (obj *Decorator) colorizeLevel(level log.LogLevel) string {
	var str string

	switch level.Code {
	case log.DEBUG:
		str = color.BlackString(level.Code)
		break
	case log.INFO:
		str = color.BlueString(level.Code)
		break
	case log.NOTICE:
		str = color.CyanString(level.Code)
		break
	case log.WARNING:
		str = color.YellowString(level.Code)
		break
	case log.ERROR:
		str = color.RedString(level.Code)
		break
	case log.CRITICAL, log.ALERT:
		s := color.New(color.FgWhite, color.BgRed).SprintFunc()
		str = fmt.Sprint(s(level.Code))
		break
	case log.EMERGENCY:
		s := color.New(color.FgWhite, color.Bold, color.BgHiRed).SprintFunc()
		str = fmt.Sprint(s(level.Code))
		break
	default:
		str = level.Code
	}

	return str
}
