package main

import (
	"bufio"
	"io"
	"strings"
)

type NodeCreationInstruction struct {
	BaseID       string
	ConnectToIDS []string
}

func NewNodeCreationInstruction(line string) NodeCreationInstruction {
	op := NodeCreationInstruction{}

	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)

	assert(scanner.Scan(), "Could not scan for baseID")
	baseID := scanner.Text()
	assert(baseID[len(baseID)-1] == ':', "Invalid baseID: "+baseID)
	op.BaseID = baseID[:len(baseID)-1]

	for scanner.Scan() {
		op.ConnectToIDS = append(op.ConnectToIDS, scanner.Text())
	}

	return op
}

func NewGraphFrom(r io.Reader) Graph {
	scanner := bufio.NewScanner(r)
	graph := NewGraph()

	for scanner.Scan() {
		line := scanner.Text()
		err := graph.Execute(NewNodeCreationInstruction(line))
		assert(err == nil, err)
	}

	return graph
}
