package kata

func Solution(word string) string {
	newStr := ""
	for i := len(word) - 1; i >= 0; i-- {
		newStr += string(word[i])
	}
	return newStr
}
