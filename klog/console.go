package klog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

type consoleLogger struct {
	l level
}

func NewConsoleLogger(l level) *consoleLogger {
	return &consoleLogger{l: l}
}

func (cl *consoleLogger) log(format string, lv level, a ...interface{}) {
	if lv < cl.l {
		return
	}
	now := time.Now().Format("[2006-01-02 15:04:05.000]")
	pc, file, line, ok := runtime.Caller(2)
	msg := fmt.Sprintf(format, a...)
	var logs string
	if ok {
		logs = fmt.Sprintf("%s\t%s\t%s\t[%s:%s:%d]", lv.String(), now, msg, runtime.FuncForPC(pc).Name(), path.Base(file), line)
	} else {
		logs = fmt.Sprintf("%s\t%s", now, msg)
	}
	_, _ = fmt.Fprintln(os.Stdout, logs)
}
