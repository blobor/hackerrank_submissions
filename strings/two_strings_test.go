package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		s1       string
		s2       string
		expected string
	}{
		{"Test case 00_1", "hello", "world", "YES"},
		{"Test case 00_1", "hi", "world", "NO"},
		{"Test case 06_1", "wouldyoulikefries", "abcabcabcabcabcabc", "NO"},
		{"Test case 06_2", "hackerrankcommunity", "cdecdecdecde", "YES"},
		{"Test case 06_3", "jackandjill", "wentupthehill", "YES"},
		{"Test case 06_4", "writetoyourparents", "fghmqzldbc", "NO"},
		{"Test case 07_1", "aardvark", "apple", "YES"},
		{"Test case 07_2", "beetroot", "sandals", "NO"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := TwoStrings(tc.s1, tc.s2)
			assert.Equal(t, tc.expected, res)
		})
	}
}
