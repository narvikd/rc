package fileutils

import (
	"fmt"
	"io/ioutil"
)

// WriteToFile writes content to filePath. (It uses 600 permissions)
func WriteToFile(filePath string, content []byte) error {
	err := ioutil.WriteFile(filePath, content, 0600)
	if err != nil {
		return fmt.Errorf("couldn't create file: %s. Error: %w", filePath, err)
	}
	return nil
}
