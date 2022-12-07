package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int, 26)
	for oldKey, slcLetters := range in {
		for _, letter := range slcLetters {
			out[strings.ToLower(letter)] = oldKey
		}
	}
	return out
}
