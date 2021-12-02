package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r, readErr := os.Open("../input.txt")
	if readErr != nil {
		panic("error reading input file")
	}

	scanner := bufio.NewScanner(r)

	// can either use holdArray or slidingWindow to compute the count. Tried
	// doint it two different ways.
	count := holdArray(scanner)

	fmt.Printf("number of increases: %d\n", count)
}

// slidingWindow keeps track of the sum of the sliding window and compares it
// with the previous window's sum. It doesn't hold the whole input in memory
// so this would work even if the input was very large.
func slidingWindow(scanner *bufio.Scanner) int {
	wind := make([]int, 0, 3) // basically acts like a queue
	prev := 0                 // previous sum
	count := 0                // number of times the sliding window has increased
	for scanner.Scan() {
		l := scanner.Text()
		cur, parseErr := strconv.ParseInt(l, 10, 64)
		if parseErr != nil {
			panic("non integer found on line")
		}

		if len(wind) < 3 {
			wind = append(wind, int(cur))
			prev += int(cur)
		} else {
			wind = append(wind[1:], int(cur))

			curSum := sum(wind)
			if curSum > prev {
				count++
			}

			prev = curSum
		}
	}

	return count
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

// holdArray holds all the values from the text file in an array and compares
// the current value (at the end of the array) with the value 3 indexes away.
// This is because the middle two values will be the same between consecutive
// sliding windows (for the first window, the first value will be different and
// for the second window, the last value will be different). If the second
// window's last value is greater than the first window's first value, the sum
// will be greater too, so saving some work by not computing the sum.
// Note, that depending on the size of the input, holding the whole file in
// memory would sometimes not be possible, at which point the above method is
// better.
func holdArray(scanner *bufio.Scanner) int {
	arr := []int{}
	count := 0

	for scanner.Scan() {
		l := scanner.Text()
		cur, parseErr := strconv.ParseInt(l, 10, 64)
		if parseErr != nil {
			panic("non integer found on line")
		}

		arr = append(arr, int(cur))

		if len(arr) < 4 {
			continue
		}

		if arr[len(arr)-1] > arr[len(arr)-4] {
			count++
		}
	}
	return count
}
