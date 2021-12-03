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

type rate struct {
	gamma   []byte // most common bits
	epsilon []byte // least common bits
}

func (r *rate) calculate(bArr []*bits) {
	for i, b := range bArr {
		r.gamma[i] = b.mostCommon()
		r.epsilon[i] = b.leastCommon()
	}
}

func (r *rate) powerConsumption() int {
	gamma := helper.ParseBinary(string(r.gamma))
	epsilon := helper.ParseBinary(string(r.epsilon))

	return gamma * epsilon
}

func main() {
	scanner := helper.ScanFile("../input.txt")
	pc := calculatePowerConsumption(scanner)
	fmt.Printf("power consumption: %d\n", pc)
}

func calculatePowerConsumption(scanner *bufio.Scanner) int {
	var arr []*bits

	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			arr = make([]*bits, len(line))
			for i := range arr {
				arr[i] = &bits{}
			}

			firstLine = false
		}

		for i, r := range line {
			arr[i].increment(r)
		}
	}

	r := &rate{
		gamma:   make([]byte, len(arr)),
		epsilon: make([]byte, len(arr)),
	}

	r.calculate(arr)
	return r.powerConsumption()

}
