package kata

func RepeatStr(repetitions int, value string) string {
	str := ""
	for i := 1; i <= repetitions; i++ {
		str += value
	}

	return str
}
