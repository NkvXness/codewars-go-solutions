package kata

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	result := dadYearsOld - sonYearsOld*2

	if result < 0 {
		return -result
	}

	return result
}
