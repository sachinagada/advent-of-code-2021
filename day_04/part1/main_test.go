package main

import (
	"testing"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func TestCalculatePowerConsumption(t *testing.T) {
	scanner := helper.ScanFile("../example.txt")
	innerMain(scanner)
}
