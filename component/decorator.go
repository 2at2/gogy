package component

import (
    "fmt"
    "time"
    "gogy/model/log"
    "regexp"
    "strings"
    "github.com/fatih/color"
    "reflect"
    "gogy/model"
)

type Decorator struct{
}

func (o *Decorator) DecorateRequest(req model.Request) {
    fmt.Println()
    c := color.New(color.FgGreen, color.Bold)
    c.Println("Request")

    fmt.Printf(" • Query: %s", color.CyanString(req.Query))
    fmt.Println()
    fmt.Printf(" • Time start: %s", color.CyanString(fmt.Sprint(req.TimeStart)))
    fmt.Println()
    fmt.Printf(" • Time end: %s", color.CyanString(fmt.Sprint(req.TimeEnd)))
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

        var level string
        switch entity.Level.Code {
        case log.DEBUG:
            level = color.BlackString(entity.Level.Code)
            break
        case log.INFO, log.NOTICE:
            level = color.BlueString(entity.Level.Code)
            break
        case log.WARNING:
            level = color.YellowString(entity.Level.Code)
            break
        case log.ERROR:
            level = color.RedString(entity.Level.Code)
            break
        case log.CRITICAL, log.ALERT:
            s := color.New(color.FgWhite, color.BgRed).SprintFunc()
            level = fmt.Sprint(s(entity.Level.Code))
            break
        case log.EMERGENCY:
            s := color.New(color.FgWhite, color.Bold, color.BgHiRed).SprintFunc()
            level = fmt.Sprint(s(entity.Level.Code))
            break
        default:
            level = entity.Level.Code
        }
        fmt.Printf("%s ", level)

        message := entity.Message
        if placeholders {
            message = o.replacePlaceholders(message, entity.Source)
        }

        fmt.Printf("%s ", message)

        fmt.Println()
    }
}

func (obj *Decorator) replacePlaceholders(str string, placeholders map[string]interface{}) string {
    r := regexp.MustCompile(":\\w+")
    for _, key:= range r.FindAllString(str, -1) {
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