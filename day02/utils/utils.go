package utils

import (
	"os"
	"strings"
)

const (
	Rock     int = 1
	Paper    int = 2
	Scissors int = 3
)

const (
	X int = Rock
	Y int = Paper
	Z int = Scissors
)

const (
	A int = Rock
	B int = Paper
	C int = Scissors
)

const (
	Win  int = 6
	Draw int = 3
	Lost int = 0
)

func Part1(filename string) int {
	var data []string = strings.Split(GetFileData(filename), "\n")
	var result int = GetTotalScore(data)

	return result
}

func GetFileData(filePath string) string {
	data, _ := os.ReadFile(filePath)

	return string(data)
}

// must to return the total addition score
func GetTotalScore(data []string) int {
	var result int = 0

	for _, v := range getRoundScore(data) {
		result += v
	}

	return result
}

// must to return int array with score for each round
func getRoundScore(rounds []string) []int {
	var roundScores []int

	for _, round := range rounds {
		roundScores = append(roundScores, computeRoundScore(round))
	}
	return roundScores
}

func computeRoundScore(round string) int {
	var result int = 0

	switch strings.Join(strings.Split(round, " "), "") {
	case "AX":
		result = Draw + X
	case "AY":
		result = Win + Y
	case "AZ":
		result = Lost + Z
	case "BX":
		result = Lost + X
	case "BY":
		result = Draw + Y
	case "BZ":
		result = Win + Z
	case "CX":
		result = Win + X
	case "CY":
		result = Lost + Y
	case "CZ":
		result = Draw + Z
	}

	return result
}
