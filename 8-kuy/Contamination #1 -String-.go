package kata

func Contamination(text, char string) string {
	str := ""

	for i := 1; i <= len(text); i++ {
		if char != "" || text != "" {
			str += char
		}
	}

	return str
}
