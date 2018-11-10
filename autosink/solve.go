package autoSink

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jmbarzee/kattis/autosink/graph"
)

func Solve(r io.Reader, w io.Writer) {
	bufR := bufio.NewReader(r)
	nodes := FetchTuples(bufR)
	conns := FetchTuples(bufR)
	g := graph.BuildGraph(nodes, conns)
	paths := FetchTuples(bufR)
	for _, path := range paths {
		cost, err := g.Dijkstra(path[0], path[1])
		if err != nil {
			fmt.Fprintln(w, "NO")
		} else {
			fmt.Fprintln(w, cost)
		}
	}
}

// FetchTuples gathers input from stdin
func FetchTuples(r *bufio.Reader) [][]string {
	n := ReadConstraint(r)
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

// ReadConstraint fetches N from the first line
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
