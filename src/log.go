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
	logFile, err := ConcatenateFileWithCurrentExeDir(logName)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(f)

	return nil
}
