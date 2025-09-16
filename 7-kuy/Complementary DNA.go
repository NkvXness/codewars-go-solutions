package kata

import "strings"

func DNAStrand(dna string) string {
	mapDNA := map[string]string{
		"A": "T",
		"T": "A",
		"G": "C",
		"C": "G",
	}

	var result strings.Builder

	for _, char := range dna {
		if value, exist := mapDNA[string(char)]; exist {
			result.WriteString(value)
		}
	}

	return result.String()
}
