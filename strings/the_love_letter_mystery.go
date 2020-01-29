package strings

import "math"

// Strings: The Love-Letter Mystery
// More info: https://www.hackerrank.com/challenges/the-love-letter-mystery

// TheLoveLetterMystery returns the integer representing the minimum number of operations needed to make the string a palindrome.
func TheLoveLetterMystery(s string) int32 {
	var result int32
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		r1, r2 := runes[i], runes[j]
		result += int32(math.Abs(float64(r1 - r2)))
	}

	return result
}
