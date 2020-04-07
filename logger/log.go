package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetOutput(os.Stdout)
}
func DBLog(fields logrus.Fields, level logrus.Level, info string) {
	entry := logrus.WithFields(fields)
	switch level {
	case logrus.WarnLevel:
		entry.Warn(info)
	case logrus.ErrorLevel:
		entry.Error(info)
	case logrus.InfoLevel:
		entry.Info(info)
	}

}
