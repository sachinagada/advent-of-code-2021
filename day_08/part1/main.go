package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func main() {
	scanner := helper.ScanFile("../input.txt")
	output := parseInput(scanner)
	count := countOutput(output)
	fmt.Println(count)
}

func parseInput(scanner *bufio.Scanner) []string {
	var output []string
	for scanner.Scan() {
		l := scanner.Text()

		arr := strings.Split(l, " | ")
		output = append(output, strings.Fields(arr[1])...)
	}
	return output
}

func countOutput(output []string) int {
	count := 0
	for _, s := range output {
		switch len(s) {
		case 2, 3, 4, 7:
			count++
		default:
		}
	}
	return count
}
