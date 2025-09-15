package kata

import "math"

func EquableTriangle(a, b, c int) bool {
	p := a + b + c

	halfP := p / 2
	s := halfP * (halfP - a) * (halfP - b) * (halfP - c)
	result := int(math.Sqrt(float64(s)))

	if p == result {
		return true
	}

	return false
}
