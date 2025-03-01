package main

import (
	"fmt"
)

type GraphvizProperty struct {
	Key   string
	Value string
}

func NewGraphvizProperty(key string, value string) *GraphvizProperty {
	return &GraphvizProperty{
		Key:   key,
		Value: value}
}

// build : fontname="Helvetica"
func (d *GraphvizProperty) String() string {
	return fmt.Sprintf("%s=\"%s\"", d.Key, d.Value)
}
