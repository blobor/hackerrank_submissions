package strings

// Strings: Palindrome Index
// More info: https://www.hackerrank.com/challenges/palindrome-index

// PalindromeIndex returns the index of the character to remove or -1.
func PalindromeIndex(s string) int32 {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		j := len(runes) - 1 - i

		if runes[i] == runes[j] {
			continue
		}

		if runes[j] == runes[i+1] && runes[j-1] == runes[i+2] {
			return int32(i)
		}

		if runes[i] == runes[j-1] && runes[i+1] == runes[j-2] {
			return int32(j)
		}
	}

	return -1
}
