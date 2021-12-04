package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	"github.com/sachinagada/advent-of-code-2021/helper"
)

func main() {
	scanner := helper.ScanFile("./input.txt")
	score := playBingo(scanner, highestBoard)
	fmt.Println("score", score)
}

type bingo struct {
	board  [5][5]int
	marked [5][5]string
}

func NewBingoBoard() *bingo {
	b := &bingo{
		board:  [5][5]int{},
		marked: [5][5]string{},
	}

	for i := 0; i < 5; i++ {
		b.marked[i] = [5]string{}
	}

	return b
}

func (b *bingo) parse(l string, index int) {
	arr := strings.Fields(l)
	row := [5]int{}
	for i, s := range arr {
		row[i] = helper.ParseInt(s)
	}

	b.board[index] = row
}

func (b *bingo) checkRow(i int) bool {
	for j := 0; j < len(b.marked); j++ {
		if b.marked[i][j] == "" {
			return false
		}
	}
	// not an empty string in whole line so is bingo
	return true
}

func (b *bingo) checkCol(j int) bool {
	for i := 0; i < len(b.marked); i++ {
		if b.marked[i][j] == "" {
			return false
		}
	}

	return true
}

func (b *bingo) isBingo() bool {
	for i := 0; i < len(b.marked); i++ {
		if b.checkRow(i) || b.checkCol(i) {
			return true
		}
	}
	return false
}

// returns if it got bingo and the index of input where it was able to get bingo
func (b *bingo) checkWinner(input []int) (int, bool) {
	for in, v := range input {
	BOARD:
		for i := 0; i < len(b.board); i++ {
			for j := 0; j < len(b.board); j++ {
				if b.board[i][j] == v {
					b.marked[i][j] = "x"
					if b.isBingo() {
						return in, true
					}
					break BOARD
				}
			}
		}
	}
	return 0, false
}

func (b *bingo) score(multiplier int) int {
	sum := 0
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board); j++ {
			if b.marked[i][j] == "" {
				sum += b.board[i][j]
			}
		}
	}

	return sum * multiplier
}

func playBingo(scanner *bufio.Scanner, chooseBoard func([]int) int) int {
	firstLine := true
	var input []int
	var bingoBoards []*bingo

	b := NewBingoBoard()
	index := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			continue
		}

		if firstLine {
			input = drawnNumbers(strings.Split(l, ","))
			// insert input in first line
			firstLine = false
			continue
		}

		b.parse(l, index)
		index++

		if index%5 == 0 {
			bingoBoards = append(bingoBoards, b)
			b = NewBingoBoard()
			index = 0
		}
	}

	// winner holds the index for the winning Bingo number in each board
	winners := make([]int, len(bingoBoards))
	for i, bb := range bingoBoards {
		if inputIndex, won := bb.checkWinner(input); won {
			winners[i] = inputIndex
		}
	}

	chosenIndex := chooseBoard(winners)
	winningBoard := bingoBoards[chosenIndex]

	inputIndex := winners[chosenIndex]
	multiplier := input[inputIndex]

	score := winningBoard.score(multiplier)
	return score
}

// parse the random numbers drawn for bingo
func drawnNumbers(input []string) []int {
	output := make([]int, len(input))
	for i, s := range input {
		output[i] = helper.ParseInt(s)
	}
	return output
}

// used for part 1
// returns index of bingo board with lowest number needed to be drawn to win
func lowestBoard(arr []int) int {
	minimum := math.MaxInt16
	index := 0
	for i, v := range arr {
		if v > 0 && v < minimum {
			minimum = v
			index = i
		}
	}
	return index
}

// use highestBoard for part 2. Returns the index of the board that requires
// the most number of numbers to be picked for the board to win
func highestBoard(arr []int) int {
	max := -1
	index := 0
	for i, v := range arr {
		if v > 0 && v > max {
			max = v
			index = i
		}
	}
	return index
}
