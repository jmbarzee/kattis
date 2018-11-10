package autoSink

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Solve(r io.Reader, w io.Writer) {
	bufR := bufio.NewReader(r)
	peepHeap, t := FetchHeap(bufR)

	schedule := make([]int, t)

	for len(*peepHeap) > 0 {
		peep := heap.Pop(peepHeap).(Peep)
		pos := peep.Limit
		for pos >= 0 {
			if schedule[pos] == 0 {
				schedule[pos] = peep.Value
				break
			} else {
				pos--
			}
		}
	}

	total := 0
	for _, deposite := range schedule {
		total += deposite
	}
	fmt.Fprintf(w, "%v\n", total)
}

type (
	Peep struct {
		Limit int
		Value int
	}
	PeepHeap []Peep
)

func (h PeepHeap) Len() int            { return len(h) }
func (h PeepHeap) Less(i, j int) bool  { return h[i].Value > h[j].Value }
func (h PeepHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PeepHeap) Push(x interface{}) { *h = append(*h, x.(Peep)) }

func (h *PeepHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// FetchHeap gathers input from stdin
func FetchHeap(r *bufio.Reader) (*PeepHeap, int) {
	n, t := ReadConstraints(r)
	peeps := &PeepHeap{}
	for i := 0; i < n; i++ {
		line, err := r.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		lineString := strings.TrimSuffix(string(line), "\n")
		nums := strings.Split(lineString, " ")
		value, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		limit, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		heap.Push(peeps, Peep{Limit: limit, Value: value})
	}
	return peeps, t
}

// ReadConstraints fetches n and t from the first line
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
	t, err := strconv.Atoi(allNums[1])
	if err != nil {
		panic(err)
	}
	return n, t
}
