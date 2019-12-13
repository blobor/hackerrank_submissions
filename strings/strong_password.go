package strings

import (
	str "strings"
	"unicode"
)

// Strings: Strong Password
// More info: https://www.hackerrank.com/challenges/strong-password

// Strong password criteria:
// Its length is at least 6.
// It contains at least one digit.
// It contains at least one lowercase English character.
// It contains at least one uppercase English character.
// It contains at least one special character. The special characters are: !@#$%^&*()-+

// numbers = "0123456789"
// lower_case = abcdefghijklmnopqrstuvwxyz
// upper_case = ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// special_characters = !@#$%^&*()-+

const minLen = 6

var specialCharacters = map[rune]struct{}{
	'!': struct{}{},
	'@': struct{}{},
	'#': struct{}{},
	'$': struct{}{},
	'%': struct{}{},
	'^': struct{}{},
	'&': struct{}{},
	'*': struct{}{},
	'(': struct{}{},
	')': struct{}{},
	'-': struct{}{},
	'+': struct{}{},
}

// MinimumNumber finds the minimum number of characters that must add to make password strong.
func MinimumNumber(n int32, password string) int32 {
	res := int32(minLen)
	if n == 0 {
		return res
	}

	password = str.TrimSpace(password)
	runes := []rune(password)
	n = int32(len(runes))
	rules := make(map[int]struct{})
	rulesCount := 4

	for _, r := range runes {
		if unicode.IsDigit(r) {
			rules[0] = struct{}{}
		}

		if 'a' <= r && r <= 'z' {
			rules[1] = struct{}{}
		}

		if 'A' <= r && r <= 'Z' {
			rules[2] = struct{}{}
		}

		if _, ok := specialCharacters[r]; ok {
			rules[3] = struct{}{}
		}
	}

	return max(res-n, int32(rulesCount-len(rules)))
}

func max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}
