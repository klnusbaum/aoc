package main

import (
	"fmt"
	"os"
	"slices"
	"sync"

	"github.com/klnusbaum/aoc/go/driver"
)

func main() {
	d := driver.Driver{
		Solver: D6P2Solver{},
	}

	if err := d.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "error solving: %s", err)
		os.Exit(1)
	}
}

type D6P2Solver struct{}

func (s D6P2Solver) Solve(input []string) (string, error) {
	a := parseArea(input)
	startX, startY := curLocation(a)
	numLoops := 0
	var numLoopsMu sync.Mutex
	var numLoopsGroup sync.WaitGroup

	for i, row := range a {
		for j := range row {
			if i == startX && j == startY {
				// can't place new obstacle in start position
				continue
			}

			if a[i][j] == '#' {
				// location already has an obstacle in it
				continue
			}

			numLoopsGroup.Add(1)
			go func(sa area, sx, sy, tx, ty int) {
				defer numLoopsGroup.Done()
				if hasLoop(a, startX, startY, i, j) {
					numLoopsMu.Lock()
					numLoops++
					numLoopsMu.Unlock()
				}
			}(a, startX, startY, i, j)
		}
	}

	numLoopsGroup.Wait()
	return fmt.Sprintf("Total loops: %d", numLoops), nil
}

func hasLoop(a area, startX, startY, tryX, tryY int) bool {
	curDir := UP
	curX, curY := startX, startY
	locationTracker := newLocationTracker()
	locationTracker.logVector(curX, curY, UP)
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

		if a[nextX][nextY] == '#' || (nextX == tryX && nextY == tryY) {
			curDir = nextDir(curDir)
		} else {
			curX, curY = nextX, nextY
		}

		if locationTracker.hasVector(curX, curY, curDir) {
			return true
		} else {
			locationTracker.logVector(curX, curY, curDir)
		}
	}

	return false
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

type locationTracker map[int]map[int]map[direction]bool

func newLocationTracker() locationTracker {
	return make(map[int]map[int]map[direction]bool)
}

func (t locationTracker) logVector(x, y int, dir direction) {
	if t[x] == nil {
		t[x] = make(map[int]map[direction]bool)
	}

	if t[x][y] == nil {
		t[x][y] = make(map[direction]bool)
	}

	t[x][y][dir] = true
}

func (t locationTracker) hasVector(x, y int, dir direction) bool {
	if t[x] == nil || t[x][y] == nil {
		return false
	}

	return t[x][y][dir]
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
