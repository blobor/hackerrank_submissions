package strings

import "testing"

func TestAlternatingCharacters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		s        string
		expected int32
	}{
		{"Test case 00", "AAAA", int32(3)},
		{"Test case 01", "BBBBB", int32(4)},
		{"Test case 02", "ABABABAB", int32(0)},
		{"Test case 03", "BABABA", int32(0)},
		{"Test case 04", "AAABBB", int32(4)},
		{"Test case 05", "AAABBBAABB", int32(6)},
		{"Test case 06", "AABBAABB", int32(4)},
		{"Test case 07", "ABABABAA", int32(1)},
		{"Test case 08", "ABBABBAA", int32(3)},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := AlternatingCharacters(tc.s)

			if res != tc.expected {
				msg := "Expected result: %v, but returned: %v"
				t.Errorf(msg, tc.expected, res)
			}
		})
	}
}
