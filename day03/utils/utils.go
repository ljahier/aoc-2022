package utils

import (
	"fmt"
	"os"
	"strings"
)

func Part1(filename string) int {
	var data []string = strings.Split(GetFileData(filename), "\n")
	var result int = GetSumOfElementsPriority(data)

	return result
}

// func Part2(filename string) int {
// 	var data []string = strings.Split(GetFileData(filename), "\n")
// 	var result int

// 	return result
// }

func GetFileData(filePath string) string {
	data, _ := os.ReadFile(filePath)

	return strings.Trim(string(data), "\n")
}

func GetSumOfElementsPriority(rucksacks []string) int {
	var result int = 0

	for _, rucksack := range rucksacks {
		compartments := getCompartments(rucksack)
		fmt.Println("DEBUG commonItem = ", getCommonItem(compartments))
		// get item that appears in both compartments
		// - lowercase item have priorities 1 through 26
		// - uppercase item have priorities 27 through 52
		// compute the sum of priorities
	}

	return result
}

func getCompartments(rucksack string) []string {
	// split row in 2 rows and each rows must have the same length
	items := strings.Split(rucksack, "")
	itemsLen := len(items)
	var rucksackCompartment1 []string
	var rucksackCompartment2 []string

	for i := 0; i < itemsLen/2; i++ {
		rucksackCompartment1 = append(rucksackCompartment1, items[i])
	}

	for i := itemsLen / 2; i < itemsLen; i++ {
		rucksackCompartment2 = append(rucksackCompartment2, items[i])
	}

	if len(rucksackCompartment1) != len(rucksackCompartment2) {
		panic("Two compartments lengths do not match")
	}

	return []string{
		strings.Join(rucksackCompartment1, ""),
		strings.Join(rucksackCompartment2, ""),
	}
}

func getCommonItem(rucksackCompartments []string) string {
	var compartment1 []string = strings.Split(rucksackCompartments[0], "")
	var compartment2 []string = strings.Split(rucksackCompartments[1], "")
	var items map[string]int
	var result string

	for i := 0; i < len(compartment1); i++ {
		items[compartment1[i]] += 1
	}

	for i := 0; i < len(compartment2); i++ {
		if items[compartment2[i]] > 0 {
			result = compartment2[i]

			return result
		}
	}

	return result
}
