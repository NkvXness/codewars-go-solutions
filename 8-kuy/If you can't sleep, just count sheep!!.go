package kata

import "strconv"

func countSheep(num int) string {
	result := ""
	for i := 1; i <= num; i++ {
		result += strconv.Itoa(i) + " sheep..."
	}
	return result

}
