package kata

func CountPositivesSumNegatives(numbers []int) []int {
	arr := []int{}
	count := 0
	sum := 0

	for _, value := range numbers {
		if value > 0 {
			count++
		}

		if value < 0 {
			sum += value
		}
	}

	arr = append(arr, count, sum)

	return arr
}
