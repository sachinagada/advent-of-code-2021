package main

import (
	"fmt"
	"math"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func main() {
	scanner := helper.ScanFile("../input.txt")

	prev := math.MaxInt64
	count := 0
	for scanner.Scan() {
		l := scanner.Text()
		cur := helper.ParseInt(l)

		if cur > prev {
			count++
		}

		prev = cur
	}
	fmt.Printf("number of inceases: %d\n", count)
}
