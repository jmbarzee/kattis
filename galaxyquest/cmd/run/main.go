package main

import (
	"fmt"
	"os"

	"github.com/jmbarzee/kattis/galaxyquest"
)

func main() {
	g := uni.ReadUniverse(os.Stdin)
	count := uni.FindMaxGalaxyCount(g)
	if count == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println(count)
	}
}
