package utils

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1("../test.txt")
	want := 15

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
