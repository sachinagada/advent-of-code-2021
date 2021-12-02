package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, readErr := os.Open("../input.txt")
	if readErr != nil {
		panic("error reading input file")
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
