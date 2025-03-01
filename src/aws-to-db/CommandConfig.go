package main

import (
	"fmt"
)

type CommandConfig struct {
	Name string `json:"name,omitempty"`
	Cmd  string `json:"cmd,omitempty"`
}

func (d CommandConfig) String() string {
	return fmt.Sprintf("CommandConfig Name(%s) cmd(%s)", d.Name, d.Cmd)
}
