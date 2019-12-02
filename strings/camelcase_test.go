package strings

import "testing"

func TestCameCase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		s        string
		expected int32
	}{
		{"Test case 00", "saveChangesInTheEditor", int32(5)},
		{"Test case 01", "singleword", int32(1)},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := CamelCase(tc.s)
			if res != tc.expected {
				msg := "Expected result: %v, but returned: %v"
				t.Errorf(msg, tc.expected, res)
			}
		})
	}
}
