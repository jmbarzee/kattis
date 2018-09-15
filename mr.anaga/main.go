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

// Anagrams is returns the number or words which didn't share a character set with another word
func Anagrams(words []string) int {
	m := make(map[string]int)
	for _, w := range words {
		sorted := SortString(w)
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

// =============== Implementations for Kattis ================
func main() {
	c := AnagramsReader(os.Stdin)
	fmt.Println(c)
}

// AnagramsReader is a wrapper around Anagrams
func AnagramsReader(r io.Reader) int {
	bufR := bufio.NewReader(r)
	n := ReadConstraintN(bufR)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		line, err := bufR.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		words[i] = string(line)
	}
	return Anagrams(words)
}

// ReadConstraintN fetches N from the first line
func ReadConstraintN(r *bufio.Reader) int {
	line, err := r.ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(line), " ")

	i, err := strconv.Atoi(s[0])
	if err != nil {
		panic(err)
	}
	return i
}
