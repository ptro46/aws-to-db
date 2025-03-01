package main

import (
	"fmt"
)

type GraphvizNode struct {
	Name       string
	Properties []*GraphvizProperty
}

func NewGraphvizNode(name string) *GraphvizNode {
	return &GraphvizNode{
		Name:       name,
		Properties: make([]*GraphvizProperty, 0, 0)}
}

func (graphvizNode *GraphvizNode) addProperty(key string, value string) {
	graphvizNode.Properties = append(graphvizNode.Properties, NewGraphvizProperty(key, value))
}

// build : "ELBDealsBackend" [fontname="Helvetica" peripheries="1"  fontsize="10" label="ELBDealsBackend"];
func (graphvizNode *GraphvizNode) String() string {
	graphvizNodeDescription := fmt.Sprintf("\"%s\"", graphvizNode.Name)
	for index, graphvizProperty := range graphvizNode.Properties {
		if index == 0 {
			graphvizNodeDescription = graphvizNodeDescription + " ["
		}
		graphvizNodeDescription = graphvizNodeDescription + fmt.Sprintf("%+v ", graphvizProperty)
	}
	if len(graphvizNode.Properties) > 0 {
		graphvizNodeDescription = graphvizNodeDescription + " ];"
	}
	return graphvizNodeDescription
}
