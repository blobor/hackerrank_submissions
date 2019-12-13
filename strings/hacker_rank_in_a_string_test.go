package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHackerrankInString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name: "Test case 00",
			s: "hereiamstackerrank",
			expected: "YES",
		},
		{
			name: "Test case 01",
			s: "hackerworld",
			expected: "NO",
		},
		{
			name: "Test case 02",
			s: "hhaacckkekraraannk",
			expected: "YES",
		},
		{
			name: "Test case 03",
			s: "rhbaasdndfsdskgbfefdbrsdfhuyatrjtcrtyytktjjt",
			expected: "NO",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := HackerrankInString(tc.s)
			assert.Equal(t, tc.expected, res)
		})
	}
}
