package kata

func GetGrade(a, b, c int) rune {
	avrg := (a + b + c) / 3

	switch {
	case avrg <= 100 && avrg >= 90:
		return 'A'
	case avrg < 90 && avrg >= 80:
		return 'B'
	case avrg < 80 && avrg >= 70:
		return 'C'
	case avrg < 70 && avrg >= 60:
		return 'D'
	case avrg < 60 && avrg >= 0:
		return 'F'
	default:
		return '?'
	}

}
