package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/klnusbaum/aoc/go/driver"
)

func main() {
	d := driver.Driver{
		Solver: D6P1Solver{},
	}

	if err := d.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "error solving: %s", err)
		os.Exit(1)
	}
}

type D6P1Solver struct{}

func (s D6P1Solver) Solve(input []string) (string, error) {
	a := parseArea(input)

	curDir := UP
	curX, curY := curLocation(a)
	locationTracker := newLocationTracker()
	locationTracker.logLocation(curX, curY)
	for {
		nextX, nextY := curX, curY
		switch curDir {
		case UP:
			nextX = nextX - 1
		case DOWN:
			nextX = nextX + 1
		case LEFT:
			nextY = nextY - 1
		case RIGHT:
			nextY = nextY + 1
		}

		if nextY < 0 || nextY >= len(a) {
			break
		}

		if nextX < 0 || nextX >= len(a[0]) {
			break
		}

		if a[nextX][nextY] == '#' {
			curDir = nextDir(curDir)
		} else {
			locationTracker.logLocation(nextX, nextY)
			curX, curY = nextX, nextY
		}
	}

	return fmt.Sprintf("Total locations: %d", locationTracker.total()), nil
}

func parseArea(input []string) [][]rune {
	var res [][]rune

	for line := range slices.Values(input) {
		res = append(res, []rune(line))
	}
	return res
}

func curLocation(a [][]rune) (int, int) {
	for x, row := range a {
		for y, r := range row {
			if r == '^' {
				return x, y
			}
		}
	}

	return 0, 0
}

func nextDir(curDir direction) direction {
	switch curDir {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	default:
		return UP
	}
}

type area [][]rune

type locationTracker map[int]map[int]bool

func newLocationTracker() locationTracker {
	return make(map[int]map[int]bool)
}

func (t locationTracker) logLocation(x, y int) {
	if t[x] == nil {
		t[x] = make(map[int]bool)
	}

	t[x][y] = true
}

func (t locationTracker) total() int {
	total := 0
	for _, row := range t {
		total += len(row)
	}
	return total
}

type direction int

const (
	UP    direction = iota
	DOWN  direction = iota
	LEFT  direction = iota
	RIGHT direction = iota
)
