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
			msg:    "safe if remove 2",
			input:  "1 2 2 3 4",
			isSafe: true,
		},
		{
			msg:    "unsafe decrease too fast",
			input:  "9 7 6 2 1",
			isSafe: false,
		},
		{
			msg:    "safe sorted if remove 3",
			input:  "1 3 2 4 5",
			isSafe: true,
		},
		{
			msg:    "safe descends properly if 4 removed",
			input:  "8 6 4 4 1",
			isSafe: true,
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
			msg:    "safe if remove last",
			input:  "72 75 78 79 79",
			isSafe: true,
		},
		{
			msg:    "safe if remove high last",
			input:  "86 87 88 91 96",
			isSafe: true,
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
