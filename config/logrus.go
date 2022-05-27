package configs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{
		// DisableTimestamp: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyMsg:   "message",
		},
	})
	if os.Getenv("ENVIRONMENT") == "live" {
		logrus.SetLevel(logrus.WarnLevel)
	} else {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.DebugLevel)
	}
}
