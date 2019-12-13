package strings

// Strings: Mars Exploration
// More info: https://www.hackerrank.com/challenges/mars-exploration

// MarsExploration returns an integer representing the number of letters SOS changed during transmission
func MarsExploration(s string) int32 {
	res := int32(0)
	runes := []rune(s)

	for i := 0; i < len(runes); i += 3 {
		if 'S' != runes[i] {
			res++
		}

		if 'O' != runes[i+1] {
			res++
		}

		if 'S' != runes[i+2] {
			res++
		}
	}

	return res
}
