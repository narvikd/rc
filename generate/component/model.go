package component

import (
	"path/filepath"
	"strings"
)

type Model struct {
	Name                 string
	ComponentDir         string
	IsReturnComponentSet bool
}

func NewModel(name string, isReturnComponentSet bool) *Model {
	newName := strings.Title(name)
	return &Model{
		Name:                 newName,
		ComponentDir:         filepath.Join("src", "components", newName),
		IsReturnComponentSet: isReturnComponentSet,
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

func getTSXPlaceHolderWithReturn() string {
	return `import React, { FC } from 'react';
import './REPLACE.css';

interface REPLACEProps {}

const REPLACE: FC<REPLACEProps> = () => {
  return (
      <div className="REPLACE">
          REPLACE Component
      </div>
  )
};

export default REPLACE;
`
}
