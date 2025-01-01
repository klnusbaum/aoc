package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/klnusbaum/aoc/go/driver"
)

func main() {
	d := driver.Driver{
		Solver: D1P2Solver{},
	}

	if err := d.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "error solving: %s", err)
		os.Exit(1)
	}
}

type D1P2Solver struct{}

func (s D1P2Solver) Solve(input []string) (string, error) {
	l1, l2, err := numSlices(input)
	if err != nil {
		return "", fmt.Errorf("failed to get number lists: %s", err)
	}

	instances := make(map[int]int)
	for v := range slices.Values(l2) {
		instances[v] = instances[v] + 1
	}

	sum := 0
	for v := range slices.Values(l1) {
		amount := instances[v]
		sum += amount * v
	}

	return fmt.Sprintf("Similarity is %d", sum), nil

}

func numSlices(input []string) ([]int, []int, error) {
	var l1, l2 []int
	for line := range slices.Values(input) {
		vals := strings.Fields(line)
		n1, err := strconv.Atoi(vals[0])
		if err != nil {
			return nil, nil, errors.New("invalid number")
		}
		n2, err := strconv.Atoi(vals[1])
		if err != nil {
			return nil, nil, errors.New("invalid number")
		}
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}
	return l1, l2, nil
}
