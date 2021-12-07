package main

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func main() {
	scanner := helper.ScanFile("./input.txt")
	horizontal := parseInput(scanner)
	cost := calculateCost(horizontal, progressiveFuel)
	fmt.Println("cost", cost)

}

func parseInput(scanner *bufio.Scanner) []int {
	var horizontal []int
	for scanner.Scan() {
		l := scanner.Text()
		sArr := strings.Split(l, ",")

		horizontal = make([]int, len(sArr))
		for i, s := range sArr {
			v := helper.ParseInt(s)
			horizontal[i] = v
		}
	}
	return horizontal

}

func calculateCost(horizontal []int, f func(i, v int) int) int {
	sort.Ints(horizontal)

	// get the highest position that the crab is in
	cost := make([]int, horizontal[len(horizontal)-1])

	for i := range cost {
		for _, v := range horizontal {
			cost[i] += f(i, v)
		}
	}

	min := math.MaxInt32
	minPosition := 0
	for i, v := range cost {
		if v < min {
			min = v
			minPosition = i
		}
	}

	return cost[minPosition]
}

// used for part 1
func constantFuel(min, position int) int {
	return int(math.Abs(float64(min - position)))
}

// used for part 2
func progressiveFuel(min, position int) int {
	n := int(math.Abs(float64(min - position)))

	return n * (n + 1) / 2
}
