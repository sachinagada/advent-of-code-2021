package main

import (
	"bufio"
	"fmt"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

type bits struct {
	zeros int
	ones  int
}

func (b *bits) increment(r rune) {
	switch r {
	case '0':
		b.zeros++
	case '1':
		b.ones++
	default:
		panic("invalid bit")
	}
}

func (b *bits) mostCommon() byte {
	if b.zeros > b.ones {
		return '0'
	}
	return '1'
}

func (b *bits) leastCommon() byte {
	if b.ones < b.zeros {
		return '1'
	}
	return '0'
}

func main() {
	scanner := helper.ScanFile("../input.txt")
	ls := lifeSupport(scanner)
	fmt.Printf("life support rating: %d\n", ls)
}

func lifeSupport(scanner *bufio.Scanner) int {
	diagnostics := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		diagnostics = append(diagnostics, line)
	}

	o2Gen := calculate(diagnostics, func(b *bits) byte {
		return b.mostCommon()
	})

	co2Scrub := calculate(diagnostics, func(b *bits) byte {
		return b.leastCommon()
	})

	return o2Gen * co2Scrub
}

func calculate(arr []string, bitCriteria func(b *bits) byte) int {
	for i := 0; i < len(arr[0]); i++ {
		b := &bits{}
		for _, a := range arr {
			b.increment(rune(a[i]))
		}

		// Now keep numbers with bitCriteria
		startingVal := bitCriteria(b)
		subset := make([]string, 0, len(arr))
		for _, a := range arr {
			if a[i] == startingVal {
				subset = append(subset, a)
			}
		}

		arr = subset
		if len(subset) == 1 {
			break
		}
	}

	return helper.ParseBinary(arr[0])
}
