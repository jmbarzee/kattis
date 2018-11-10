package uni

import (
	"fmt"
	"sort"
)

func SameGalaxy(p1, p2 Point, distance int64) bool {
	xDiff := p1.X - p2.X
	x2 := xDiff * xDiff
	yDiff := p1.Y - p2.Y
	y2 := yDiff * yDiff
	d2 := distance * distance
	return (x2 + y2) <= d2
}

type sortByX []Point

func (s sortByX) Less(i, j int) bool { return s[i].X < s[j].X }
func (s sortByX) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByX) Len() int           { return len(s) }

func splitByX(points []Point, distance int64, half int) [][]Point {
	pointLists := make([][]Point, 0)
	sort.Sort(sortByX(points))
	begin := 0
	var i int
	for i = 0; i < len(points)-1; i++ {
		if inRangeX(points[i], points[i+1], distance) {
			continue
		}
		if i+1-begin <= half {
			begin = i + 1
			continue
		}
		pointLists = append(pointLists, points[begin:i+1])
		begin = i + 1
	}
	if len(points) > 0 && i+1-begin > half {
		fmt.Printf("len(points)=%v -> [%v:%v]\n", len(points), begin, i+1)
		pointLists = append(pointLists, points[begin:i+1])
	}
	return pointLists
}

func inRangeX(p1, p2 Point, distance int64) bool {
	xDiff := p1.X - p2.X
	x2 := xDiff * xDiff
	d2 := distance * distance
	return x2 <= d2
}

type sortByY []Point

func (s sortByY) Less(i, j int) bool { return s[i].X < s[j].X }
func (s sortByY) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByY) Len() int           { return len(s) }

func splitByY(points []Point, distance int64, half int) [][]Point {
	pointLists := make([][]Point, 0)
	sort.Sort(sortByY(points))
	begin := 0
	var i int
	for i = 0; i < len(points)-1; i++ {
		if inRangeY(points[i], points[i+1], distance) {
			continue
		}
		if i+1-begin <= half {
			begin = i + 1
			continue
		}
		pointLists = append(pointLists, points[begin:i+1])
		begin = i + 1
	}
	if i+1-begin > half {
		pointLists = append(pointLists, points[begin:i+1])
	}
	return pointLists
}

func inRangeY(p1, p2 Point, distance int64) bool {
	yDiff := p1.Y - p2.Y
	y2 := yDiff * yDiff
	d2 := distance * distance
	return y2 <= d2
}
