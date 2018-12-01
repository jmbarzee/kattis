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

	n := ReadConstraint(bufR)
	for i := 0; i < n; i++ {
		moves := FetchRow(bufR)
		minH := minHeight(moves, 0, 0)
		fmt.Fprintf(w, "%v\n", minH)
	}
}

func minHeight(moves []int, height, maxHeight int) int {
	if height < 0 {
		return -1
	}

	if len(moves) == 0 {
		if height == 0 {
			return maxHeight
		} else {
			return -1
		}
	} else {
		move := moves[0]
		movesLeft := []int{}
		if len(moves) > 1 {
			movesLeft = moves[1:len(moves)]
		}
		up := minHeight(movesLeft, height+move, max(height+move, maxHeight))
		down := minHeight(movesLeft, height-move, maxHeight)
		return minNegativeAvoid(up, down)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func minNegativeAvoid(a, b int) int {
	if a == -1 {
		return b
	} else if b == -1 {
		return a
	} else {
		return min(a, b)
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// FetchRow gathers input from stdin
func FetchRow(r *bufio.Reader) []int {
	n := ReadConstraint(r)
	moves := make([]int, n)

	line, err := r.ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	lineString := strings.TrimSuffix(string(line), "\n")
	nums := strings.Split(lineString, " ")
	for i, num := range nums {
		moves[i], err = strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
	}
	return moves
}

// ReadConstraint fetches n from the first line
func ReadConstraint(r *bufio.Reader) int {
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
	return n
}
