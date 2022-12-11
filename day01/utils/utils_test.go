package utils

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1("../test.txt")
	want := 24000

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2("../test.txt")
	want := 45000

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
