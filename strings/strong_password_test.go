package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		len      int32
		pass     string
		expected int32
	}{
		{"Test case 00", int32(3), "Ab1", int32(3)},
		{"Test case 01", int32(11), "#HackerRank", int32(1)},
		{"Test case 38", int32(5), "jnhqj", int32(3)},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := MinimumNumber(tc.len, tc.pass)
			assert.Equal(t, tc.expected, res)
		})
	}
}
