package main

import (
	"bytes"
	"fmt"
)

type ConfigAwsToDBConfig struct {
	DB       DbConfig        `json:"db,omitempty"`
	Commands []CommandConfig `json:"commands,omitempty"`
	Graph    []GraphConfig   `json:"graphConfig,omitempty"`
}

func (d *ConfigAwsToDBConfig) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("ConfigAwsToDBConfig\n")
	buffer.WriteString("Database\n")
	buffer.WriteString(fmt.Sprintf("    %+v\n", d.DB))
	buffer.WriteString("Commands\n")
	for _, command := range d.Commands {
		buffer.WriteString(fmt.Sprintf("    %+v\n", command))
	}
	return buffer.String()
}
