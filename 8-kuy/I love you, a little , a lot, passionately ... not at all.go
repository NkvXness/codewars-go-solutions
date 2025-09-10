package kata

func HowMuchILoveYou(n int) string {
	arr := []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}

	return arr[(n-1)%len(arr)]
}
