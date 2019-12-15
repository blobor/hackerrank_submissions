package strings

// Strings: Super Reduced String
// More info: https://www.hackerrank.com/challenges/reduced-string

// SuperReducedString returns the super reduced string or Empty String if the final string is empty.
func SuperReducedString(s string) string {
	res := []rune{}
	for _, ch := range s {
		resLen := len(res)
		if resLen == 0 {
			res = append(res, ch)
			continue
		}

		if res[resLen-1] == ch {
			res = res[:resLen-1]
			continue
		}

		res = append(res, ch)
	}

	if len(res) == 0 {
		return "Empty String"
	}

	return string(res)
}
