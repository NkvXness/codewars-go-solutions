package kata

func CountBy(x, n int) []int {
	arr := []int{}
	for i := x; i <= x*n; i += x {
		arr = append(arr, i)
	}
	return arr
}
