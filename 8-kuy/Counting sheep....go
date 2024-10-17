package kata

func CountSheeps(numbers []bool) int {
	var count int
	for _, value := range numbers {
		if value == true {
			count++
		}
	}
	return count
}
