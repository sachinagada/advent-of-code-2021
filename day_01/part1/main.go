package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	r, readErr := os.Open("../input.txt")
	if readErr != nil {
		panic("error reading input file")
	}

	scanner := bufio.NewScanner(r)

	prev := int64(math.MaxInt64)
	count := 0
	for scanner.Scan() {
		l := scanner.Text()
		cur, parseErr := strconv.ParseInt(l, 10, 64)
		if parseErr != nil {
			panic("non integer found on line")
		}

		if cur > prev {
			count++
		}

		prev = cur
	}
	fmt.Printf("number of inceases: %d\n", count)
}
