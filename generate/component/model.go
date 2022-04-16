package component

import "strings"

type Model struct {
	name string
}

func NewModel(name string) *Model {
	return &Model{name: strings.Title(name)}
}

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
