package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver(t *testing.T) {
	tests := []struct {
		msg   string
		input []string
		want  int
	}{
		{
			msg: "example",
			input: []string{
				"190: 10 19",
				"3267: 81 40 27",
				"83: 17 5",
				"156: 15 6",
				"7290: 6 8 6 15",
				"161011: 16 10 13",
				"192: 17 8 14",
				"21037: 9 7 18 13",
				"292: 11 6 16 20",
			},
			want: 3749,
		},
		{
			msg: "single",
			input: []string{
				"3267: 81 40 27",
			},
			want: 3267,
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			solver := D7P1Solver{}
			got, err := solver.Solve(tt.input)
			require.NoError(t, err)
			assert.Equal(t, fmt.Sprintf("Total of valid equations: %d", tt.want), got)
		})
	}
}
