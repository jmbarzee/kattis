package rainbow

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// TestSolve verifies that Solve will pass Kattis
func TestSolve(t *testing.T) {
	cases := testSolveProvider(t)

	for i, c := range cases {
		t.Run(fmt.Sprintf("test:%v", i), func(t *testing.T) {
			r := bytes.NewReader(c.input)
			w := bytes.NewBuffer(make([]byte, 0))

			Solve(r, w)
			expected := string(c.output)
			b, err := ioutil.ReadAll(w)
			if err != nil {
				t.Error(err)
			}
			actual := string(b)

			if expected != actual {
				s := fmt.Sprintf("TestCase:%v Failed\n", i)
				s += "\tExp.\tAct.\n"
				expectedR := bytes.NewBuffer([]byte(expected))
				actualR := bytes.NewBuffer([]byte(actual))
				for {
					ex, err := expectedR.ReadBytes('\n')
					if err != nil {
						break
					}
					exS := strings.TrimSuffix(string(ex), "\n")
					ac, err := actualR.ReadBytes('\n')
					if err != nil {
						break
					}
					acS := strings.TrimSuffix(string(ac), "\n")
					failed := ""
					if exS != acS {
						failed = "*** Broken ***"
					}
					s += fmt.Sprintf("\t%s\t%s\t%s\n", exS, acS, failed)
				}
				t.Errorf(s)
			}
		})
	}
}

type testCase struct {
	input  []byte
	output []byte
}

func testSolveProvider(t *testing.T) []testCase {

	manualTests := make([]testCase, 4)
	for i := range manualTests {
		inName := fmt.Sprintf("tests/manual/%v.input", i)
		inFile, err := os.Open(inName)
		if err != nil {
			t.Error(err)
		}
		manualTests[i].input, err = ioutil.ReadAll(inFile)
		if err != nil {
			t.Error(err)
		}

		outName := fmt.Sprintf("tests/manual/%v.output", i)
		outFile, err := os.Open(outName)
		if err != nil {
			t.Error(err)
		}
		manualTests[i].output, err = ioutil.ReadAll(outFile)
		if err != nil {
			t.Error(err)
		}
	}

	generatedTests := make([]testCase, 0)
	for i := range generatedTests {
		inName := fmt.Sprintf("tests/generated/%v.input", i)
		inFile, err := os.Open(inName)
		if err != nil {
			t.Error(err)
		}
		generatedTests[i].input, err = ioutil.ReadAll(inFile)
		if err != nil {
			t.Error(err)
		}

		outName := fmt.Sprintf("tests/generated/%v.output", i)
		outFile, err := os.Open(outName)
		if err != nil {
			t.Error(err)
		}
		generatedTests[i].output, err = ioutil.ReadAll(outFile)
		if err != nil {
			t.Error(err)
		}
	}

	return append(manualTests, generatedTests...)
}
