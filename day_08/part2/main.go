package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func main() {
	scanner := helper.ScanFile("../input.txt")
	sum := searchSegment(scanner)
	fmt.Println(sum)
}

func searchSegment(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		l := scanner.Text()

		arr := strings.Split(l, " | ")
		input := strings.Fields(arr[0])
		output := strings.Fields(arr[1])

		numMap := determineNumbers(input)
		number := calculateOutput(numMap, output)
		sum += number
	}
	return sum
}

func determineNumbers(input []string) map[int]string {
	sort.Slice(input, func(i, j int) bool {
		return len(input[i]) < len(input[j])
	})
	numMap := map[int]string{
		1: input[0], // 1 has 2 letters
		7: input[1], // 7 has 3 letters
		4: input[2], // 4 has 4 letters
		8: input[9], // 8 has 7 letters
	}

	sixLetterInd := make([]int, 0, 2) // will be index for input for 0 and 6
	for i := 6; i <= 8; i++ {
		if isSubstring(input[i], numMap[4]) {
			numMap[9] = input[i]
			continue
		}
		sixLetterInd = append(sixLetterInd, i)
	}

	for _, ind := range sixLetterInd {
		// 0 contains all letters of 7
		if isSubstring(input[ind], numMap[7]) {
			numMap[0] = input[ind]
		} else {
			// 6 is the last remaining 6 letter signal
			numMap[6] = input[ind]
		}
	}

	fiveLetterInd := make([]int, 0, 2) // will be index for input for 2 and 5
	for i := 3; i <= 5; i++ {
		if isSubstring(input[i], numMap[1]) {
			// the signal for 3 should have all letters as 1
			numMap[3] = input[i]
			continue
		}
		fiveLetterInd = append(fiveLetterInd, i)
	}

	for _, ind := range fiveLetterInd {
		// singal for 6 encompasses all the letters of 5 (but not 2)
		if isSubstring(numMap[6], input[ind]) {
			numMap[5] = input[ind]
		} else {
			numMap[2] = input[ind]
		}
	}

	return numMap
}

// checks to see if string s has all the letters of substring
func isSubstring(s, sub string) bool {
SUBSTRING:
	for _, subR := range sub {
		for _, r := range s {
			if r == subR {
				continue SUBSTRING
			}
		}
		// went through the whole string and didn't find a similar char as substring
		return false

	}
	return true
}

func calculateOutput(numMap map[int]string, output []string) int {
	num := make([]int, 4)

	for i, v := range output {
		switch {
		case len(v) == 2:
			num[i] = 1
		case len(v) == 3:
			num[i] = 7
		case len(v) == 4:
			num[i] = 4
		case len(v) == 7:
			num[i] = 8
		default:
			for k, numLetters := range numMap {
				switch k {
				case 1, 4, 7, 8:
					continue
				default:
					if isEqual(numLetters, v) {
						num[i] = k
					}
				}
			}
		}
	}

	return getDigit(num)
}

func isEqual(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	r1 := strings.Split(s1, "")
	sort.Strings(r1)

	r2 := strings.Split(s2, "")
	sort.Strings(r2)

	for i, v := range r1 {
		if r2[i] != v {
			return false
		}
	}

	return true
}

func getDigit(output []int) int {
	sum := 0
	for _, v := range output {
		sum = sum*10 + v
	}
	return sum
}
