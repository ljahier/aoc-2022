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

func Part2(filename string) int {
	var data []string = strings.Split(strings.Trim(GetFileData(filename), "\n"), "\n")
	var result int = GetTotalScoreCheat(data)

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

func GetTotalScoreCheat(data []string) int {
	var result int = 0

	for _, v := range getRoundScoreCheat(data) {
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

func getRoundScoreCheat(rounds []string) []int {
	var roundScores []int

	for _, round := range rounds {
		roundScores = append(roundScores, computeRoundScoreCheat(round))
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

func computeRoundScoreCheat(round string) int {
	userMove := strings.Split(round, " ")[1]
	elveMove := strings.Split(round, " ")[0]
	winShapes := map[string]string{"A": "Y", "B": "Z", "C": "X"}
	loseShapes := map[string]string{"A": "Z", "B": "X", "C": "Y"}
	drawShapes := map[string]string{"A": "X", "B": "Y", "C": "Z"}
	var userCheatMove string

	switch userMove {
	case "Y":
		userCheatMove = drawShapes[elveMove]
	case "X":
		userCheatMove = loseShapes[elveMove]
	case "Z":
		userCheatMove = winShapes[elveMove]
	}
	arr := []string{
		elveMove,
		userCheatMove,
	}
	cheatRound := strings.Join(arr, "")

	return computeRoundScore(cheatRound)
}
