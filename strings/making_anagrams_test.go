package strings

import "testing"

func TestMakeAnagram(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a        string
		b        string
		expected int32
	}{
		{"Test case 00", "cde", "abc", int32(4)},
		{"Test case 01", "cdde", "aabc", int32(6)},
		{"Test case 02", "fcrxzwscanmligyxyvym", "jxwtrhvujlmrpdoqbisbwhmgpmeoke", int32(30)},
		{"Test case 03", "showman", "woman", int32(2)},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := MakeAnagram(tc.a, tc.b)

			if res != tc.expected {
				msg := "Expected result: %v, but returned: %v"
				t.Errorf(msg, tc.expected, res)
			}
		})
	}
}
