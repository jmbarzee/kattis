package graph

import (
	"strconv"
)

func BuildGraph(nodes int, edges [][]string) *Graph {
	g := NewGraph(nodes)

	for _, edge := range edges {
		n, err := strconv.Atoi(edge[0])
		if err != nil {
			panic(err)
		}
		m, err := strconv.Atoi(edge[1])
		if err != nil {
			panic(err)
		}
		c, err := strconv.ParseFloat(edge[2], 32)
		if err != nil {
			panic(err)
		}
		addEdge(g.Nodes[n], g.Nodes[m], float64(c))
	}

	return g
}

const (
	maxNodes = 200
	maxEdges = 1000
)

// func GenerateGraph() *Graph {
// 	g := NewGraph()

// 	for i := 0; i < maxNodes; i++ {
// 		name := strconv.Itoa(i)
// 		cost := 1 //rand.Intn(10)
// 		node := NewNode(name, cost)
// 		g.addNode(node)
// 	}

// 	for i := 0; i < maxEdges; i++ {
// 		src := strconv.Itoa(rand.Intn(maxNodes))
// 		dst := strconv.Itoa(rand.Intn(maxNodes))
// 		if src == dst {
// 			i--
// 			continue
// 		}
// 		g.addConn(src, dst)
// 	}
// 	return g
// }
