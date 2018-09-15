package main

import (
	"os"
	"testing"
)

// TestUniqueTrees verifies that UniqueTrees will pass Kattis
func TestUniqueTrees(t *testing.T) {
	cases := testUniqueTreesProvider()

	for i, c := range cases {
		f, err := os.Open(c.fileName)
		if err != nil {
			t.Error(err)
		}
		ceilings := fetchCeilings(f)
		count := UniqueTrees(ceilings)
		if count != c.result {
			t.Errorf("TestCase:%v Failed \n\texpected:%v \n\tresult:%v", i, c.result, count)
		}
	}
}

type testCase struct {
	result   int
	fileName string
}

func testUniqueTreesProvider() []testCase {
	cases := []testCase{
		testCase{
			result:   4,
			fileName: "tests/0.test",
		},
		testCase{
			result:   2,
			fileName: "tests/1.test",
		},
		// testCase{
		// 	result:   1,
		// 	fileName: "tests/2.test",
		// },
		// testCase{
		// 	result:   0,
		// 	fileName: "tests/3.test",
		// },
		// testCase{
		// 	result:   3,
		// 	fileName: "tests/4.test",
		// },
		// testCase{
		// 	result:   1,
		// 	fileName: "tests/5.test",
		// },
	}
	return cases
}
