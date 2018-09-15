package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Val  int32
	West *Node
	East *Node
}

func NewNode(i int32) *Node {
	return &Node{
		Val: i,
	}
}

func (n *Node) Insert(i int32) *Node {
	if n == nil {
		return NewNode(i)
	}

	if i < n.Val {
		n.West = n.West.Insert(i)
	} else {
		n.East = n.East.Insert(i)
	}
	return n
}

func Equals(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil && b != nil {
		return false
	} else if a != nil && b == nil {
		return false
	}
	return Equals(a.West, b.West) && Equals(a.East, b.East)
}

// Main is your number one homie
func Main() {
	fmt.Println(UniqueTrees(fetchCeilings(os.Stdin)))
}

func UniqueTrees(ceilings [][]int32) int {

	trees := make([]*Node, len(ceilings))
	for i, ceiling := range ceilings {
		for _, num := range ceiling {
			trees[i] = trees[i].Insert(num)
		}
	}

	uniqueTrees := make([]*Node, 1)
	uniqueTrees[0] = trees[0]
	trees = trees[1:len(trees)]
Outer:
	for _, tree := range trees {
		for _, uniqueTree := range uniqueTrees {
			if Equals(tree, uniqueTree) {
				continue Outer
			}
		}
		uniqueTrees = append(uniqueTrees, tree)
	}

	return len(uniqueTrees)
}

// fetchLists gathers input from stdin
func fetchCeilings(r io.Reader) [][]int32 {
	bufR := bufio.NewReader(r)
	n, k := ReadConstraints(bufR)
	ceilings := make([][]int32, n)
	for i := 0; i < n; i++ {
		line, err := bufR.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		lineString := strings.TrimSuffix(string(line), "\n")

		ceiling := make([]int32, k)
		allNums := strings.Split(lineString, " ")
		for j, s := range allNums {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			ceiling[j] = int32(num)
		}
		ceilings[i] = ceiling
	}
	return ceilings
}

// ReadConstraints fetches N from the first line
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
