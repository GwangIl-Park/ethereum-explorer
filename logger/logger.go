package logger

import (
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
)

var (
	Logger logrus.Logger
)

func NewLogger(verbosity string) error {
	formatter := runtime.Formatter{
		ChildFormatter: &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			DisableSorting: true,
		},
		Line:true,
		File:true,
		Package:true,
	}
	Logger.SetFormatter(&formatter)
	Logger.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(verbosity)

	if err != nil {
		return err
	}

	Logger.SetLevel(level)

	return nil
}