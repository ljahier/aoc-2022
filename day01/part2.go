package main

import (
	"fmt"
	"strings"

	"github.com/ljahier/aoc/utils"
)

func Part2() {
	var data []string = strings.Split(utils.GetFileData("input.txt"), "\n\n")

	numberOfCalories := utils.GetTotalThreeMostCalories(data)

	fmt.Println(numberOfCalories)
}
