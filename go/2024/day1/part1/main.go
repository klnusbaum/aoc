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
		Solver: D1P1Solver{},
	}

	if err := d.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "error solving: %s", err)
		os.Exit(1)
	}
}

type D1P1Solver struct{}

func (s D1P1Solver) Solve(input []string) (string, error) {
	l1, l2, err := numSlices(input)
	if err != nil {
		return "", fmt.Errorf("failed to get number lists: %s", err)
	}

	slices.Sort(l1)
	slices.Sort(l2)

	sum := 0
	for i := range l1 {
		if l1[i] > l2[i] {
			sum += l1[i] - l2[i]
		} else {
			sum += l2[i] - l1[i]
		}
	}

	return fmt.Sprintf("Sum of differences is %d", sum), nil
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
