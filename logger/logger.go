package logger

import (
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
)

var (
	Logger log.Logger
)

func SetLogger(verbosity string) error {
	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}}
	formatter.Line = true
	formatter.File = true
	formatter.Package = true
	Logger.SetFormatter(&formatter)
	Logger.SetOutput(os.Stdout)
	level, err := log.ParseLevel(verbosity)

	if err != nil {
		return err
	}

	Logger.SetLevel(level)

	return nil
}