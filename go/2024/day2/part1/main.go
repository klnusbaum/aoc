package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/klnusbaum/aoc/go/driver"
)

func main() {
	d := driver.Driver{
		Solver: D2P1Solver{},
	}

	if err := d.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "error solving: %s", err)
		os.Exit(1)
	}
}

type D2P1Solver struct{}

func (s D2P1Solver) Solve(input []string) (string, error) {
	numSafe := 0

	for line := range slices.Values(input) {
		r, err := newReport(line)
		if err != nil {
			return "", fmt.Errorf("couldn't make report: %s", err)
		}

		if r.isSafe() {
			numSafe++
		}
	}

	return fmt.Sprintf("Num safe: %d", numSafe), nil
}

type report []int

func newReport(line string) (report, error) {
	var result []int
	levels := strings.Fields(line)

	for v := range slices.Values(levels) {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse report: %s", err)
		}
		result = append(result, i)
	}

	return result, nil
}

func (r report) isSafe() bool {
	return r.isGradualAscending() || r.isGradualDescending()
}

func (r report) isGradualAscending() bool {
	if !slices.IsSorted(r) {
		return false
	}

	for i := range len(r) - 1 {
		if !(r[i+1] >= r[i]+1 && r[i+1] <= r[i]+3) {
			return false
		}
	}

	return true
}

func (r report) isGradualDescending() bool {
	if !slices.IsSortedFunc(r, func(a, b int) int {
		return b - a
	}) {
		return false
	}

	for i := range len(r) - 1 {
		if !(r[i+1] <= r[i]-1 && r[i+1] >= r[i]-3) {
			return false
		}
	}

	return true
}
