package kata

func FindMultiples(integer, limit int) []int {
	arr := []int{}

	for i := integer; integer <= limit; integer += i {
		arr = append(arr, integer)
	}

	return arr
}
