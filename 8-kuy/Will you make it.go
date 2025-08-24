package kata

func ZeroFuel(distanceToPump int, mpg int, fuelLeft int) bool {
	return fuelLeft*mpg >= distanceToPump
}
