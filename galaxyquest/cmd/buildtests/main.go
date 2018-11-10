package main

import (
	"fmt"
	"os"
	"path"

	"github.com/jmbarzee/kattis/galaxyquest"
	"github.com/jmbarzee/kattis/galaxyquest/tests"
)

func main() {
	for i := 0; i < 5; i++ {
		u, sol := uni.GenerateUniverseTest()
		testCase := testhelper.TestCase{
			Uni:      u,
			Expected: sol,
			FileName: fmt.Sprintf("%v.test", i),
		}
		err := writeTestToFile(testCase)
		if err != nil {
			panic(err)
		}
	}
}
func writeTestToFile(test testhelper.TestCase) error {
	root := "tests/generated"
	f1, err := os.OpenFile(path.Join(root, test.FileName), os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f1.Close()

	fmt.Fprintf(f1, "%v %v\n", test.Uni.Distance, len(test.Uni.Points))
	for _, p := range test.Uni.Points {
		fmt.Fprintf(f1, "%v %v\n", p.X, p.Y)
	}
	f2, err := os.OpenFile(path.Join(root, test.FileName+".sol"), os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f1.Close()

	fmt.Fprintf(f2, "%v\n", test.Expected)
	return nil
}
