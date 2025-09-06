package kata

func Digitize(n int) []int {
	arr := []int{}

	if n == 0 {
		return []int{n}
	}

	for n > 0 {
		arr = append(arr, n%10)
		n /= 10
	}

	return arr
}
