package log

import "errors"

const (
	EMERGENCY_CODE = 8
	ALERT_CODE = 7
	CRITICAL_CODE = 6
	ERROR_CODE = 5
	WARNING_CODE = 4
	NOTICE_CODE = 3
	INFO_CODE = 2
	DEBUG_CODE = 1

	EMERGENCY_SHORT_STRING = "EMER"
	ALERT_SHORT_STRING = "ALER"
	CRITICAL_SHORT_STRING = "CRIT"
	ERROR_SHORT_STRING = "ERRO"
	WARNING_SHORT_STRING = "WARN"
	NOTICE_SHORT_STRING = "NOTI"
	INFO_SHORT_STRING = "INFO"
	DEBUG_SHORT_STRING = "DEBU"

	EMERGENCY_LONG_STRING = "emergency"
	ALERT_LONG_STRING = "alert"
	CRITICAL_LONG_STRING = "critical"
	ERROR_LONG_STRING = "error"
	WARNING_LONG_STRING = "warning"
	NOTICE_LONG_STRING = "notice"
	INFO_LONG_STRING = "info"
	DEBUG_LONG_STRING = "debug"
)

type LogLevel struct {
	Code int
}

func (o *LogLevel) SetFromString(v string) error {
	switch v {
	case DEBUG_LONG_STRING:
	case DEBUG_SHORT_STRING:
		o.Code = DEBUG_CODE
		break
	case INFO_LONG_STRING:
	case INFO_SHORT_STRING:
		o.Code = INFO_CODE
		break
	case NOTICE_LONG_STRING:
	case NOTICE_SHORT_STRING:
		o.Code = NOTICE_CODE
		break
	case WARNING_LONG_STRING:
	case WARNING_SHORT_STRING:
		o.Code = WARNING_CODE
		break
	case ERROR_LONG_STRING:
	case ERROR_SHORT_STRING:
		o.Code = ERROR_CODE
		break
	case CRITICAL_LONG_STRING:
	case CRITICAL_SHORT_STRING:
		o.Code = CRITICAL_CODE
		break
	case ALERT_LONG_STRING:
	case ALERT_SHORT_STRING:
		o.Code = ALERT_CODE
		break
	case EMERGENCY_LONG_STRING:
	case EMERGENCY_SHORT_STRING:
		o.Code = EMERGENCY_CODE
		break
	default:
		return errors.New("Wrong log-level")
	}
	return nil
}
