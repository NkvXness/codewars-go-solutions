package kata

import "strings"

func IsPalindrome(str string) bool {
	count := 0
	str2 := strings.ToLower(str)

	for i := 0; i <= len(str2)-1; i++ {
		if str2[i] == str2[len(str2)-1-i] {
			count += 1
		}
	}

	if count == len(str2) {
		return true
	} else {
		return false
	}
}
