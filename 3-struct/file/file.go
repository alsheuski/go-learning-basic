// Package file implements base file operations
package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(fileName string) ([]byte, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		errorMessage := fmt.Sprintf("Error reading file %v: %v\n", fileName, err)
		return nil, errors.New(errorMessage)
	}

	return file, nil
}

func SaveFile(fileName string, data []byte) error {
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		errorMessage := fmt.Sprintf("Error writing file %v: %v\n", fileName, err)
		return errors.New(errorMessage)
	}

	return nil
}

func IsJSON(fileName string) bool {
	ext := filepath.Ext(fileName)
	return strings.EqualFold(ext, ".json")
}
