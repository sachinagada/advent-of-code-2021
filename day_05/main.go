package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func main() {
	scanner := helper.ScanFile("./input.txt")
	overlap := navigateVents(scanner)
	fmt.Println("overlap:", overlap)

}

func navigateVents(scanner *bufio.Scanner) int {
	v := make([][]int, 1)
	v[0] = make([]int, 1)

	for scanner.Scan() {
		l := scanner.Text()
		endpoints := strings.Split(l, " -> ")
		e1 := NewCoordinate(endpoints[0])
		e2 := NewCoordinate(endpoints[1])

		v = markVent(v, e1, e2)
	}

	overLap := countOverlaps(v)
	return overLap
}

func markVent(v [][]int, e1, e2 coordinates) [][]int {
	if e1.x == e2.x {
		v = markVertical(v, e1, e2)
		return v
	}

	v = markLine(v, e1, e2)
	return v
}

func markVertical(v [][]int, e1, e2 coordinates) [][]int {
	ys := []int{e1.y, e2.y}
	sort.Ints(ys)

	v = incrementRows(v, ys[1])
	v = incrementColumn(v, e1.x)

	for y := ys[0]; y <= ys[1]; y++ {
		v[y][e1.x]++
	}

	return v
}

func markLine(v [][]int, e1, e2 coordinates) [][]int {
	ys := []int{e1.y, e2.y}
	sort.Ints(ys)

	xs := []int{e1.x, e2.x}
	sort.Ints(xs)

	v = incrementColumn(v, xs[1])
	v = incrementRows(v, ys[1])

	l := calcLine(e1, e2)

	for x := xs[0]; x <= xs[1]; x++ {
		y := l.slope*x + l.b
		v[y][x]++
	}

	return v
}

func incrementRows(v [][]int, yIndex int) [][]int {
	reqLen := yIndex + 1
	vLen := len(v)
	if vLen < reqLen {
		add := make([][]int, reqLen-vLen)
		v = append(v, add...)

		// initialize each row
		for i := vLen; i < len(v); i++ {
			v[i] = make([]int, len(v[0]))
		}
	}

	return v
}

func incrementColumn(v [][]int, xIndex int) [][]int {
	reqLen := xIndex + 1
	cols := len(v[0])
	if cols < reqLen {
		for i := 0; i < len(v); i++ {
			add := make([]int, reqLen-cols)
			v[i] = append(v[i], add...)
		}
	}
	return v
}

func countOverlaps(v [][]int) int {
	sum := 0
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[0]); j++ {
			if v[i][j] >= 2 {
				sum++
			}
		}
	}

	return sum
}

type coordinates struct {
	x int
	y int
}

func NewCoordinate(s string) coordinates {
	arr := strings.Split(s, ",")
	x := helper.ParseInt(arr[0])
	y := helper.ParseInt(arr[1])

	return coordinates{
		x: x,
		y: y,
	}
}

type line struct {
	slope int
	b     int
}

func calcLine(c1, c2 coordinates) line {
	diffY := c2.y - c1.y
	diffX := c2.x - c1.x
	slope := diffY / diffX

	// y = mx + b
	// b = y - mx
	b := c2.y - slope*c2.x

	return line{
		slope: slope,
		b:     b,
	}
}
