package main

import (
	"testing"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func TestInnerMain(t *testing.T) {
	scanner := helper.ScanFile("../example.txt")
	innerMain(scanner)
}
