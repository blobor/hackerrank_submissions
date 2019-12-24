package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		s        string
		expected int32
	}{
		{"Test case 00", "aaab", int32(3)},
		{"Test case 01", "baa", int32(0)},
		{"Test case 02", "aaa", int32(-1)},
		{"Test case 03", "quyjjdcgsvvsgcdjjyq", int32(1)},
		{"Test case 04", "hgygsvlfwcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcflvsgygh", int32(8)},
		{"Test case 05", "fgnfnidynhxebxxxfmxixhsruldhsaobhlcggchboashdlurshxixmfxxxbexhnydinfngf", int32(33)},
		{"Test case 06", "bsyhvwfuesumsehmytqioswvpcbxyolapfywdxeacyuruybhbwxjmrrmjxwbhbyuruycaexdwyfpaloyxbcpwsoiqtymhesmuseufwvhysb", int32(23)},
		{"Test case 07", "fvyqxqxynewuebtcuqdwyetyqqisappmunmnldmkttkmdlnmnumppasiqyteywdquctbeuwenyxqxqyvf", int32(25)},
		{"Test case 08", "mmbiefhflbeckaecprwfgmqlydfroxrblulpasumubqhhbvlqpixvvxipqlvbhqbumusaplulbrxorfdylqmgfwrpceakceblfhfeibmm", int32(44)},
		{"Test case 09", "tpqknkmbgasitnwqrqasvolmevkasccsakvemlosaqrqwntisagbmknkqpt", int32(20)},
		{"Test case 10", "lhrxvssvxrhl", int32(-1)},
		{"Test case 11", "prcoitfiptvcxrvoalqmfpnqyhrubxspplrftomfehbbhefmotfrlppsxburhyqnpfmqlaorxcvtpiftiocrp", int32(14)},
		{"Test case 12", "kjowoemiduaaxasnqghxbxkiccikxbxhgqnsaxaaudimeowojk", int32(-1)},
		{"Test case 13", "hgygsvlfcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcwflvsgygh", int32(44)},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := PalindromeIndex(tc.s)
			assert.Equal(t, tc.expected, res)
		})
	}
}
