package config

import "github.com/sirupsen/logrus"

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "Jan 01 2006 15:04:05",
	})
	log.SetLevel(logrus.Level(6))
	return log
}
