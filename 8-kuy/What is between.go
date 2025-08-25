package kata

func Between(a, b int) []int {
	arr := []int{}
	for ; a <= b; a++ {
		arr = append(arr, a)
	}
	return arr
}
