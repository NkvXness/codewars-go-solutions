package kata

func Maps(x []int) []int {
	arr := []int{}
	for _, value := range x {
		arr = append(arr, value*2)
	}
	return arr
}
