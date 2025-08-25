package kata

func Well(x []string) string {
	count := 0

	for _, value := range x {
		if value == "good" {
			count++
		}
	}

	if count == 1 || count == 2 {
		return "Publish!"
	} else if count > 2 {
		return "I smell a series!"
	} else {
		return "Fail!"
	}
}
