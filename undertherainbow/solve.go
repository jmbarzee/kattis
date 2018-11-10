package rainbow

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Solve(r io.Reader, w io.Writer) {
	bufR := bufio.NewReader(r)

	locations := FetchLocations(bufR)

	// set initial node for "base case" (if you were to imagine this as a recursive solution)
	locations[len(locations)-1].MinPenalty = 0

	// find min penalty for all nodes starting at the node right before the destination
	for i := len(locations) - 2; i > -1; i-- {

		// find penalty if we were to next travel to the location at j
		for j := i + 1; j < len(locations); j++ {
			travelDistance := locations[j].Distance - locations[i].Distance
			partialPenalty := (400 - travelDistance) * (400 - travelDistance) // long way around for (400 - travelDistance)^2
			totalPenalty := partialPenalty + locations[j].MinPenalty

			// check if option is shorter
			if totalPenalty < locations[i].MinPenalty {
				locations[i].MinPenalty = totalPenalty
			}

			// possible optimization by realizing minPenalties are getting worse.
		}
	}

	fmt.Fprintf(w, "%v\n", locations[0].MinPenalty)
}

type (
	Locations struct {
		Distance   int
		MinPenalty int
	}
)

// FetchLocations gathers input from stdin
func FetchLocations(r *bufio.Reader) []Locations {
	const maxDistance = 30000
	n := ReadConstraints(r)
	locations := make([]Locations, n+1)
	for i := 0; i < n+1; i++ {
		line, err := r.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		lineString := strings.TrimSuffix(string(line), "\n")
		nums := strings.Split(lineString, " ")
		distance, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		locations[i].Distance = distance
		locations[i].MinPenalty = maxDistance * maxDistance // this should overshoot any min penalty so all solutions

	}
	return locations
}

// ReadConstraints fetches n from the first line
func ReadConstraints(r *bufio.Reader) int {
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
