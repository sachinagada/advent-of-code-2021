package main

import (
	"testing"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func TestInnerMain(t *testing.T) {
	scanner := helper.ScanFile("./example.txt")
	overlap := navigateVents(scanner)
	if overlap != 12 {
		t.Errorf("expected 12; received %d", overlap)
	}
}
