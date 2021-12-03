package helper

import (
	"bufio"
	"os"
	"strconv"
)

func ScanFile(path string) *bufio.Scanner {
	r, readErr := os.Open(path)
	if readErr != nil {
		panic("error reading input file")
	}

	scanner := bufio.NewScanner(r)
	return scanner
}

func ParseInt(s string) int {
	i, parseErr := strconv.ParseInt(s, 10, 64)
	if parseErr != nil {
		panic("non integer found on line")
	}

	return int(i)
}

func ParseBinary(s string) int {
	i, parseErr := strconv.ParseInt(s, 2, 64)
	if parseErr != nil {
		panic("unexpected error:" + parseErr.Error())
	}

	return int(i)
}
