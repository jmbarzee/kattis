package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	c := Anagrams(os.Stdin)
	fmt.Println(c)
}

// Anagrams exits exclusivly for testing purposes. Otherwise, everything would be in main()
func Anagrams(r io.Reader) int {
	m := make(map[string]int)

	bufR := bufio.NewReader(r)
	n := ReadConstraintN(bufR)
	for i := 0; i < n; i++ {
		line, err := bufR.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		sorted := SortString(string(line))
		m[sorted]++
	}

	count := 0
	for _, i := range m {
		if i == 1 {
			count++
		}
	}
	return count
}

// ReadConstraintN fetches N from the first line
func ReadConstraintN(r *bufio.Reader) int {
	line, err := r.ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(line), " ")[0]
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// SortString plays ping pong and drives a 1950's sports car.
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortChars(r))
	return string(r)
}

type sortChars []rune

func (s sortChars) Less(i, j int) bool { return s[i] < s[j] }
func (s sortChars) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortChars) Len() int           { return len(s) }
