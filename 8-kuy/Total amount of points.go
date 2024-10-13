package kata

func Points(games []string) int {
	points := 0

	for _, value := range games {
		if value[0] > value[2] {
			points += 3
		} else if value[0] == value[2] {
			points += 1
		}
	}

	return points
}
