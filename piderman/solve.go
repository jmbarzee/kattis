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

var cache [][]cacheItem

const absoluteMaxHeight = 500

type cacheItem struct {
	minHeight int
}

func Solve(r io.Reader, w io.Writer) {
	bufR := bufio.NewReader(r)

	n := ReadConstraint(bufR)
	for i := 0; i < n; i++ {
		// fetch moves
		moves := FetchRow(bufR)

		// init cache
		cache = make([][]cacheItem, len(moves))
		for j := range cache {
			cache[j] = make([]cacheItem, absoluteMaxHeight)
		}

		// fill cache
		for pos := len(cache) - 1; pos >= 0; pos-- {
			for height := range cache[pos] {
				movesLeft := moves[pos:len(moves)]

				minHeight := minHeight(movesLeft, height, height)
				cache[pos][height].minHeight = minHeight
			}
		}

		if cache[0][0].minHeight == -1 {
			fmt.Fprintf(w, "IMPOSSIBLE\n")
		} else {
			height := 0
			for pos, move := range moves {
				if pos == len(moves)-1 {
					fmt.Fprintf(w, "D\n")
				} else {
					up := -1
					down := -1
					if height+move < absoluteMaxHeight {
						up = cache[pos+1][height+move].minHeight
					}
					if height-move >= 0 {
						down = cache[pos+1][height-move].minHeight
					}
					if up == -1 {
						fmt.Fprintf(w, "D")
						height = height - move
					} else if down == -1 {
						fmt.Fprintf(w, "U")
						height = height + move
					} else if up < down {
						fmt.Fprintf(w, "U")
						height = height + move
					} else {
						fmt.Fprintf(w, "D")
						height = height - move
					}
				}
			}

		}
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
	} else if len(moves) == 1 {
		if height-moves[0] == 0 {
			return maxHeight
		} else {
			return -1
		}
	} else {
		move := moves[0]
		// movesLeft := []int{}
		// if len(moves) > 1 {
		// 	movesLeft = moves[1:len(moves)]
		// }
		// up := minHeight(movesLeft, height+move, max(height+move, maxHeight))
		// down := minHeight(movesLeft, height-move, maxHeight)
		pos := len(cache) - len(moves)
		up := -1
		down := -1
		if height == 3 && pos == 2 {
			fmt.Printf("")
		}
		if height+move < absoluteMaxHeight {
			up = cache[pos+1][height+move].minHeight
		}
		if height-move >= 0 {
			ctemp := cache[pos+1][height-move].minHeight
			if ctemp != -1 {
				down = max(ctemp, height)
			}
		}
		return minNegativeAvoid(up, down)
	}

}

func printCache() {
	for j := absoluteMaxHeight - 1; j >= 0; j-- {
		for i := range cache {
			fmt.Printf("%.3v ", cache[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
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
