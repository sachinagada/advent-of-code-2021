package main

import (
	"testing"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func TestCalculatePowerConsumption(t *testing.T) {
	scanner := helper.ScanFile("../example.txt")

	if pc := calculatePowerConsumption(scanner); pc != 198 {
		t.Errorf("expected 198; got %d", pc)
	}
}
