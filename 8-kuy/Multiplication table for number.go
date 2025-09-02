package kata

import (
	"fmt"
	"strings"
)

func MultiTable(number int) string {
	str := ""

	for i := 1; i <= 10; i++ {
		result := i * number
		str += fmt.Sprintf("%d * %d = %d\n", i, number, result)
	}

	return strings.TrimRight(str, "\n")
}
