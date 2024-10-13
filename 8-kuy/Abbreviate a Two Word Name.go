package kata

import (
	"strings"
)

func AbbrevName(name string) string {
	arr := strings.Fields(name)
	return strings.ToUpper(string(arr[0][0]) + "." + string(arr[1][0]))
}
