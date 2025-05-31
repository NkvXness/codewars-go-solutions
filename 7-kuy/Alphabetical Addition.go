package kata

func AddLetters(letters []rune) rune {

	if len(letters) == 0 {
		return 'z'
	}

	var result rune

	for _, char := range letters {
		result += char - 'a' + 1
	}

	for result > 26 {
		result -= 26
	}

	return result + 'a' - 1

}
