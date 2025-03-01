package main

import (
	"fmt"
)

type GraphvizDiGraph struct {
	Name   string
	Nodes  map[string]*GraphvizNode
	Arrows []*GraphvizArrow
}

func NewGraphvizDiGraph(name string) *GraphvizDiGraph {
	return &GraphvizDiGraph{
		Name:   name,
		Nodes:  make(map[string]*GraphvizNode),
		Arrows: make([]*GraphvizArrow, 0, 0)}
}

func (graphvizDiGraph *GraphvizDiGraph) addNode(node *GraphvizNode) {
	graphvizDiGraph.Nodes[node.Name] = node
}
func (graphvizDiGraph *GraphvizDiGraph) addArrow(arrow *GraphvizArrow) {
	graphvizDiGraph.Arrows = append(graphvizDiGraph.Arrows, arrow)
}

func (graphvicDiGraph *GraphvizDiGraph) findNode(nodeName string) *GraphvizNode {
	for _, currentNode := range graphvicDiGraph.Nodes {
		if currentNode.Name == nodeName {
			return currentNode
		}
	}
	return nil
}

// build : all the digraph
func (graphvizDiGraph *GraphvizDiGraph) String(graphodes []GraphConfig) string {
	graphvizDigraphDescription := fmt.Sprintf("digraph \"%s\" {\n", graphvizDiGraph.Name)

	for _, graphvizNode := range graphvizDiGraph.Nodes {
		graphvizDigraphDescription = graphvizDigraphDescription + fmt.Sprintf("\t%+v\n", graphvizNode)
	}
	graphvizDigraphDescription = graphvizDigraphDescription + "\n"

	rankString := "" /*`
	{rank = same; "ELBZendeskFrontend"; "ELBBackOfficeFrontend"; "ELBDealsFrontend"}
	{rank = same; "ELBZendeskBackend"; "ELBBackOfficeBackend"; "ELBDealsBackend"; "ELBPartenairesBackend"}
	{rank = same; "ELBGatewayBackend"; "ELBCypheringBackend"; "ELBFilerBackend"}
	{rank = same; "bastion-admin-1"; "bastion-admin-vpn-2"}
	`*/
	ranks := make(map[int]string)
	for _, node := range graphodes {
		ranks[node.Rank] += fmt.Sprintf(`; "%s"`, node.Name)
	}

	for i := 1; i < len(ranks)+1; i++ {
		rankString += fmt.Sprintf("{rank = same%s}\n", ranks[i])
	}

	graphvizDigraphDescription = graphvizDigraphDescription + "\n" + rankString + "\n\n"

	for _, graphvizArrow := range graphvizDiGraph.Arrows {
		graphvizDigraphDescription = graphvizDigraphDescription + fmt.Sprintf("\t%+v\n", graphvizArrow)
	}
	graphvizDigraphDescription = graphvizDigraphDescription + "\n}\n"

	return graphvizDigraphDescription
}
