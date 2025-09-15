package kata

func SizeToNumber(s string) (int, bool) {
	if s == "" {
		return 0, false
	} else if s == "m" {
		return 38, true
	}

	sizeMap := map[byte]int{'s': 36, 'l': 40}
	size, ok := sizeMap[s[len(s)-1]]
	if !ok {
		return 0, false
	}

	for _, r := range s[:len(s)-1] {
		if r != 'x' {
			return 0, false
		}
		size += sizeMap[s[len(s)-1]] - 38
	}

	return size, true
}
