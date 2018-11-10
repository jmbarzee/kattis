package uni_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/jmbarzee/kattis/galaxyquest"
	"github.com/jmbarzee/kattis/galaxyquest/tests"
)

// TestFindMaxGalaxyCount verifies that FindMaxGalaxyCount will pass Kattis
func TestFindMaxGalaxyCount(t *testing.T) {
	var cases []testhelper.TestCase
	cases = append(cases, manualTestProvider(t)...)
	cases = append(cases, automaticTestProvider(t)...)

	for i, c := range cases {
		t.Run(fmt.Sprintf("test:%v - %v ", i, c.FileName), func(t *testing.T) {
			count := uni.FindMaxGalaxyCount(c.Uni)
			if count != c.Expected {
				t.Errorf("testhelper.TestCase:%v Failed \n\texpected:%v \n\tresult:%v", i, c.Expected, count)
			}
		})
	}
}

func manualTestProvider(t *testing.T) []testhelper.TestCase {
	cases := []testhelper.TestCase{
		testhelper.TestCase{
			Expected: 0,
			FileName: "0.test",
		},
		testhelper.TestCase{
			Expected: 4,
			FileName: "1.test",
		},
		testhelper.TestCase{
			Expected: 0,
			FileName: "2.test",
		},
		testhelper.TestCase{
			Expected: 3,
			FileName: "3.test",
		},
		testhelper.TestCase{
			Expected: 3,
			FileName: "4.test",
		},
		testhelper.TestCase{
			Expected: 0,
			FileName: "5.test",
		},
		testhelper.TestCase{
			Expected: 2,
			FileName: "6.test",
		},
		testhelper.TestCase{
			Expected: 3,
			FileName: "7.test",
		},
		testhelper.TestCase{
			Expected: 0,
			FileName: "8.test",
		},
	}
	return testhelper.FetchUniverses(t, cases, "tests/manual")
}

func automaticTestProvider(t *testing.T) []testhelper.TestCase {
	var files []string
	root := "tests/generated"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".test" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	cases := make([]testhelper.TestCase, len(files))
	for i := range cases {
		cases[i].FileName = files[i]
	}
	cases = testhelper.FetchUniverses(t, cases, "")
	cases = testhelper.FetchSolutions(t, cases, "")
	return cases
}
