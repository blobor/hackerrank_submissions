package strings

// Strings: Pangrams
// More info: https://www.hackerrank.com/challenges/pangrams

// Pangrams returns the string "pangram" if the input string is a pangram. Otherwise, it should return "not pangram".
func Pangrams(s string) string {
	set := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := set[r]; ok {
			continue
		}

		if 'a' <= r && r <= 'z' {
			set[r] = struct{}{}
		}

		if 'A' <= r && r <= 'Z' {
			r += 'a' - 'A'
			set[r] = struct{}{}
		}
	}

	if len(set) == 26 {
		return "pangram"
	}

	return "not pangram"
}
