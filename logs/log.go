package logs

import (
	"github.com/sirupsen/logrus"
)

type FieldLogger interface {
	WithFields(map[string]interface{}) Logger
}

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

type FieldLog struct {
	Logger logrus.FieldLogger
}

func (f *FieldLog) WithFields(fields map[string]interface{}) Logger {
	return f.Logger.WithFields(fields)
}
