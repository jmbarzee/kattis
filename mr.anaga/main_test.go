package main

import (
	"os"
	"testing"
)

func TestAnagrams(t *testing.T) {
	cases := testAnagramsProvider()

	for i, c := range cases {
		f, err := os.Open(c.fileName)
		if err != nil {
			t.Error(err)
		}
		anagrams := Anagrams(f)
		if anagrams != c.result {
			t.Errorf("TestCase:%v Failed \n\texpected:%v \n\tresult:%v", i, c.result, anagrams)
		}
	}
}

type testCase struct {
	result   int
	fileName string
}

func testAnagramsProvider() []testCase {
	cases := []testCase{
		testCase{
			result:   1,
			fileName: "tests/0.test",
		},
		testCase{
			result:   2,
			fileName: "tests/1.test",
		},
		testCase{
			result:   1,
			fileName: "tests/2.test",
		},
		testCase{
			result:   0,
			fileName: "tests/3.test",
		},
		testCase{
			result:   3,
			fileName: "tests/4.test",
		},
		testCase{
			result:   1,
			fileName: "tests/5.test",
		},
	}
	return cases
}
