package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/klnusbaum/aoc/go/driver"
)

func main() {
	d := driver.Driver{
		Solver: D3P1Solver{},
	}

	if err := d.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "error solving: %s", err)
		os.Exit(1)
	}
}

type D3P1Solver struct{}

func (s D3P1Solver) Solve(input []string) (string, error) {
	totalSum := 0
	st := &stateTracker{
		enabled: true,
	}
	for _, line := range input {
		lineSum, err := doMuls(line, st)
		if err != nil {
			return "", fmt.Errorf("error processing line: %s", err)
		}
		totalSum += lineSum
	}
	return fmt.Sprintf("Total: %d", totalSum), nil
}

func doMuls(line string, s *stateTracker) (int, error) {
	sum := 0
	for _, c := range line {
		switch c {
		case 'd':
			s.state = IN_D
		case 'o':
			if s.state == IN_D {
				s.state = IN_O
			} else {
				s.reset()
			}
		case 'n':
			if s.state == IN_O {
				s.state = IN_N
			} else {
				s.reset()
			}
		case '\'':
			if s.state == IN_N {
				s.state = IN_APPOSTROPHE
			} else {
				s.reset()
			}
		case 't':
			if s.state == IN_APPOSTROPHE {
				s.state = IN_T
			} else {
				s.reset()
			}
		case 'm':
			s.state = IN_M
		case 'u':
			if s.state == IN_M {
				s.state = IN_U
			} else {
				s.reset()
			}
		case 'l':
			if s.state == IN_U {
				s.state = IN_L
			} else {
				s.reset()
			}
		case '(':
			if s.state == IN_L {
				s.state = IN_MUL_OPEN_PAREN
			} else if s.state == IN_O {
				s.state = IN_DO_OPEN_PAREN
			} else if s.state == IN_T {
				s.state = IN_DONT_OPEN_PAREN
			} else {
				s.reset()
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if s.state == IN_MUL_OPEN_PAREN {
				s.state = IN_FIRST_NUM
				s.firstNum += string(c)
			} else if s.state == IN_FIRST_NUM {
				s.firstNum += string(c)
			} else if s.state == IN_COMMA {
				s.state = IN_SECOND_NUM
				s.secondNum += string(c)
			} else if s.state == IN_SECOND_NUM {
				s.secondNum += string(c)
			} else {
				s.reset()
			}
		case ',':
			if s.state == IN_FIRST_NUM {
				s.state = IN_COMMA
			} else {
				s.reset()
			}
		case ')':
			if s.state == IN_SECOND_NUM && s.enabled {
				res, err := s.doMul()
				if err != nil {
					return 0, fmt.Errorf("error multiplying: %s", err)
				}
				sum += res
			} else if s.state == IN_DO_OPEN_PAREN {
				s.enabled = true
			} else if s.state == IN_DONT_OPEN_PAREN {
				s.enabled = false
			}
			s.reset()
		default:
			s.reset()
		}
	}

	return sum, nil
}

type stateTracker struct {
	state     state
	enabled   bool
	firstNum  string
	secondNum string
}

func (t *stateTracker) reset() {
	t.state = RESET
	t.firstNum = ""
	t.secondNum = ""
}

func (t *stateTracker) doMul() (int, error) {
	a, err := strconv.Atoi(t.firstNum)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse first number: %s", err)
	}
	b, err := strconv.Atoi(t.secondNum)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse first number: %s", err)
	}

	return a * b, nil
}

type state int

const (
	RESET state = iota

	IN_D           state = iota
	IN_O           state = iota
	IN_N           state = iota
	IN_APPOSTROPHE state = iota
	IN_T           state = iota

	IN_M          state = iota
	IN_U          state = iota
	IN_L          state = iota
	IN_FIRST_NUM  state = iota
	IN_COMMA      state = iota
	IN_SECOND_NUM state = iota

	IN_MUL_OPEN_PAREN  state = iota
	IN_DO_OPEN_PAREN   state = iota
	IN_DONT_OPEN_PAREN state = iota
	IN_CLOSING_PAREN   state = iota
)
