package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMuls(t *testing.T) {
	tests := []struct {
		msg   string
		input string
		want  int
	}{
		{
			msg:   "simple",
			input: "mul(1,2)",
			want:  2,
		},
		{
			msg:   "not enabled",
			input: "don't()mul(1,2)",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			got, err := doMuls(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
