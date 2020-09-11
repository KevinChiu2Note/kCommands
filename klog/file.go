package klog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

type fileLogger struct {
	lv       level
	fp       string
	errFile  *os.File
	infoFile *os.File
	fs       int64
}

func NewFileLogger(fp, errFn, infoFn string, fs int64, lv level) (*fileLogger, error) {
	fl := &fileLogger{
		lv: lv,
		fp: fp,
		fs: fs,
	}
	p := path.Join(fp, infoFn)
	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	fl.infoFile = f

	p = path.Join(fp, errFn)
	f, err = os.OpenFile(p, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	fl.errFile = f
	return fl, nil
}

func (fl *fileLogger) log(format string, lv level, a ...interface{}) {
	if lv < fl.lv {
		return
	}
	now := time.Now().Format("[2006-01-02 15:04:05.000]")
	pc, file, line, ok := runtime.Caller(2)
	msg := fmt.Sprintf(format, a...)
	var logs string
	if ok {
		logs = fmt.Sprintf("%s\t%s\t%s\t[%s:%s:%d]", lv.String(), now, msg, runtime.FuncForPC(pc).Name(), path.Base(file), line)
	} else {
		logs = fmt.Sprintf("%s\t%s\n", now, msg)
	}
	if lv == ERROR {
		if fl.checkFileSize(fl.errFile) {
			fl.errFile = fl.backupAndGetNewFile(fl.errFile)
		}
		_, _ = fmt.Fprintf(fl.errFile, "%s\n",logs)
	} else {
		if fl.checkFileSize(fl.infoFile) {
			fl.infoFile = fl.backupAndGetNewFile(fl.infoFile)
		}
		_, _ = fmt.Fprintf(fl.infoFile,"%s\n", logs)
	}
}

func (fl *fileLogger) checkFileSize(f *os.File) bool {
	stat, err := f.Stat()
	if err != nil {
		return false
	}
	return stat.Size() >= fl.fs
}

func (fl *fileLogger) backupAndGetNewFile(f *os.File) *os.File {
	logName := f.Name()
	_ = f.Close()
	now := time.Now().Format("20060102_150405")
	newName := fmt.Sprintf("%s-%s", logName, now)
	_ = os.Rename(path.Join(fl.fp, logName), path.Join(fl.fp, newName))
	file, err := os.OpenFile(path.Join(fl.fp, logName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}
	return file
}
