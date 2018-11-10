package graph

import (
	"container/heap"
	"fmt"
)

func (g *Graph) Dijkstra(srcName, dstName int) (float64, error) {
	for _, n := range g.Nodes {
		n.DijkCost = 1
		n.DijkIndex = -1
		n.DijkVisited = false
	}

	src := g.Nodes[srcName]
	if src == nil {
		err := fmt.Errorf("Warning! Tried to path from a node not in list!")
		panic(err)
	}

	src.DijkCost = 1
	q := make(DijkQueue, 1)
	q[0] = src
	heap.Init(&q)

	for len(q) > 0 {
		crnt := heap.Pop(&q).(*Node)
		crnt.DijkVisited = true
		//fmt.Printf("visit: %v(%v)\n", crnt.Name, crnt.DijkCost)
		if crnt.Name == dstName {
			return crnt.DijkCost, nil
		}
		for _, edge := range crnt.Edges {
			//fmt.Printf("\t\t-> %v(%v)\n", edge.Dst.Name, edge.Cost)
			if edge.Dst.DijkVisited {
				// Already visited
				continue

			} else if edge.Dst.DijkIndex == -1 {
				// Found Node
				edge.Dst.DijkCost = crnt.DijkCost * edge.Cost
				heap.Push(&q, edge.Dst)

			} else if edge.Dst.DijkCost < crnt.DijkCost*edge.Cost {
				// Old Node, shorter path
				edge.Dst.DijkCost = crnt.DijkCost * edge.Cost
				heap.Fix(&q, edge.Dst.DijkIndex)
			}
		}
	}

	return -1, fmt.Errorf("Dijksta failed to find a path from the source to the destination")
}

// A DijkQueue implements heap.Interface and holds DijkNode.
type DijkQueue []*Node

func (dq DijkQueue) Len() int { return len(dq) }

func (dq DijkQueue) Less(i, j int) bool {
	// We want Pop to give us the highest priority so we use greater than here.
	if dq[i].DijkCost != dq[j].DijkCost {
		return dq[i].DijkCost > dq[j].DijkCost
	}
	return dq[i].Name < dq[j].Name
}

func (dq DijkQueue) Swap(i, j int) {
	dq[i], dq[j] = dq[j], dq[i]
	dq[i].DijkIndex = i
	dq[j].DijkIndex = j
}

func (dq *DijkQueue) Push(nodeI interface{}) {
	i := len(*dq)
	node := nodeI.(*Node)
	node.DijkIndex = i
	*dq = append(*dq, node)
}

func (dq *DijkQueue) Pop() interface{} {
	old := *dq
	n := len(old)
	item := old[n-1]
	item.DijkIndex = -1 // for safety
	*dq = old[0 : n-1]
	return item
}
