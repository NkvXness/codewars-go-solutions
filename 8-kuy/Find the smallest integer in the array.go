package kata

func SmallestIntegerFinder(numbers []int) int {
	min := 0

	for i := 0; i < len(numbers); i++ {
		if i == 0 || numbers[i] < min {
			min = numbers[i]
		}
	}

	return min
}
