package strings

// Strings: Anagram
// More info: https://www.hackerrank.com/challenges/anagram

// Anagram returns the minimum number of characters to change to make the words anagrams, or -1 if it's not possible.
func Anagram(s string) int32 {
	runes := []rune(s)

	if len(runes)%2 != 0 {
		return -1
	}

	half := len(runes) / 2
	set := make(map[rune]int)
	for i, r := range runes {
		val, ok := set[r]
		if ok {
			if i < half {
				set[r] = val + 1
			} else {
				newVal := val - 1
				set[r] = newVal
				if newVal == 0 {
					delete(set, r)
				}
			}
		} else {
			set[r] = 1

			if i >= half {
				set[r] *= -1
			}
		}
	}

	if len(set) == 0 {
		return 0
	}

	p1 := 0
	p2 := 0
	for _, count := range set {
		if count > 0 {
			p1 += count
		} else {
			p2 += (count * -1)
		}
	}

	if p1 > p2 {
		return int32(p2)
	}
	return int32(p1)
}
