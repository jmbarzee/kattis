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
	Solve(os.Stdin, os.Stdout)
}

var cache [][][]int

func Solve(r io.Reader, w io.Writer) {
	bufR := bufio.NewReader(r)

	rooms, k := FetchRows(bufR)
	fmt.Printf("rooms:\n%v\n", rooms)
	cache = make([][][]int, len(rooms))
	for r := range cache {
		cache[r] = make([][]int, k+1)
		for j := range cache[r] {
			cache[r][j] = make([]int, 3)
		}
	}
	maxVal := 0
	for r := len(cache) - 1; r >= 0; r-- {
		for kk := range cache[r] {
			for un := range cache[r][kk] {
				cache[r][kk][un] = maxValue(rooms, r, un-1, kk)
			}
		}
	}
	maxVal = cache[0][k][0]
	fmt.Fprintf(w, "%v\n", maxVal)
}

func maxValue(rooms [][]int, r, uncloseableRoom, k int) int {
	if r >= len(rooms)-1 {
		return 0
	}
	ret := 0

	for i := 0; i < r; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("maxValue(..., %v, %v, %v)\n", r, uncloseableRoom, k)

	if k == 0 {
		ret = rooms[r][0] + rooms[r][1] + cache[r+1][k][0]
	} else if k == len(rooms)-1-r {
		if uncloseableRoom == 1 {
			ret = rooms[r][1] + cache[r+1][k-1][2]
		} else if uncloseableRoom == 0 {
			ret = rooms[r][0] + cache[r+1][k-1][1]
		} else if uncloseableRoom == -1 {
			max1 := rooms[r][0] + cache[r+1][k-1][1]
			max2 := rooms[r][1] + cache[r+1][k-1][2]
			if max1 > max2 {
				ret = max1
			} else {
				ret = max2
			}
		} else {
			panic(fmt.Errorf("unclosable room is %v", uncloseableRoom))
		}

	} else if k < len(rooms)-1-r {
		if uncloseableRoom == 1 {
			max1 := rooms[r][1] + cache[r+1][k-1][2]
			max2 := rooms[r][0] + rooms[r][1] + cache[r+1][k][0]
			if max1 > max2 {
				ret = max1
			} else {
				ret = max2
			}
		} else if uncloseableRoom == 0 {
			max1 := rooms[r][0] + cache[r+1][k-1][1]
			max2 := rooms[r][0] + rooms[r][1] + cache[r+1][k][0]
			if max1 > max2 {
				ret = max1
			} else {
				ret = max2
			}
		} else if uncloseableRoom == -1 {
			max1 := rooms[r][0] + cache[r+1][k-1][1]
			max2 := rooms[r][0] + rooms[r][1] + cache[r+1][k][0]
			max3 := rooms[r][1] + cache[r+1][k-1][2]
			if max1 > max2 && max1 > max3 {
				ret = max1
			} else if max2 > max1 && max2 > max3 {
				ret = max2
			} else {
				ret = max3
			}
		} else {
			panic(fmt.Errorf("unclosable room is %v", uncloseableRoom))
		}

	} else {
		ret = 0
		// panic(fmt.Errorf("Need to close more rooms, maxValue(..., %v, %v, %v)", r, uncloseableRoom, k))
	}

	for i := 0; i < r; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("=> %v\n", ret)
	return ret
}

// FetchRows gathers input from stdin
func FetchRows(r *bufio.Reader) ([][]int, int) {
	const maxDistance = 30000
	n, k := ReadConstraints(r)
	rooms := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		rooms[i] = make([]int, 2)
		line, err := r.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		lineString := strings.TrimSuffix(string(line), "\n")
		nums := strings.Split(lineString, " ")
		rooms[i][0], err = strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		rooms[i][1], err = strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

	}
	return rooms, k
}

// ReadConstraints fetches n from the first line
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
