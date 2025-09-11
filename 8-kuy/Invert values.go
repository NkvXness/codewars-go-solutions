package kata

func Invert(arr []int) []int {
	arr2 := make([]int, len(arr))

	for i, value := range arr {
		arr2[i] = value * -1
	}

	return arr2
}
