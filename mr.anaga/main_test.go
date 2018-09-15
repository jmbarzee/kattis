package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"testing"
)

// TestAnagrams verifies that anagrams will pass Kattis
func TestAnagrams(t *testing.T) {
	cases := testAnagramsProvider()

	for i, c := range cases {
		f, err := os.Open(c.fileName)
		if err != nil {
			t.Error(err)
		}
		anagrams := AnagramsReader(f)
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

type benchmarkTest struct {
	n int
	k int
}

func (b benchmarkTest) name() string {
	return fmt.Sprintf("( n:%v k:%v )", b.n, b.k)
}

func (b benchmarkTest) getWords() []string {
	words := make([]string, b.n)
	for i := range words {
		b := make([]byte, b.k)
		rand.Read(b)
		words[i] = string(b)
	}
	return words
}

func BenchmarkAnagrams(b *testing.B) {
	tests := benchmarkAnagramsProviderQ1()
	for _, test := range tests {
		words := test.getWords()
		b.Run(test.name(), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				Anagrams(words)
			}
		})
	}
}

func benchmarkAnagramsProvider() []benchmarkTest {
	tests := []benchmarkTest{}
	for i := 10; i < 1000001; i = i * 10 {
		for j := 10; j < 1000001; j = j * 10 {
			tests = append(tests, benchmarkTest{
				k: i,
				n: j,
			})
		}
	}
	return tests
}

func benchmarkAnagramsProviderQ1() []benchmarkTest {
	tests := []benchmarkTest{}
	for i := 10; i < 1000; i = i * 2 {
		tests = append(tests, benchmarkTest{
			n: 2000,
			k: i,
		})
	}
	return tests
}
