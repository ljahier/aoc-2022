package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(filename string) int {
	var data []string = strings.Split(getFileData(filename), "\n")
	var result int = getSumOfPairs(data)

	return result
}

func getFileData(filePath string) string {
	data, _ := os.ReadFile(filePath)

	return strings.Trim(string(data), "\n")
}

func getSumOfPairs(pairs []string) int {
	var result int = 0

	for _, pairStr := range pairs {
		pair := convertStrLineToIntLine(pairStr) // => [24, 68]
		result += isPair(pair)
	}

	return result
}

func convertStrLineToIntLine(pairStr string) []string {
	var pair []string = []string{}
	pairSplited := strings.Split(pairStr, ",")

	for _, pairWithDash := range pairSplited {
		pair = append(pair, pairWithDash)
	}

	return pair
}

func isPair(pair []string) int {
	nbr1 := strings.Split(pair[0], "-")
	nbr2 := strings.Split(pair[1], "-")

	fmt.Println("pair = ", pair)
	fmt.Println("unwrap(strconv.Atoi(nbr1[0])).(int) = ", unwrap(strconv.Atoi(nbr1[0])).(int))
	fmt.Println("unwrap(strconv.Atoi(nbr1[1])).(int) = ", unwrap(strconv.Atoi(nbr1[1])).(int))
	fmt.Println("unwrap(strconv.Atoi(nbr2[0])).(int) = ", unwrap(strconv.Atoi(nbr2[0])).(int))
	fmt.Println("unwrap(strconv.Atoi(nbr2[1])).(int) = ", unwrap(strconv.Atoi(nbr2[1])).(int))

	if unwrap(strconv.Atoi(nbr1[0])).(int) <= unwrap(strconv.Atoi(nbr2[0])).(int) && unwrap(strconv.Atoi(nbr1[1])).(int) >= unwrap(strconv.Atoi(nbr2[1])).(int) {
		return 1
	} else if unwrap(strconv.Atoi(nbr1[0])).(int) >= unwrap(strconv.Atoi(nbr2[0])).(int) && unwrap(strconv.Atoi(nbr1[1])).(int) <= unwrap(strconv.Atoi(nbr2[1])).(int) {
		return 1
	}

	return 0
}

func unwrap(val interface{}, err error) interface{} {
	return val
}
