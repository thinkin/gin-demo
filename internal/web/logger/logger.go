package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"gin-demo/internal/web/configs"
)

var L = logrus.New()

// InitLog init logger
// refer from https://github.com/sirupsen/logrus
func InitLog() {
	timestampFormat := "2006-01-02 15:04:05.000"
	switch configs.APPConfig.Log.Formatter {
	case "json":
		L.Formatter = &logrus.JSONFormatter{
			TimestampFormat: timestampFormat,
		}
	case "text":
		L.Formatter = &logrus.TextFormatter{
			TimestampFormat: timestampFormat,
		}
	default:
		L.Formatter = &logrus.TextFormatter{
			TimestampFormat: timestampFormat,
		}
	}

	logDir := configs.APPConfig.Log.Dir
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Sprintf("create log dir failed, %s", err.Error()))
	}
	logFile := fmt.Sprintf("%s/%s.log", logDir, configs.APPConfig.App)
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("init log open file failed, %s", err.Error()))
	}
	L.SetOutput(file)

	switch configs.APPConfig.Log.Level {
	case "trace":
		L.SetLevel(logrus.TraceLevel)
	case "debug":
		L.SetLevel(logrus.DebugLevel)
	case "info":
		L.SetLevel(logrus.InfoLevel)
	case "warn":
		L.SetLevel(logrus.WarnLevel)
	case "fatal":
		L.SetLevel(logrus.FatalLevel)
	case "panic":
		L.SetLevel(logrus.PanicLevel)
	default:
		L.SetLevel(logrus.DebugLevel)
	}
}
