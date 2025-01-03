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
		Solver: D5P1Solver{},
	}

	if err := d.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "error solving: %s", err)
		os.Exit(1)
	}
}

type D5P1Solver struct{}

func (s D5P1Solver) Solve(input []string) (string, error) {
	middle := 0
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			middle = i
			break
		}
	}

	rules := parseRules(input[:middle])
	updates := parseUpdates(input[middle+1:])

	var invalidUpdates [][]string
	for update := range slices.Values(updates) {
		if !isValidUpdate(update, rules) {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	var fixedUpdates [][]string
	for update := range slices.Values(invalidUpdates) {
		fixedUpdates = append(fixedUpdates, fixUpdate(update, rules))
	}

	sum, err := addUpdateMiddles(fixedUpdates)
	if err != nil {
		return "", fmt.Errorf("couldn't sum middle values: %s", err)
	}

	return fmt.Sprintf("middle sum: %d", sum), nil
}

func parseRules(input []string) map[string]map[string]bool {
	rules := make(map[string]map[string]bool)
	for line := range slices.Values(input) {
		fields := strings.Split(line, "|")
		if rules[fields[0]] == nil {
			rules[fields[0]] = make(map[string]bool)
		}

		rules[fields[0]][fields[1]] = true
	}

	return rules
}

func parseUpdates(input []string) [][]string {
	var res [][]string
	for line := range slices.Values(input) {
		res = append(res, strings.Split(line, ","))
	}
	return res
}

func isValidUpdate(update []string, rules map[string]map[string]bool) bool {
	for i := 1; i < len(update); i++ {
		ruleSet := rules[update[i]]
		for j := i - 1; j >= 0; j-- {
			if ruleSet[update[j]] {
				return false
			}
		}
	}

	return true
}

func fixUpdate(update []string, rules map[string]map[string]bool) []string {
	for i := 1; i < len(update); i++ {
		ruleSet := rules[update[i]]
		lastPlace := i
		for j := i - 1; j >= 0; j-- {
			if ruleSet[update[j]] {
				update[j], update[lastPlace] = update[lastPlace], update[j]
				lastPlace = j
			}
		}
	}

	return update
}

func addUpdateMiddles(updates [][]string) (int, error) {
	sum := 0
	for update := range slices.Values(updates) {
		mid := update[(len(update)-1)/2]
		val, err := strconv.Atoi(mid)
		if err != nil {
			return 0, fmt.Errorf("couldn't parse mid value: %s", err)
		}
		sum += val
	}
	return sum, nil
}
