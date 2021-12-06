package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func main() {
	scanner := helper.ScanFile("./input.txt")
	count := countFish(scanner, 80)
	fmt.Println("number of fish:", count)

}

func countFish(scanner *bufio.Scanner, days int) int {
	// lf holds the number of fish at each internal timer. So index 0
	// holds the number of fish with internal timer of 0.
	lf := make([]int, 9)
	for scanner.Scan() {
		l := scanner.Text()
		sArr := strings.Split(l, ",")

		for _, s := range sArr {
			intTimer := helper.ParseInt(s)
			lf[intTimer]++
		}
	}

	for i := 0; i < days; i++ {
		// as each day progresses, the internal timer decreases by a day so
		// shift the slice. The ones at internal timer of 0, would give birth
		// to a new fish each so that would be the addition to index 8. The
		// parents that gave birth would reset their timer to day 6 so add the
		// parent count to day 6 count.
		newLF := append(lf[1:], lf[0])
		newLF[6] += lf[0]
		lf = newLF
	}

	sum := 0
	for _, val := range lf {
		sum += val
	}

	return sum
}
