package strings

import "testing"

import "github.com/stretchr/testify/assert"

func TestPangrams(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"Test case 00",
			"We promptly judged antique ivory buckles for the next prize",
			"pangram",
		},
		{
			"Test case 01",
			"We promptly judged antique ivory buckles for the prize",
			"not pangram",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := Pangrams(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
