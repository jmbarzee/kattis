package graph

type Graph struct {
	Nodes []*Node
}

func NewGraph(nodes int) *Graph {
	g := &Graph{
		Nodes: make([]*Node, nodes),
	}
	for i := range g.Nodes {
		g.Nodes[i] = NewNode(i)
	}
	return g
}

func addEdge(n, m *Node, cost float64) {
	n.Edges = append(n.Edges, &Edge{
		Dst:  m,
		Src:  n,
		Cost: cost,
	})
	m.Edges = append(m.Edges, &Edge{
		Dst:  n,
		Src:  m,
		Cost: cost,
	})
}

type Edge struct {
	Dst  *Node
	Src  *Node
	Cost float64
}

type Node struct {
	Name  int
	Edges []*Edge

	// For Dijkstras
	DijkCost    float64
	DijkIndex   int
	DijkVisited bool
}

func NewNode(name int) *Node {
	return &Node{
		Name:  name,
		Edges: make([]*Edge, 0),

		DijkCost:    0,
		DijkIndex:   -1,
		DijkVisited: false,
	}
}
