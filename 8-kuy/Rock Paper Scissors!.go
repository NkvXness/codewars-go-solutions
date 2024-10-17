package kata

func Rps(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	} else if len(p1) == 8 && len(p2) == 5 {
		return "Player 1 won!"
	} else if len(p1) == 4 && len(p2) == 8 {
		return "Player 1 won!"
	} else if len(p1) == 5 && len(p2) == 4 {
		return "Player 1 won!"
	} else {
		return "Player 2 won!"
	}
}
