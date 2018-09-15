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
	stars, distance := fetchStarsAndDistance(os.Stdin)
	count, err := FindHalfGalaxyCount(stars, distance)
	if err != nil {
		fmt.Println("NO")
	} else {
		fmt.Println(count)
	}
}

func FindHalfGalaxyCount(stars *StarList, distance int64) (int, error) {
	half := stars.Len / 2
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
			return starsFound, nil
		}
		if stars.Len <= half {
			return 0, fmt.Errorf("Not Enough Stars Remaining to find half.")
		}
	}
	return 0, fmt.Errorf("Reached end of Universe. Something went wrong.")
}

// fetchStarsAndDistance gathers input from stdin
func fetchStarsAndDistance(r io.Reader) (*StarList, int64) {
	bufR := bufio.NewReader(r)
	d, n := ReadConstraints(bufR)
	stars := NewStarList()
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
		stars.Append(&Star{
			X: int64(x),
			Y: int64(y),
		})

	}
	return stars, int64(d)
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
