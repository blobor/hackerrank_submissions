package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateLeft(t *testing.T) {
	testCases := []struct {
		name     string
		a        []int32
		d        int32
		expected []int32
	}{
		{
			"Test case 00",
			[]int32{1, 2, 3, 4, 5},
			int32(4),
			[]int32{5, 1, 2, 3, 4},
		},
		{
			"Test case 01",
			[]int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51},
			int32(10),
			[]int32{77, 97, 58, 1, 86, 58, 26, 10, 86, 51, 41, 73, 89, 7, 10, 1, 59, 58, 84, 77},
		},
		{
			"Test case 02",
			[]int32{33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60, 87, 97},
			int32(13),
			[]int32{87, 97, 33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := RotateLeft(tc.a, tc.d)
			assert.Equal(t, tc.expected, res)
		})
	}
}
