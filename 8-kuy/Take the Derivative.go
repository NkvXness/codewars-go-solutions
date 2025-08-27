package kata

import "fmt"

func Derive(coefficient, exponent int) string {
	result := coefficient * exponent
	newExp := exponent - 1

	return fmt.Sprintf("%dx^%d", result, newExp)
}
