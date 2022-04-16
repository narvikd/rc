package osutils

import (
	"fmt"
	"os"
)

// CreateDir creates dirs even if the parent dir doesn't exist. (It uses 755 permissions)
func CreateDir(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("couldn't create directory: %s. Error: %w", path, err)
	}
	return nil
}
