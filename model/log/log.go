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
	Object    string
	Source    map[string]interface{}
	Exception *Exception
}

type Exception struct {
	Message string
	Code    int
	Trace   []Trace
}

type Trace struct {
	File string
	Line int
}
