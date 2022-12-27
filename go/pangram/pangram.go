package pangram

import "strings"

// IsPangram checks if the given string is a pangram.
func IsPangram(s string) bool {
	runeCounts := countRunes(strings.ToLower(s))
	for asc := 'a'; asc <= 'z'; asc++ {
		if runeCounts[asc] == 0 {
			return false
		}
	}
	return true
}

// countRunes builds a map of runes and how often they appear in the given string.
func countRunes(s string) map[rune]int {
	m := make(map[rune]int)
	for _, runeVal := range s {
		m[runeVal] += 1
	}
	return m
}
