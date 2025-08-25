package kata

import "math"

func SquareOrSquareRoot(arr []int) []int {
	newArr := []int{}

	for _, value := range arr {
		sqrt := math.Sqrt(float64(value))
		if sqrt == float64(int(sqrt)) {
			newArr = append(newArr, int(sqrt))
		} else {
			newArr = append(newArr, value*value)
		}
	}

	return newArr
}
