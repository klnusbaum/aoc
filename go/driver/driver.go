package driver

import (
	"bufio"
	"fmt"
	"os"
)

type Solver interface {
	Solve(input []string) (string, error)
}

type Driver struct {
	InFile string
	Solver Solver
}

func (d Driver) Solve() error {
	i, err := d.input()
	if err != nil {
		return fmt.Errorf("failed to collect input: %s", err)
	}

	result, err := d.Solver.Solve(i)
	if err != nil {
		return fmt.Errorf("failed to solve: %s", err)
	}

	fmt.Println(result)
	return nil
}

func (d Driver) input() ([]string, error) {
	f, err := os.Open(d.InFile)
	if err != nil {
		return nil, fmt.Errorf("couldn't open input file: %s", err)
	}

	sc := bufio.NewScanner(f)
	var result []string
	for sc.Scan() {
		result = append(result, sc.Text())
	}

	if sc.Err() != nil {
		return nil, fmt.Errorf("scanner failed: %s", err)
	}

	return result, nil
}
