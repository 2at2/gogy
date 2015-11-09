package log

import "errors"

const (
	EMERGENCY = "EMER"
	ALERT     = "ALER"
	CRITICAL  = "CRIT"
	ERROR     = "ERRO"
	WARNING   = "WARN"
	NOTICE    = "NOTI"
	INFO      = "INFO"
	DEBUG     = "DEBU"
)

type LogLevel struct {
	Code string
}

func (o *LogLevel) SetFromString(v string) error {
	switch v {
	case "debug":
		o.Code = DEBUG
		break
	case "info":
		o.Code = INFO
		break
	case "notice":
		o.Code = NOTICE
		break
	case "warning":
		o.Code = WARNING
		break
	case "error":
		o.Code = ERROR
		break
	case "critical":
		o.Code = CRITICAL
		break
	case "alert":
		o.Code = ALERT
		break
	case "emergency":
		o.Code = EMERGENCY
		break
	default:
		return errors.New("Wrong log-level")
	}
	return nil
}
