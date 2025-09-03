package kata

import "sort"

func MergeArrays(a, b []int) []int {

	if len(a) == 0 && len(b) == 0 {
		return []int{}
	}

	merged := make([]int, 0, len(a)+len(b))
	merged = append(merged, a...)
	merged = append(merged, b...)

	sort.Ints(merged)

	uniq := merged[:0]
	for i, v := range merged {
		if i == 0 || v != merged[i-1] {
			uniq = append(uniq, v)
		}
	}

	return uniq
}
