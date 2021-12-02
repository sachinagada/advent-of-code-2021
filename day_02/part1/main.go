package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	up      = "up"
	down    = "down"
	forward = "forward"
)

func main() {
	r, readErr := os.Open("../input.txt")
	if readErr != nil {
		panic("error reading input file")
	}

	scanner := bufio.NewScanner(r)

	var depth int64
	var horizontal int64
	for scanner.Scan() {
		line := scanner.Text()
		dir := strings.Split(line, " ")
		units, parseErr := strconv.ParseInt(dir[1], 10, 64)
		if parseErr != nil {
			panic("unexpected input: " + line)
		}

		switch string(dir[0]) {
		case up:
			depth -= units
		case down:
			depth += units
		case forward:
			horizontal += units
		}
	}

	fmt.Println("product:", depth*horizontal)
}
