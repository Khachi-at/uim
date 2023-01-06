package logger

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type Setting struct {
	FileName    string
	Level       string
	Format      string
	RollingDays uint
}

func Init(settings Setting) error {
	if settings.Level == "" {
		settings.Level = "debug"
	}
	ll, err := logrus.ParseLevel(settings.Level)
	if err == nil {
		std.SetLevel(ll)
	} else {
		std.Error("Invalid log level")
	}

	if settings.FileName == "" {
		return nil
	}

	if settings.RollingDays == 0 {
		settings.RollingDays = 7
	}

	writer, err := rotatelogs.New(
		settings.FileName+".%Y%m%d",
		rotatelogs.WithLinkName(settings.FileName),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithRotationCount(settings.RollingDays),
	)
	if err != nil {
		return err
	}

	var logfr logrus.Formatter
	if settings.Format == "json" {
		logfr = &logrus.JSONFormatter{
			DisableTimestamp: false,
		}
	} else {
		logfr = &logrus.TextFormatter{
			DisableColors: true,
		}
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
	}, logfr)

	std.AddHook(lfsHook)
	return nil
}
