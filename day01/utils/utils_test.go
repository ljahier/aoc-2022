package utils

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	var data []string = strings.Split(GetFileData("../test.txt"), "\n\n")

	got := GetTotalCalories(data)
	want := 24000

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	var data []string = strings.Split(GetFileData("../test.txt"), "\n\n")

	got := GetTotalThreeMostCalories(data)
	want := 45000

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
