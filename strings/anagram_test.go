package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		s        string
		expected int32
	}{
		{"Test case 00", "aaabbb", int32(3)},
		{"Test case 01", "ab", int32(1)},
		{"Test case 02", "abc", int32(-1)},
		{"Test case 03", "mnop", int32(2)},
		{"Test case 04", "xyyx", int32(0)},
		{"Test case 05", "xaxbbbxx", int32(1)},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := Anagram(tc.s)
			assert.Equal(t, tc.expected, res)
		})
	}
}
