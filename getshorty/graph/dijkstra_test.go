package graph

// func TestGraph(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		t.Run(fmt.Sprintf("test:%v", i), func(t *testing.T) {
// 			g := GenerateGraph()
// 			trip := GenerateTrip(g)
// 			srcName := trip[0].name
// 			dstName := trip[len(trip)-1].name
// 			expectedCost := trip[len(trip)-1].total
// 			actualCost, err := g.Dijkstra(srcName, dstName)
// 			if err != nil {
// 				t.Errorf("Failed to find a generated path")
// 			}
// 			if actualCost != expectedCost {
// 				s := fmt.Sprintf("TestCase:%v Failed\n", i)
// 				s += fmt.Sprintf("\t%v -> %v\n", srcName, dstName)
// 				s += "Path:\n"
// 				s += "\tExp.\tAct.\n"
// 				crnt := g.Map[dstName]
// 				for i := len(trip) - 1; i > -1; i-- {
// 					exName := trip[i].name
// 					exCost := trip[i].cost

// 					acName := ""
// 					if crnt != nil {
// 						acName = crnt.Name
// 					}

// 					differs := ""
// 					if exName != acName {
// 						differs = "*** path diverges ***"
// 					}

// 					if crnt != nil {
// 						acCost := crnt.Cost

// 						s += fmt.Sprintf("\t%s(%v)\t%s(%v)\t%s\n", exName, exCost, acName, acCost, differs)
// 						crnt = crnt.DijkPrev
// 					} else {
// 						s += fmt.Sprintf("\t%s(%v)\t\t%s\n", exName, exCost, differs)
// 					}
// 				}
// 				t.Errorf(s)
// 			}
// 		})
// 	}
// }

// func GenerateTrip(g *Graph) []stop {
// 	for _, n := range g.Map {
// 		n.GenVisited = false
// 	}

// 	// Randomly select starting node
// 	srcName := strconv.Itoa(rand.Intn(maxNodes))

// 	crnt := g.Map[srcName]
// 	if crnt == nil {
// 		err := fmt.Errorf("Warning! Tried to path from a node not in list!")
// 		panic(err)
// 	}
// 	crnt.GenVisited = true

// 	// Randomly select trip length (check bounds)
// 	l := rand.Intn(maxEdges / 10)
// 	trip := make([]stop, 1)
// 	trip[0] = stop{
// 		name:  crnt.Name,
// 		cost:  0,
// 		total: 0,
// 	}

// 	// Randlomly select Next Node
// 	for i := 1; i < l; i++ {
// 		if len(crnt.Conns) == 0 {
// 			break
// 		}
// 		//advance and add to trip
// 		for _, conn := range crnt.Conns {
// 			if !conn.GenVisited {
// 				crnt = conn
// 				crnt.GenVisited = true
// 				trip = append(trip, stop{
// 					name:  crnt.Name,
// 					cost:  crnt.Cost,
// 					total: trip[i-1].total + crnt.Cost,
// 				})
// 			}
// 		}
// 	}
// 	return trip
// }

// type (
// 	stop struct {
// 		name  string
// 		cost  int
// 		total int
// 	}
// )
