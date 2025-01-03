package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/klnusbaum/aoc/go/driver"
)

func main() {
	d := driver.Driver{
		Solver: D4P1Solver{},
	}

	if err := d.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "error solving: %s", err)
		os.Exit(1)
	}
}

type D4P1Solver struct{}

func (s D4P1Solver) Solve(input []string) (string, error) {
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
	for i, row := range chars {
		for j, c := range row {
			if c == 'X' {
				numXMAS += countFrom(chars, i, j)
			}
		}
	}
	return numXMAS
}

func countFrom(chars [][]rune, row, col int) int {
	numXMAS := 0

	if checkForward(chars, row, col) {
		numXMAS++
	}
	if checkBackwards(chars, row, col) {
		numXMAS++
	}
	if checkUp(chars, row, col) {
		numXMAS++
	}
	if checkDown(chars, row, col) {
		numXMAS++
	}
	if checkNW(chars, row, col) {
		numXMAS++
	}
	if checkNE(chars, row, col) {
		numXMAS++
	}
	if checkSE(chars, row, col) {
		numXMAS++
	}
	if checkSW(chars, row, col) {
		numXMAS++
	}

	return numXMAS
}

func checkForward(chars [][]rune, row, col int) bool {
	if col+3 >= len(chars[row]) {
		return false
	}

	return chars[row][col+1] == 'M' && chars[row][col+2] == 'A' && chars[row][col+3] == 'S'
}

func checkBackwards(chars [][]rune, row, col int) bool {
	if col-3 < 0 {
		return false
	}

	return chars[row][col-1] == 'M' && chars[row][col-2] == 'A' && chars[row][col-3] == 'S'
}

func checkUp(chars [][]rune, row, col int) bool {
	if row-3 < 0 {
		return false
	}

	return chars[row-1][col] == 'M' && chars[row-2][col] == 'A' && chars[row-3][col] == 'S'
}

func checkDown(chars [][]rune, row, col int) bool {
	if row+3 >= len(chars) {
		return false
	}

	return chars[row+1][col] == 'M' && chars[row+2][col] == 'A' && chars[row+3][col] == 'S'
}

func checkNE(chars [][]rune, row, col int) bool {
	if row-3 < 0 || col-3 < 0 {
		return false
	}

	return chars[row-1][col-1] == 'M' && chars[row-2][col-2] == 'A' && chars[row-3][col-3] == 'S'
}

func checkNW(chars [][]rune, row, col int) bool {
	if row-3 < 0 || col+3 >= len(chars[row]) {
		return false
	}

	return chars[row-1][col+1] == 'M' && chars[row-2][col+2] == 'A' && chars[row-3][col+3] == 'S'
}

func checkSE(chars [][]rune, row, col int) bool {
	if row+3 >= len(chars) || col+3 >= len(chars[row]) {
		return false
	}

	return chars[row+1][col+1] == 'M' && chars[row+2][col+2] == 'A' && chars[row+3][col+3] == 'S'
}

func checkSW(chars [][]rune, row, col int) bool {
	if row+3 >= len(chars) || col-3 < 0 {
		return false
	}

	return chars[row+1][col-1] == 'M' && chars[row+2][col-2] == 'A' && chars[row+3][col-3] == 'S'
}
