package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReports(t *testing.T) {
	tests := []struct {
		msg    string
		input  string
		isSafe bool
	}{
		{
			msg:    "safe descending",
			input:  "7 6 4 2 1",
			isSafe: true,
		},
		{
			msg:    "unsafe increase too fast",
			input:  "1 2 7 8 9",
			isSafe: false,
		},
		{
			msg:    "unsafe increase too slow",
			input:  "1 2 2 3 4",
			isSafe: false,
		},
		{
			msg:    "unsafe decrease too fast",
			input:  "9 7 6 2 1",
			isSafe: false,
		},
		{
			msg:    "unsafe not sorted",
			input:  "1 3 2 4 5",
			isSafe: false,
		},
		{
			msg:    "unsafe too slow of decrease",
			input:  "8 6 4 4 1",
			isSafe: false,
		},
		{
			msg:    "safe ascending",
			input:  "1 3 6 7 9",
			isSafe: true,
		},
		{
			msg:    "unsafe all the same",
			input:  "4 4 4 4 4",
			isSafe: false,
		},
		{
			msg:    "safe inc 1",
			input:  "1 2 3 4 5",
			isSafe: true,
		},
		{
			msg:    "unsafe last the same",
			input:  "72 75 78 79 79",
			isSafe: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			r, err := newReport(tt.input)
			require.NoError(t, err)

			assert.Equal(t, r.isSafe(), tt.isSafe)
		})
	}
}
