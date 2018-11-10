package graph

import (
	"math/rand"
	"strconv"
)

func BuildGraph(nodes, conns [][]string) *Graph {
	g := NewGraph()
	for _, node := range nodes {
		name := node[0]
		cost, err := strconv.Atoi(node[1])
		if err != nil {
			panic(err)
		}
		g.addNode(NewNode(name, cost))
	}

	for _, c := range conns {
		g.addConn(c[0], c[1])
	}

	return g
}

const (
	maxNodes = 200
	maxEdges = 1000
)

func GenerateGraph() *Graph {
	g := NewGraph()

	for i := 0; i < maxNodes; i++ {
		name := strconv.Itoa(i)
		cost := 1 //rand.Intn(10)
		node := NewNode(name, cost)
		g.addNode(node)
	}

	for i := 0; i < maxEdges; i++ {
		src := strconv.Itoa(rand.Intn(maxNodes))
		dst := strconv.Itoa(rand.Intn(maxNodes))
		if src == dst {
			i--
			continue
		}
		g.addConn(src, dst)
	}
	return g
}
