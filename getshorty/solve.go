package getshorty

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jmbarzee/kattis/getshorty/graph"
)

func Solve(r io.Reader, w io.Writer) {
	bufR := bufio.NewReader(r)
	for {
		n, m := ReadConstraints(bufR)
		if n == 0 && m == 0 {
			break
		}
		conns := FetchTuples(bufR, m)
		g := graph.BuildGraph(n, conns)
		cost, err := g.Dijkstra(0, n-1)
		if err != nil {
			fmt.Fprintln(w, "Failed")
		} else {
			fmt.Fprintf(w, "%.4f\n", cost)
		}
	}
}

// FetchTuples gathers input from stdin
func FetchTuples(r *bufio.Reader, n int) [][]string {
	tuples := make([][]string, n)
	for i := 0; i < n; i++ {
		line, err := r.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		lineString := strings.TrimSuffix(string(line), "\n")
		tuples[i] = strings.Split(lineString, " ")
	}
	return tuples
}

// ReadConstraint fetches N, M from the first line
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
	m, err := strconv.Atoi(allNums[1])
	if err != nil {
		panic(err)
	}
	return n, m
}
