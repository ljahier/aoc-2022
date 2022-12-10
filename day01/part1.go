package main

import (
	"fmt"
	"strings"

	"github.com/ljahier/aoc/utils"
)

func Part1() {
	var data []string = strings.Split(utils.GetFileData("input.txt"), "\n\n")

	numberOfCalories := utils.GetTotalCalories(data)

	fmt.Println(numberOfCalories)
}
