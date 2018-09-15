package main

import "sort"

type (
	Point struct {
		X int64
		Y int64
	}
)

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
	if i+1-begin > half {
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

type (
	Star struct {
		Point
		Next *Star
		Prev *Star
	}

	StarList struct {
		Head *Star
		Tail *Star
		Len  int
	}
)

func NewStarList() *StarList {
	return &StarList{
		Len:  0,
		Head: nil,
		Tail: nil,
	}
}

func (sl *StarList) Remove(star *Star) {
	sl.Len--
	if sl.Head == star {
		// Star is at the Front
		if sl.Len > 0 {
			star.Next.Prev = nil
		}
		sl.Head = star.Next
	} else if sl.Tail == star {
		// Star is at the Back
		if sl.Len > 0 {
			star.Prev.Next = nil
		}
		sl.Tail = star.Prev
	} else {
		// Star is in the Middle
		star.Prev.Next = star.Next
		star.Next.Prev = star.Prev
	}
}

func (sl *StarList) Append(star *Star) {
	sl.Len++
	if sl.Len == 1 {
		sl.Head = star
		sl.Tail = star
	}
	sl.Tail.Next = star
	star.Prev = sl.Tail
	sl.Tail = star
}

func SameGalaxy(s1, s2 *Star, distance int64) bool {
	xDiff := s1.X - s2.X
	x2 := xDiff * xDiff
	yDiff := s1.Y - s2.Y
	y2 := yDiff * yDiff
	d2 := distance * distance
	return (x2 + y2) <= d2
}
