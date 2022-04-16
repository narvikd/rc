package component

import (
	"fmt"
	"github.com/TwiN/go-color"
	"path/filepath"
	"rc/pkg/fileutils"
	"rc/pkg/osutils"
	"strings"
)

// Create handles component creation. It returns an error if it can't create its files / directories.
func Create(name string) error {
	m := NewModel(name)

	errCreateDir := osutils.CreateDir(m.componentDir)
	if errCreateDir != nil {
		return errCreateDir
	}

	errTSX := m.writeFile("tsx", m.createTSX())
	if errTSX != nil {
		return errTSX
	}

	errCSS := m.writeFile("css", m.createCSS())
	if errCSS != nil {
		return errCSS
	}

	return nil
}

// writeFile is a wrapper for "fileutils.WriteToFile" to create a file using the Model specs.
func (m *Model) writeFile(fileType string, content []byte) error {
	fileName := fmt.Sprintf("%s.%s", m.name, fileType)
	filePath := filepath.Join(m.componentDir, fileName)

	err := fileutils.WriteToFile(filePath, content)
	if err != nil {
		return err
	}

	fmt.Printf(color.InGreen("%s was successfully created at %s\n"), fileName, filePath)
	return nil
}

// createTSX uses getTSXPlaceHolder() as the template to generate the new .tsx
func (m *Model) createTSX() []byte {
	return []byte(strings.ReplaceAll(getTSXPlaceHolder(), "REPLACE", m.name))
}

// createCSS creates a simple CSS file.
func (m *Model) createCSS() []byte {
	return []byte(fmt.Sprintf(".%s {}", m.name))
}
