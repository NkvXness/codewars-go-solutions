package kata

import (
	"strconv"
)

func ToCsvText(array [][]int) string {
	var csvString string

	for i, row := range array {
		var rowString string
		for j, value := range row {
			rowString += strconv.Itoa(value)
			if j < len(row)-1 {
				rowString += ","
			}
		}
		csvString += rowString
		if i < len(array)-1 {
			csvString += "\n"
		}
	}

	return csvString
}
