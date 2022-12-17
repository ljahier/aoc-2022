package main

import (
	"fmt"
	"os"

	"github.com/ljahier/aoc-2022/utils"
)

func main() {
	var filename string = "./input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	fmt.Println("Part 1 : ", utils.Part1(filename))
	// fmt.Println("Part 2 : ", utils.Part2(filename))
}
