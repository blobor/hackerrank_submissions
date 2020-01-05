package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGemstones(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      []string
		expected int32
	}{
		{
			name:     "Test case 27",
			arr:      []string{"abcdde", "baccd", "eeabg"},
			expected: 2,
		},
		{
			name:     "Test case 28",
			arr:      []string{"basdfj", "asdlkjfdjsa", "bnafvfnsd", "oafhdlasd"},
			expected: 4,
		},
		{
			name:     "Test case 29",
			arr:      []string{"vtrjvgbj", "mkmjyaeav", "sibzdmsk"},
			expected: 0,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := Gemstones(tc.arr)
			assert.Equal(t, tc.expected, res)
		})
	}
}
