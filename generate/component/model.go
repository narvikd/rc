package component

import (
	"path/filepath"
	"strings"
)

type Model struct {
	name         string
	componentDir string
}

func NewModel(name string) *Model {
	newName := strings.Title(name)
	return &Model{
		name:         newName,
		componentDir: filepath.Join("src", "components", newName),
	}
}

// getTSXPlaceHolder returns a simple functional component template.
func getTSXPlaceHolder() string {
	return `import React, { FC } from 'react';
import './REPLACE.css';

interface REPLACEProps {}

const REPLACE: FC<REPLACEProps> = () => (
  <div className="REPLACE">
    REPLACE Component
  </div>
);

export default REPLACE;
`
}
