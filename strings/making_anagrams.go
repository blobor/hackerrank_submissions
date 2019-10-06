package strings

// Strings: Making Anagrams
// More info: https://www.hackerrank.com/challenges/ctci-making-anagrams

// MakeAnagram returns a single integer denoting the number of characters
// you must delete to make the two strings anagrams of each other.
func MakeAnagram(a string, b string) int32 {
	strMap := make(map[rune]int32)
	result := int32(0)

	for _, char := range a {
		val, ok := strMap[char]

		if ok {
			strMap[char] = val + 1
		} else {
			strMap[char] = 1
		}
	}

	for _, char := range b {
		val, ok := strMap[char]

		if ok {
			strMap[char] = val - 1
		} else {
			strMap[char] = -1
		}
	}

	for _, count := range strMap {
		result += fastAbs(count)
	}

	return result
}

func fastAbs(x int32) int32 {
	return (x ^ (x >> 31)) - (x >> 31)
}
