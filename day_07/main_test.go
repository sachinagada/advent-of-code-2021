package main

import (
	"testing"
)

func TestInnerMain(t *testing.T) {
	horizontal := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	cost := calculateCost(horizontal, constantFuel)
	if cost != 37 {
		t.Errorf("expected cost of 37; received %d", cost)
	}

	cost = calculateCost(horizontal, progressiveFuel)
	if cost != 168 {
		t.Errorf("expected cost of 168; received %d", cost)
	}
}
