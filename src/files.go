package godo

import (
	"fmt"
	"os"
	"path/filepath"
)

// ConcatenateFileWithCurrentExeDir concatenates the given file name with the current exe location
func ConcatenateFileWithCurrentExeDir(filename string) (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}

	exePath := filepath.Dir(exe)

	return fmt.Sprint(exePath, os.PathSeparator, filename), nil
}
