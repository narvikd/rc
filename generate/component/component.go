package component

import (
	"fmt"
	"github.com/TwiN/go-color"
	"io/ioutil"
	"path/filepath"
	"rc/pkg/osutils"
	"strings"
)

// Create handles component creation. It returns an error if it can't create its files / directories.
func Create(name string) error {
	m := NewModel(name)

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

func (m *Model) writeFile(fileType string, content []byte) error {
	componentDir := filepath.Join("src", "components", m.name)
	fileName := fmt.Sprintf("%s.%s", m.name, fileType)
	filePath := filepath.Join(componentDir, fileName)

	errCreateDir := osutils.CreateDir(componentDir)
	if errCreateDir != nil {
		return errCreateDir
	}

	err := ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		return fmt.Errorf("couldn't create file: %s. Error: %w", fileName, err)
	}
	fmt.Printf(color.InGreen("%s was successfully created at %s\n"), fileName, filePath)
	return nil
}

// createTSX uses getTSXPlaceHolder() as the template to generate the new .tsx
func (m *Model) createTSX() []byte {
	return []byte(strings.ReplaceAll(getTSXPlaceHolder(), "REPLACE", m.name))
}

func (m *Model) createCSS() []byte {
	return []byte(fmt.Sprintf(".%s {}", m.name))
}
