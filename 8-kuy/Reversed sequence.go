package kata

func ReverseSeq(n int) []int {
	arr := []int{}

	for ; n > 0; n-- {
		arr = append(arr, n)
	}

	return arr
}
