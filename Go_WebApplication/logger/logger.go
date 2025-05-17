package logger

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

var Log = zerolog.New(zerolog.ConsoleWriter{})

// SetupLogger :
func SetupLogger(logger zerolog.Logger) error {
	var level, filePath string

	level = "debug"
	fmt.Println(level)

	filePath = "/var/log/webapp/web-app.log"
	if filePath == "" {
		return errors.New("logger file path not found")
	}

	basePath := filepath.Dir(filePath)
	created, errStr := CheckPathExists(basePath)
	if !created {
		return errStr
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	Log = zerolog.New(zerolog.ConsoleWriter{Out: file, NoColor: false, TimeFormat: time.RFC3339}).With().Timestamp().Logger()

	return nil
}

// Check if directory exists
func CheckPathExists(dirPath string) (bool, error) {
	// Check if directory exists
	_, statErr := os.Stat(dirPath)
	if os.IsNotExist(statErr) {
		// Attempt to create directory
		mkdirErr := os.MkdirAll(dirPath, 0755)
		if mkdirErr != nil {
			return false, mkdirErr
		}
		return true, nil
	}
	return true, statErr
}
