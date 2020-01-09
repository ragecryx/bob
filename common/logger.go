package common

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	commonLog *log.Entry // Used by this module internally

	// BuilderLog is used to log messages of the builder module
	BuilderLog *log.Entry
	// ServiceLog is used to log messages of the web-service module
	ServiceLog *log.Entry
)

// InitLogging sets the default logging
// behavior for logrus.
func InitLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	commonLog = log.WithField("system", "common")
	BuilderLog = log.WithField("system", "builder")
	ServiceLog = log.WithField("system", "web")
}
