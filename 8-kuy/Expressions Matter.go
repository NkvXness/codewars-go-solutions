package kata

func ExpressionMatter(a int, b int, c int) int {
	res1 := a * (b + c)
	res2 := a * b * c
	res3 := a + b + c
	res4 := (a + b) * c

	maxResult := res1

	if res2 > maxResult {
		maxResult = res2
	}
	if res3 > maxResult {
		maxResult = res3
	}
	if res4 > maxResult {
		maxResult = res4
	}

	return maxResult
}
