package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/klnusbaum/aoc/go/driver"
)

func main() {
	d := driver.Driver{
		Solver: D4P2Solver{},
	}

	if err := d.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "error solving: %s", err)
		os.Exit(1)
	}
}

type D4P2Solver struct{}

func (s D4P2Solver) Solve(input []string) (string, error) {
	var chars [][]rune
	for line := range slices.Values(input) {
		var lineChars []rune
		for _, c := range line {
			lineChars = append(lineChars, c)
		}
		chars = append(chars, lineChars)
	}

	return fmt.Sprintf("Num xmas: %d", getCount(chars)), nil
}

func getCount(chars [][]rune) int {
	numXMAS := 0
	for i := 1; i < len(chars)-1; i++ {
		for j := 1; j < len(chars[i])-1; j++ {
			if chars[i][j] == 'A' && isX(chars, i, j) {
				numXMAS++
			}
		}
	}
	return numXMAS
}

func isX(chars [][]rune, row, col int) bool {
	return isConf1(chars, row, col) ||
		isConf2(chars, row, col) ||
		isConf3(chars, row, col) ||
		isConf4(chars, row, col)
}

func isConf1(chars [][]rune, row, col int) bool {
	return chars[row-1][col-1] == 'M' && chars[row-1][col+1] == 'M' &&
		chars[row+1][col-1] == 'S' && chars[row+1][col+1] == 'S'
}

func isConf2(chars [][]rune, row, col int) bool {
	return chars[row-1][col-1] == 'S' && chars[row-1][col+1] == 'S' &&
		chars[row+1][col-1] == 'M' && chars[row+1][col+1] == 'M'
}

func isConf3(chars [][]rune, row, col int) bool {
	return chars[row-1][col-1] == 'S' && chars[row-1][col+1] == 'M' &&
		chars[row+1][col-1] == 'S' && chars[row+1][col+1] == 'M'
}

func isConf4(chars [][]rune, row, col int) bool {
	return chars[row-1][col-1] == 'M' && chars[row-1][col+1] == 'S' &&
		chars[row+1][col-1] == 'M' && chars[row+1][col+1] == 'S'
}
