package uni

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

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
