package graph

import (
	"fmt"
)

type Graph struct {
	Map map[string]*Node
	//DijCache map[string]map[string]int
}

func NewGraph() *Graph {
	return &Graph{
		Map: make(map[string]*Node),
		//Cache: make(map[string]map[string]int),
	}
}

func (g *Graph) addNode(n *Node) {
	g.Map[n.Name] = n
}

func (g *Graph) addConn(srcName, dstName string) {
	src := g.Map[srcName]
	if src == nil {
		err := fmt.Errorf("Warning! Tried to add connection from a node not in list! %v -> %v", srcName, dstName)
		panic(err)
	}
	dst := g.Map[dstName]
	if dst == nil {
		err := fmt.Errorf("Warning! Tried to add connection to a node not in list! %v -> %v", srcName, dstName)
		panic(err)
	}
	src.Conns = append(src.Conns, dst)
}

type Node struct {
	Name  string
	Conns []*Node
	Cost  int

	// For Dijkstras
	DijkCost  int
	DijkIndex int
	DijkPrev  *Node
	DijkVisited bool

	// For Generation
	GenVisited bool
}

func NewNode(name string, cost int) *Node {
	return &Node{
		Name:  name,
		Conns: make([]*Node, 0),
		Cost:  cost,

		DijkCost:  0,
		DijkIndex: -1,
		DijkVisited: false,
	}
}

func (n *Node) SetDijkCost(prevCost int) {
	n.DijkCost = prevCost + n.Cost
}
