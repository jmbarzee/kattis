package main

import (
	"os"
	"testing"
)

// TestFindHalfGalaxyCount verifies that UniqueTrees will pass Kattis
func TestFindHalfGalaxyCount(t *testing.T) {
	cases := testFindHalfGalaxyCountProvider()

	for i, c := range cases {
		f, err := os.Open(c.fileName)
		if err != nil {
			t.Error(err)
		}
		stars, distance := fetchStarsAndDistance(f)
		count, _ := FindHalfGalaxyCount(stars, distance)
		if count != c.result {
			t.Errorf("TestCase:%v Failed \n\texpected:%v \n\tresult:%v", i, c.result, count)
		}
	}
}

type testCase struct {
	result   int
	fileName string
}

func testFindHalfGalaxyCountProvider() []testCase {
	cases := []testCase{
		testCase{
			result:   0,
			fileName: "tests/0.test",
		},
		testCase{
			result:   4,
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
