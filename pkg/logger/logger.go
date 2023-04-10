package logger

import "github.com/sirupsen/logrus"

func GetLogger() *logrus.Logger {
	log := logrus.New()
	log.SetReportCaller(true)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
