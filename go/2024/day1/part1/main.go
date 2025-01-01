package part1

import (
	"fmt"
	"os"

	"github.com/klnusbaum/aoc/go/driver"
)

func main() {
	d := driver.Driver{
		InFile: "input.txt",
		Solver: D1P1Solver{},
	}

	if err := d.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "error solving: %s", err)
		os.Exit(1)
	}
}

type D1P1Solver struct{}

func (s D1P1Solver) Solve(input []string) error {

	return nil
}
