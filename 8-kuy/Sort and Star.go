package kata

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)

	firstWord := strings.Split(arr[0], "")

	return strings.Join(firstWord, "***")
}
