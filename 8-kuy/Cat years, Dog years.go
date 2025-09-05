package kata

func CalculateYears(years int) (result [3]int) {
	for i := 0; i < len(result); i++ {
		if years == 1 {
			result[0] = 1
			result[1] = 15
			result[2] = 15
		} else if years == 2 {
			result[0] = 2
			result[1] = 24
			result[2] = 24
		} else {
			result[0] = years
			result[1] = 24 + 4*(years-2)
			result[2] = 24 + 4*(years-2) + (years - 2)
		}
	}
	return result
}
