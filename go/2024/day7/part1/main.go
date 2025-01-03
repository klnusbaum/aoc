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
		Solver: D7P1Solver{},
	}

	if err := d.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "error solving: %s", err)
		os.Exit(1)
	}
}

type D7P1Solver struct{}

func (s D7P1Solver) Solve(input []string) (string, error) {
	var equations []equation
	for line := range slices.Values(input) {
		eq, err := parseEquation(line)
		if err != nil {
			return "", fmt.Errorf("couldn't parse equation: %s", err)
		}
		equations = append(equations, eq)
	}

	var trueEqs []equation
	for eq := range slices.Values(equations) {
		if eq.IsTrue() {
			trueEqs = append(trueEqs, eq)
		}
	}

	sum := 0
	for eq := range slices.Values(trueEqs) {
		sum += eq.target
	}
	return fmt.Sprintf("Total of valid equations: %d", sum), nil
}

type equation struct {
	target int
	vals   []int
}

func (e equation) IsTrue() bool {
	pvs := possibleValues(e.vals)

	for pv := range slices.Values(pvs) {
		if pv == e.target {
			return true
		}
	}
	return false
}

func possibleValues(vals []int) []int {
	if len(vals) == 1 {
		return []int{vals[0]}
	}

	var res []int

	pvs := possibleValues(vals[:len(vals)-1])
	for pv := range slices.Values(pvs) {
		res = append(res, pv+vals[len(vals)-1])
		res = append(res, pv*vals[len(vals)-1])
	}

	return res
}

func parseEquation(input string) (equation, error) {
	chunks := strings.Split(input, ":")
	target, err := strconv.Atoi(chunks[0])
	if err != nil {
		return equation{}, fmt.Errorf("couldn't parse target value: %s", err)
	}

	valFields := strings.Fields(chunks[1])
	var vals []int
	for v := range slices.Values(valFields) {
		val, err := strconv.Atoi(v)
		if err != nil {
			return equation{}, fmt.Errorf("couldn't parse value: %s", err)
		}
		vals = append(vals, val)
	}

	return equation{
		target: target,
		vals:   vals,
	}, nil
}
