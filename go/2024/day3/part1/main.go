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
	for _, line := range input {
		lineSum, err := doMuls(line)
		if err != nil {
			return "", fmt.Errorf("error processing line: %s", err)
		}
		totalSum += lineSum
	}
	return fmt.Sprintf("Total: %d", totalSum), nil
}

func doMuls(line string) (int, error) {
	sum := 0
	s := &stateTracker{}
	for _, c := range line {
		switch c {
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
				s.state = IN_OPEN_PAREN
			} else {
				s.reset()
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if s.state == IN_OPEN_PAREN {
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
			if s.state == IN_SECOND_NUM {
				res, err := s.doMul()
				if err != nil {
					return 0, fmt.Errorf("error multiplying: %s", err)
				}
				sum += res
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
	RESET            state = iota
	IN_M             state = iota
	IN_U             state = iota
	IN_L             state = iota
	IN_OPEN_PAREN    state = iota
	IN_FIRST_NUM     state = iota
	IN_COMMA         state = iota
	IN_SECOND_NUM    state = iota
	IN_CLOSING_PAREN state = iota
)
