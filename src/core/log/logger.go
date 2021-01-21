package log

import (
	"go.uber.org/zap"
)

var Logger *logger

type logger struct {
	zLogger *zap.Logger
	sugar   *zap.SugaredLogger
}

func Default(zLogger *zap.Logger) {
	Logger = &logger{
		zLogger: zLogger,
		sugar:   zLogger.Sugar(),
	}
}

func (l *logger) Info(format string, v ...interface{}) {
	l.sugar.Infof(format, v)
}

func (l *logger) Warn(format string, v ...interface{}) {
	l.sugar.Warnf(format, v)
}

func (l *logger) Debug(format string, v ...interface{}) {
	l.sugar.Debugf(format, v)
}

func (l *logger) Error(format string, v ...interface{}) {
	l.sugar.Errorf(format, v)
}

func (l *logger) Panic(format string, v ...interface{}) {
	l.sugar.Panicf(format, v)
}

func (l *logger) DPanic(format string, v ...interface{}) {
	l.sugar.DPanicf(format, v)
}

func (l *logger) Fatal(format string, v ...interface{}) {
	l.sugar.Fatalf(format, v)
}
