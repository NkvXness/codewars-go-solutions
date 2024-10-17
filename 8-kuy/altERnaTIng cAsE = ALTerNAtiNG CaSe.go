package kata

import "strings"

func ToAlternatingCase(str string) string {
	newStr := ""

	for _, char := range str {
		if char >= 65 && char <= 90 {
			newStr += strings.ToLower(string(char))
		} else {
			newStr += strings.ToUpper(string(char))
		}
	}

	return newStr
}
