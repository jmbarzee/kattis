package testhelper

import (
	"bufio"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/jmbarzee/kattis/galaxyquest"
)

type TestCase struct {
	Uni      uni.Universe
	Expected int
	FileName string
}

func FetchUniverses(t *testing.T, cases []TestCase, dir string) []TestCase {
	for i, c := range cases {
		f, err := os.Open(path.Join(dir, c.FileName))
		if err != nil {
			t.Error(err)
		}
		cases[i].Uni = uni.ReadUniverse(f)
	}
	return cases
}

func FetchSolutions(t *testing.T, cases []TestCase, dir string) []TestCase {
	for i, c := range cases {
		f, err := os.Open(path.Join(dir, c.FileName+".sol"))
		if err != nil {
			t.Error(err)
		}
		bufR := bufio.NewReader(f)
		line, err := bufR.ReadBytes('\n')
		if err != nil {
			t.Error(err)
		}
		lineString := strings.TrimSuffix(string(line), "\n")
		allNums := strings.Split(lineString, " ")
		sol, err := strconv.Atoi(allNums[0])
		if err != nil {
			t.Error(err)
		}
		cases[i].Expected = sol
	}
	return cases
}
