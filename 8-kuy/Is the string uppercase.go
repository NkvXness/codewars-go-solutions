package kata

import "unicode"

type MyString string

func (s MyString) IsUpperCase() bool {
	for _, char := range s {
		if unicode.IsLower(char) {
			return false
		}
	}

	return true
}
