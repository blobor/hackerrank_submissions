package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTheLoveLetterMystery(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected int32
	}{
		{"Test case 00", "abc", 2},
		{"Test case 01", "abcba", 0},
		{"Test case 02", "abcd", 4},
		{"Test case 03", "cba", 2},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := TheLoveLetterMystery(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
