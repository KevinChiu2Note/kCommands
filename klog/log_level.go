package klog

const (
	DEBUG level = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type level uint8

func (lv level) String() string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "EROOR"
	case FATAL:
		return "FATAL"
	default:
		return ""
	}
}
