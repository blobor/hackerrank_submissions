package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperReducedString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test case 00", "aaabccddd", "abd"},
		{"Test case 01", "aa", "Empty String"},
		{"Test case 02", "baab", "Empty String"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := SuperReducedString(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
