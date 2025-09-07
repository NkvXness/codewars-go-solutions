package kata

func monkeyCount(n int) []int {
	arr := []int{}

	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	return arr
}
