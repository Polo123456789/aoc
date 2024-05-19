package main

import (
// "fmt"
)

func maxLabelFrom(edges map[string]*Node) int {
	labels := make(map[int]int)
	for _, edge := range edges {
		count, ok := labels[edge.Label]
		if !ok {
			labels[edge.Label] = 1
		} else {
			labels[edge.Label] = count + 1
		}
	}
	//fmt.Println("\tLabels: ", labels)

	maxLabel := -1
	maxCount := -1
	for label, count := range labels {
		if count > maxCount {
			maxCount = count
			maxLabel = label
		}
	}
	//fmt.Println("\tMax label: ", maxLabel)
	return maxLabel
}

func (g *Graph) PropagateLabels() {
	noLabelsChanged := false
	for !noLabelsChanged {
		noLabelsChanged = true
		for _, val := range g.nodes {
			//fmt.Println("Checking: ", key)
			maxLabel := maxLabelFrom(val.Edges)
			if maxLabel != val.Label {
				noLabelsChanged = false
				val.Label = maxLabel
			}
		}
	}
}
