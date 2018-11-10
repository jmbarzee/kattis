package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	g := ReadUniverse(os.Stdin)
	count := FindMaxGalaxyCount(g)
	if count == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println(count)
	}
}

type (
	Universe struct {
		Distance int64
		Points   []Point
	}

	Point struct {
		X int64
		Y int64
	}
)

func FindMaxGalaxyCount(u Universe) int {
	// TODO @jmbarzee speed up by stoping after visiting more than half

	half := len(u.Points) / 2
	PointsByX := splitByX(u.Points, u.Distance, half)
	for _, plx := range PointsByX {

		pointsByY := splitByY(plx, u.Distance, half)
		for _, ply := range pointsByY {

			n := countPosibleGalaxy(ply, u.Distance, half)
			if n != 0 {
				return n
			} else {
				return 0
			}
		}
	}
	return 0
}

func countPosibleGalaxy(points []Point, distance int64, half int) int {
	stars := NewStarList()
	for _, p := range points {
		stars.Append(&Star{
			Point: p,
		})
	}

	for i := 0; i < stars.Len; i++ {
		star := stars.Head
		stars.Remove(star)
		nextStar := stars.Head
		starsFound := 1
		for j := 0; j < stars.Len; j++ {
			crntStar := nextStar
			nextStar = nextStar.Next
			if SameGalaxy(star.Point, crntStar.Point, distance) {
				stars.Remove(crntStar)
				starsFound++
				j--
			}
		}
		if starsFound > half {
			return starsFound
		}
		if stars.Len <= half {
			return 0
		}
	}
	return 0
}

// ReadUniverse gathers input from stdin
func ReadUniverse(r io.Reader) Universe {
	bufR := bufio.NewReader(r)
	distance, n := ReadConstraints(bufR)
	u := Universe{
		Points:   make([]Point, n),
		Distance: int64(distance),
	}
	for i := 0; i < n; i++ {
		line, err := bufR.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		lineString := strings.TrimSuffix(string(line), "\n")
		point := strings.Split(lineString, " ")
		x, err := strconv.Atoi(point[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(point[1])
		if err != nil {
			panic(err)
		}
		u.Points[i] = Point{
			X: int64(x),
			Y: int64(y),
		}

	}
	return u
}

// ReadConstraints fetches constraints from the first line
func ReadConstraints(r *bufio.Reader) (int, int) {
	line, err := r.ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	lineString := strings.TrimSuffix(string(line), "\n")
	allNums := strings.Split(lineString, " ")
	n, err := strconv.Atoi(allNums[0])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(allNums[1])
	if err != nil {
		panic(err)
	}
	return n, k
}

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
