package kata

func multipleOfIndex(ints []int) []int {
	arr := []int{}
	for i, value := range ints {
		if i != 0 && value%i == 0 {
			arr = append(arr, value)
		}
	}
	return arr
}
