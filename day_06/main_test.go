package main

import (
	"testing"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func TestCountFish(t *testing.T) {
	scanner := helper.ScanFile("./example.txt")
	count := countFish(scanner, 80)
	if count != 5934 {
		t.Errorf("expected 5934; got %d", count)
	}

	scanner2 := helper.ScanFile("./example.txt")
	count = countFish(scanner2, 256)
	if count != 26984457539 {
		t.Errorf("expected 373378; got %d", count)
	}
}
