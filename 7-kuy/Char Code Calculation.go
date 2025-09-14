package kata

import "strconv"

func Calc(s string) int {
	total := 0
	total2 := 0

	for _, char := range s {
		ascii := strconv.Itoa(int(char))

		for _, digit := range ascii {
			num, _ := strconv.Atoi(string(digit))
			total += num

			if digit == '7' {
				num = 1
			}
			total2 += num
		}
	}

	return total - total2
}
