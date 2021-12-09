package main

import (
	"testing"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func TestIsSubstring(t *testing.T) {
	if !isSubstring("fab", "ab") {
		t.Errorf("ab is a valid substring of fab")
	}

	if isSubstring("tab", "cd") {
		t.Errorf("cd is not a substring of tab")
	}
}

func TestIsEqual(t *testing.T) {
	if !isEqual("cdfeb", "fbecd") {
		t.Errorf("expected equal")
	}

	if isEqual("cat", "hat") {
		t.Errorf("expected inequality")
	}
}

func TestGetDigit(t *testing.T) {
	output := []int{1, 4, 7, 8}
	dig := getDigit(output)
	if dig != 1478 {
		t.Errorf("expected 1478; received %d", dig)
	}
}

func TestSearchSegment(t *testing.T) {
	scanner := helper.ScanFile("../example.txt")
	sum := searchSegment(scanner)

	if sum != 61229 {
		t.Errorf("expected 61229; received %d", sum)
	}
}
