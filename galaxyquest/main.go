package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	points, distance := fetchStarsAndDistance(os.Stdin)
	count := FindHalfGalaxyCount(points, distance)
	if count == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println(count)
	}
}

func FindHalfGalaxyCount(points []Point, distance int64) int {
	half := len(points) / 2
	PointsByX := splitByX(points, distance, half)
	// TODO @jmbarzee speed up by stoping after visiting more than half
	for _, plx := range PointsByX {
		pointsByY := splitByY(plx, distance, half)
		for _, ply := range pointsByY {

			n := FindHalfGalaxyCountSub(ply, distance, half)
			if n != 0 {
				return n
			} else {
				return 0
			}
		}
	}
	return 0
}

func FindHalfGalaxyCountSub(points []Point, distance int64, half int) int {
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
			if SameGalaxy(star, crntStar, distance) {
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

// fetchStarsAndDistance gathers input from stdin
func fetchStarsAndDistance(r io.Reader) ([]Point, int64) {
	bufR := bufio.NewReader(r)
	d, n := ReadConstraints(bufR)
	points := make([]Point, n)
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
		points[i] = Point{
			X: int64(x),
			Y: int64(y),
		}

	}
	return points, int64(d)
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
