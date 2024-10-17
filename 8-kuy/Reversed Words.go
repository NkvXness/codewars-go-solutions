package kata

import "strings"

func ReverseWords(str string) string {
	arr := strings.Fields(str)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	newStr := strings.Join(arr, " ")

	return newStr
}
