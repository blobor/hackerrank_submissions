package strings

// Strings: HackerRank in a String!
// More info: https://www.hackerrank.com/challenges/hackerrank-in-a-string

// HackerRankInString must return a string YES or NO depending if input string contains
// the word hackerrank if a subsequence of its characters spell the word hackerrank.
func HackerrankInString(s string) string {
	q := []rune("hackerrank")

	for _, ch := range s {
		if ch == q[0] {
			q = q[1:]
		}

		if len(q) == 0 {
			return "YES"
		}
	}
	return "NO"
}
