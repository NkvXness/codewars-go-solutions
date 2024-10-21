package kata

import (
	"math"
)

func NearestSq(n int) int {
	return int(math.Pow(math.Round(math.Sqrt(float64(n))), 2))
}
