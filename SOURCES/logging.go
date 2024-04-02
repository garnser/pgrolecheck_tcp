// logging.go
package main

import (
	"log"
	"log/syslog"
	"os"
)

// SetupLogging initializes the logging system based on configuration.
func SetupLogging(logDest string, foreground bool) *log.Logger {
	if foreground {
		return log.New(os.Stdout, "", log.LstdFlags)
	}

	switch logDest {
	case "syslog":
		sysLogger, err := syslog.NewLogger(syslog.LOG_NOTICE|syslog.LOG_DAEMON, log.LstdFlags)
		if err != nil {
			log.Fatalf("Failed to initialize syslog: %v", err)
		}
		return sysLogger
	case "":
		return log.New(os.Stdout, "", log.LstdFlags)
	default:
		file, err := os.OpenFile(logDest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		return log.New(file, "", log.LstdFlags)
	}
}
