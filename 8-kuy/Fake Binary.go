package kata

import (
	"strconv"
)

func FakeBin(x string) string {
	newStr := ""

	for _, value := range x {
		if v, _ := strconv.Atoi(string(value)); v < 5 {
			newStr += "0"
		} else {
			newStr += "1"
		}
	}

	return newStr
}
