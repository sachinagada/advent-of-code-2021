package main

import (
	"testing"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func TestParseInput(t *testing.T) {
	scanner := helper.ScanFile("../example.txt")
	output := parseInput(scanner)
	count := countOutput(output)
	if count != 26 {
		t.Errorf("expected 26, received %d", count)
	}
}
