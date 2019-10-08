package arrays

// Arrays: Left Rotation
// More info: https://www.hackerrank.com/challenges/ctci-array-left-rotation

// RotateLeft performs left rotations on the array and returns the updated array.
func RotateLeft(a []int32, d int32) []int32 {
	count := len(a)
	result := make([]int32, len(a))

	for i := 0; i < count; i++ {
		oldIndex := (int(d) + i) % count
		result[i] = a[oldIndex]
	}
	return result
}
