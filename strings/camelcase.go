package strings

import (
	str "strings"
	"unicode"
)

// Strings: CamelCase
// More info: https://www.hackerrank.com/challenges/camelcase

// CamelCase returns the integer number of words in the input string.
func CamelCase(s string) int32 {
	s = str.TrimSpace(s)
	c := len(s)
	r := int32(0)

	if c < 0 {
		return r
	}

	if c > 0 {
		r++
	}

	for _, ch := range s {
		if unicode.IsUpper(ch) {
			r++
		}
	}

	return r
}
