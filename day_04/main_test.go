package main

import (
	"testing"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func TestInnerMain(t *testing.T) {
	scanner := helper.ScanFile("./example.txt")
	score := playBingo(scanner, lowestBoard)
	if score != 4512 {
		t.Errorf("expected 4512; got %d", score)
	}
}
