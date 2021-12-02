package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

const (
	up      = "up"
	down    = "down"
	forward = "forward"
)

type position struct {
	depth      int
	horizontal int
	aim        int
}

func main() {
	scanner := helper.ScanFile("./input.txt")

	p := &position{}
	// p.calculatePosition(scanner)
	p.calculateWithAim(scanner)

	fmt.Println("product:", p.depth*p.horizontal)
}

// part1
func (p *position) calculatePosition(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		dir := strings.Split(line, " ")
		units := helper.ParseInt(dir[1])

		switch string(dir[0]) {
		case up:
			p.depth -= units
		case down:
			p.depth += units
		case forward:
			p.horizontal += units
		}
	}
}

// part2
func (p *position) calculateWithAim(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		dir := strings.Split(line, " ")
		units := helper.ParseInt(dir[1])

		switch string(dir[0]) {
		case up:
			p.aim -= units
		case down:
			p.aim += units
		case forward:
			p.horizontal += units
			p.depth += units * p.aim
		}
	}
}
