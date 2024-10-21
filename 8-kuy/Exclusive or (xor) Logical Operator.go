package kata

func Xor(a bool, b bool) bool {
	if a != b || a == false && b == false {
		return a || b
	}
	return false
}
