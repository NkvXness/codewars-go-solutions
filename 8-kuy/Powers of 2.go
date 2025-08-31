package kata

import "math"

func PowersOfTwo(n int) []uint64 {
	var arr []uint64
	for i := 0; i <= n; i++ {
		arr = append(arr, uint64(math.Pow(2, float64(i))))
	}
	return arr
}
