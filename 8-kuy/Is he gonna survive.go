package kata

func Hero(bullets, dragons int) bool {
	return bullets/dragons == 2 || bullets > dragons*2
}
