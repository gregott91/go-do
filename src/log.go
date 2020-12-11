package godo

import (
	"log"
	"os"
)

const (
	logName = "debug.log"
)

// InitializeLogging sets up debug file logging
func InitializeLogging() error {
	f, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(f)

	return nil
}
