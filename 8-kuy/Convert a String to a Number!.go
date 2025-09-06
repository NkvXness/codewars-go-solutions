package kata

func StringToNumber(str string) int {
	var number int

	negative := str[0] == '-'

	startIndex := 0

	if negative {
		startIndex = 1
	}

	for i := startIndex; i < len(str); i++ {
		digit := int(str[i] - '0')
		number = number*10 + digit
	}

	if negative {
		number = -number
	}

	return number
}
