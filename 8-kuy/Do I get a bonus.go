package kata

import "strconv"

func BonusTime(salary int, bonus bool) string {
	str := "£"

	if bonus {
		salary = salary * 10
	}

	str += strconv.Itoa(salary)
	return str
}
