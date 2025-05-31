package kata

func Game(frank, sam, tom [4]int) bool {
	return frank[3] > sam[1] && frank[3] > tom[1] && frank[2] > sam[0] && frank[2] > tom[0]
}
