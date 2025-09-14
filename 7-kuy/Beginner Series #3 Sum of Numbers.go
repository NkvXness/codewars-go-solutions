package kata

func GetSum(a, b int) int {

	if a > b {
		a, b = b, a
	}

	n := b - a + 1
	s := n * (a + b) / 2

	return s
}
