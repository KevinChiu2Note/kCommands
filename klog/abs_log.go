package klog

type ILogger interface {
	log(format string, lv level, a ...interface{})
}

type baseLogger struct {
	ILogger
}

func NewLogger(l ILogger) *baseLogger {
	return &baseLogger{l}
}

func (bl *baseLogger) Debug(format string, a ...interface{}) {
	bl.ILogger.log(format, DEBUG, a...)
}
func (bl *baseLogger) Trace(format string, a ...interface{}) {
	bl.ILogger.log(format, TRACE, a...)

}
func (bl *baseLogger) Info(format string, a ...interface{}) {
	bl.ILogger.log(format, INFO, a...)
}
func (bl *baseLogger) Warning(format string, a ...interface{}) {
	bl.ILogger.log(format, WARNING, a...)

}
func (bl *baseLogger) Error(format string, a ...interface{}) {
	bl.ILogger.log(format, ERROR, a...)
}
func (bl *baseLogger) Fatal(format string, a ...interface{}) {
	bl.ILogger.log(format, FATAL, a...)
}
func (bl *baseLogger) log(format string, lv level, a ...interface{}) {
}
