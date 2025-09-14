package kata

func Cats(start, finish int) int {
	distance := finish - start

	return distance/3 + distance%3
}
