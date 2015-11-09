package log

import (
	"time"
)

type Log struct {
	Id        string
	Level     LogLevel
	Message   string
	Time      time.Time
	Host      string
	ScriptId  string
	SessionId string
	Source    map[string]interface{}
}
