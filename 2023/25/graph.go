package main

import (
	"fmt"
)

type Graph struct {
	nodes      map[string]*Node
	labelCount int
}

type Node struct {
	ID    string
	Edges map[string]*Node
	Label int
}

func NewGraph() Graph {
	return Graph{
		nodes:      make(map[string]*Node),
		labelCount: 0,
	}
}

func (g *Graph) NewNode(id string) *Node {
	if g.nodes == nil {
		g.nodes = make(map[string]*Node)
	}

	g.labelCount++
	node := &Node{
		ID:    id,
		Edges: make(map[string]*Node),
		Label: g.labelCount,
	}
	g.nodes[id] = node

	return node
}

func (g *Graph) Exists(id string) bool {
	if g.nodes == nil {
		return false
	}
	_, ok := g.nodes[id]
	return ok
}

func (g *Graph) ConnectNodes(nodeID, edgeID string) error {
	if g.nodes == nil {
		return fmt.Errorf("Cant connect nodes if there are none")
	}

	node, ok := g.nodes[nodeID]
	if !ok {
		return fmt.Errorf("Could not find node of id: %s", nodeID)
	}
	edge, ok := g.nodes[edgeID]
	if !ok {
		return fmt.Errorf("Could not find node of id: %s", edgeID)
	}

	node.Edges[edgeID] = edge
	edge.Edges[nodeID] = node
	return nil
}

func (g *Graph) Execute(instruction NodeCreationInstruction) error {
	if !g.Exists(instruction.BaseID) {
		g.NewNode(instruction.BaseID)
	}

	for _, c := range instruction.ConnectToIDS {
		if !g.Exists(c) {
			g.NewNode(c)
		}
		err := g.ConnectNodes(instruction.BaseID, c)
		if err != nil {
			return err
		}
	}

	return nil
}

// Key is the group label, value is the size of the group
type GraphGroups map[int]int

func (g *Graph) GetGroupSizes() GraphGroups {
	groupSizes := make(map[int]int)
	for _, node := range g.nodes {
		group, ok := groupSizes[node.Label]
		if !ok {
			groupSizes[node.Label] = 1
		} else {
			groupSizes[node.Label] = group + 1
		}
	}
	return groupSizes
}
