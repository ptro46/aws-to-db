package main

import (
	"fmt"
)

type GraphvizArrow struct {
	From       *GraphvizNode
	To         *GraphvizNode
	Properties []*GraphvizProperty
}

func NewGraphvizArrow(from *GraphvizNode, to *GraphvizNode) *GraphvizArrow {
	return &GraphvizArrow{
		From:       from,
		To:         to,
		Properties: make([]*GraphvizProperty, 0, 0)}
}

func (graphvizArrow *GraphvizArrow) addProperty(key string, value string) {
	graphvizArrow.Properties = append(graphvizArrow.Properties, NewGraphvizProperty(key, value))
}

// build : "ELBBackOfficeBackend" -> "ELBDealsBackend" [label="10.1.2.0/24"]
func (graphvizArrow *GraphvizArrow) String() string {
	graphvizArrowDescription := fmt.Sprintf("\"%s\" -> \"%s\"", graphvizArrow.From.Name, graphvizArrow.To.Name)
	for index, graphvizProperty := range graphvizArrow.Properties {
		if index == 0 {
			graphvizArrowDescription = graphvizArrowDescription + " ["
		}
		graphvizArrowDescription = graphvizArrowDescription + fmt.Sprintf("%+v ", graphvizProperty)
	}
	if len(graphvizArrow.Properties) > 0 {
		graphvizArrowDescription = graphvizArrowDescription + " ]"
	}
	return graphvizArrowDescription
}
