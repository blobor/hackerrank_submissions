package strings

// Strings: Alternating Characters
// More info: https://www.hackerrank.com/challenges/alternating-characters

// AlternatingCharacters must return an integer representing the minimum number of deletions to make the alternating string.
func AlternatingCharacters(s string) int32 {
	result := int32(0)
	runes := []rune(s)

	prev := runes[0]
	for index := 1; index < len(runes); index++ {
		current := runes[index]
		if prev == current {
			result++
		}

		prev = current
	}

	return result
}
