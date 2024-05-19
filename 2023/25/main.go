package main

import (
	"fmt"
	"os"
)

// Print for visualization in: https://csacademy.com/app/graph_editor/
func printGraph(g Graph) {
	for key, val := range g.nodes {
		for edge := range val.Edges {
			fmt.Println(key, edge)
		}
	}
}

func main() {
	graph := NewGraphFrom(os.Stdin)
	graph.PropagateLabels()
	groups := graph.GetGroupSizes()
	assert(len(groups) == 2, "Only two groups should be present")
	val := 1
	for _, group := range groups {
		val *= group
	}
	fmt.Println(val)
}
