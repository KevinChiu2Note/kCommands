package klog

import "testing"

func TestLog(t *testing.T) {
	logger := NewLogger(NewConsoleLogger(DEBUG))
	logger.Debug("debug.................")
	logger.Trace("trace........")
	logger.Info("info......................")
	logger.Warning("warning..")
	logger.Fatal("fatal...............................")
	logger.Error("error.......")
}

func TestNewFileLogger(t *testing.T) {
	fl, err := NewFileLogger("./", "err.log", "info.log", 1024*1024*10, DEBUG)
	if err != nil {
		panic(err)
	}
	logger := NewLogger(fl)
	for i := 0; i < 15000; i++ {
		logger.Debug("debug.................")
		logger.Debug("debug.................")
		logger.Debug("debug.................")
		logger.Debug("debug.................")
		logger.Debug("debug.................")

		logger.Error("debug.................")
	}
}
