package utils

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1("../test.txt")
	fmt.Println("GOT = ", got)
	want := 2

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

// func TestPart2(t *testing.T) {
// 	got := Part2("../test.txt")
// 	want := 12

// 	if got != want {
// 		t.Errorf("got %d, wanted %d", got, want)
// 	}
// }
