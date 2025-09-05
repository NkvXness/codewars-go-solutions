package kata

func CloseCompare(a, b float64, margin ...float64) int {
	m := 0.0
	if len(margin) > 0 {
		m = margin[0]
	}

	diff := a - b

	if abs(diff) <= m {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
