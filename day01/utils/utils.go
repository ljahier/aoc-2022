package utils

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetFileData(filePath string) string {
	data, _ := os.ReadFile(filePath)

	return string(data)
}

func GetTotalCalories(data []string) int {
	var maxValue int = 0

	for _, values := range data {
		splitedByNewLine := strings.Split(values, "\n")
		nbr := computeNbrArray(splitedByNewLine)
		if nbr > maxValue {
			maxValue = nbr
		}
	}

	return maxValue
}

func GetTotalThreeMostCalories(data []string) int {
	var result = 0
	var calories []int

	for _, values := range data {
		splitedByNewLine := strings.Split(values, "\n")
		calories = append(calories, computeNbrArray(splitedByNewLine))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	for _, value := range calories[:3] {
		result += value
	}
	return result
}

func computeNbrArray(arr []string) int {
	var result int = 0

	for _, value := range arr {
		nbr, _ := strconv.Atoi(value)
		result += nbr
	}

	return result
}
