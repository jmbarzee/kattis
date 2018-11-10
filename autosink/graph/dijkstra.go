package graph

import (
	"container/heap"
	"fmt"
)

func (g *Graph) Dijkstra(srcName, dstName string) (int, error) {
	for _, n := range g.Map {
		n.DijkCost = 0
		n.DijkIndex = -1
		n.DijkVisited = false
		n.DijkPrev = nil
	}
	
	src := g.Map[srcName]
	if src == nil {
		err := fmt.Errorf("Warning! Tried to path from a node not in list!")
		panic(err)
	}

	src.DijkCost = 0
	q := make(DijkQueue, 1)
	q[0] = src
	heap.Init(&q)

	for len(q) > 0 {
		crnt := heap.Pop(&q).(*Node)
		crnt.DijkVisited = true
		//fmt.Printf("h:%v\t%s(%v) -> %v\n", len(q), crnt.Name, crnt.Cost, len(crnt.Conns))
		if crnt.Name == dstName {
			return crnt.DijkCost, nil
		}
		for _, conn := range crnt.Conns {
			if conn.DijkVisited {
				continue
			} else if conn.DijkIndex == -1  {
				conn.DijkPrev = crnt
				conn.SetDijkCost(crnt.DijkCost)
				heap.Push(&q, conn)
			} else if conn.DijkCost > crnt.DijkCost+conn.Cost {
				conn.DijkPrev = crnt
				conn.SetDijkCost(crnt.DijkCost)
				heap.Fix(&q, conn.DijkIndex)
			}
		}
	}

	return -1, fmt.Errorf("Dijksta failed to find a path from the source to the destination")
}

// A DijkQueue implements heap.Interface and holds DijkNode.
type DijkQueue []*Node

func (dq DijkQueue) Len() int { return len(dq) }

func (dq DijkQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use greater than here.
	if dq[i].Cost != dq[j].Cost {
		return dq[i].Cost < dq[j].Cost
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
