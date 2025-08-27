package kata

import "strconv"

func SumMix(arr []any) int {
	sum := 0
	for _, value := range arr {
		switch v := value.(type) {
		case int:
			sum += v
		case string:
			num, err := strconv.Atoi(v)
			if err == nil {
				sum += num
			}
		}
	}
	return sum
}
