package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type LogxLevel logrus.Level

const (
	PanicLevel LogxLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

type Logx struct {
	prefix   string
	loglevel logrus.Level
	log      *logrus.Logger
}

func NewLogx(prefix string, loglevel LogxLevel) Logger {
	logx := &Logx{
		prefix:   prefix,
		loglevel: logrus.Level(loglevel),
		log:      logrus.StandardLogger(),
	}

	logx.log.SetLevel(logx.loglevel)
	logx.log.SetFormatter(&logrus.JSONFormatter{})

	return logx
}

func (l *Logx) Info(args ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", l.prefix, fmt.Sprint(args...))
	l.log.Info(msg)
}

func (l *Logx) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("[%s] %s", l.prefix, msg)
	l.log.Info(msg)
}

func (l *Logx) Trace(args ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", l.prefix, fmt.Sprint(args...))
	l.log.Trace(msg)
}

func (l *Logx) Tracef(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("[%s] %s", l.prefix, msg)
	l.log.Tracef(msg)
}

func (l *Logx) Debug(args ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", l.prefix, fmt.Sprint(args...))
	l.log.Debug(msg)
}

func (l *Logx) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("[%s] %s", l.prefix, msg)
	l.log.Debugf(msg)
}

func (l *Logx) Error(args ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", l.prefix, fmt.Sprint(args...))
	l.log.Error(msg)
}

func (l *Logx) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("[%s] %s", l.prefix, msg)
	l.log.Errorf(msg)
}

func (l *Logx) Fatal(args ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", l.prefix, fmt.Sprint(args...))
	l.log.Fatal(msg)
}

func (l *Logx) Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("[%s] %s", l.prefix, msg)
	l.log.Fatalf(msg)
}
