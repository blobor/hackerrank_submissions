package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarsExploration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected int32
	}{
		{"Test case 00", "SOSSPSSQSSOR", int32(3)},
		{"Test case 01", "SOSSOT", int32(1)},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := MarsExploration(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
