package strings

// Strings: Gemstones
// More info: https://www.hackerrank.com/challenges/gem-stones

// Gemstones an integer representing the number of gemstones found in the list of rocks.
func Gemstones(arr []string) int32 {
	set := make(map[rune]int)

	for _, s := range arr {
		runeSet := make(map[rune]struct{})
		for _, r := range s {
			runeSet[r] = struct{}{}
		}

		for r := range runeSet {
			set[r]++
		}
	}

	for r, i := range set {
		if i != len(arr) {
			delete(set, r)
		}
	}

	return int32(len(set))
}
