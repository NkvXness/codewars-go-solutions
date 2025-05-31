package kata

import "math"

func MapVector(vector, circle1, circle2 []float64) []float64 {

	dx1 := vector[0] - circle1[0]
	dy1 := vector[1] - circle1[1]
	r1 := circle1[2]
	r2 := circle2[2]

	scale := r2 / r1
	dx2 := dx1 * scale
	dy2 := dy1 * scale

	x2 := circle2[0] + dx2
	y2 := circle2[1] + dy2

	x2 = math.Round(x2*100) / 100
	y2 = math.Round(y2*100) / 100

	return []float64{x2, y2}

}
