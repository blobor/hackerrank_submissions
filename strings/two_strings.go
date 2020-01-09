package strings

// Strings: Two Strings
// More info: https://www.hackerrank.com/challenges/two-strings

// TwoStrings returns a string, either "YES" or "NO" based on whether the strings share a common substring.
func TwoStrings(s1 string, s2 string) string {
	s1Map := strToRuneMap(s1)
	s2Map := strToRuneMap(s2)

	for r := range s1Map {
		if _, ok := s2Map[r]; ok {
			return "YES"
		}
	}

	return "NO"
}

func strToRuneMap(s string) map[rune]struct{} {
	m := make(map[rune]struct{})

	for _, r := range s {
		m[r] = struct{}{}
	}

	return m
}
