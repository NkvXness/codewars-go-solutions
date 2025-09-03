package kata

func SetAlarm(employed, vacation bool) bool {
	if employed == true && vacation == false {
		return true
	}
	return false
}
