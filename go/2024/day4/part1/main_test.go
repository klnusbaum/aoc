package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver(t *testing.T) {
	tests := []struct {
		msg   string
		input []string
		want  string
	}{
		{
			msg: "one line",
			input: []string{
				"MMMSXXMASM",
			},
			want: "Num xmas: 1",
		},
		{
			msg: "two line line",
			input: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
			},
			want: "Num xmas: 2",
		},
		{
			msg: "simple",
			input: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			want: "Num xmas: 18",
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			solver := D4P1Solver{}
			got, err := solver.Solve(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
