package kata

func combat(health, damage float64) float64 {
	if damage < health {
		return health - damage
	}
	return 0.0
}
