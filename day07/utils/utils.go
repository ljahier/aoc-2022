package utils

import (
	"os"
	"strings"
)

func Part1(filename string) int {
	var input []string = strings.Split(getFileData(filename), "\n")
	// var result int = getSum(data)

	coreCompute(input)

	return 0
}

func getFileData(filePath string) string {
	data, _ := os.ReadFile(filePath)

	return strings.Trim(string(data), "\n")
}

// func getSum(data string) int {
// 	parsingFile(data)
// 	return 0
// }
