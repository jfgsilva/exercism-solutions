package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	strSlice := strings.Split(word, "")
	// isletter checks if char is letter or other char to be used as helper
	isletter := func(str string) bool {
		if str == "-" || str == " " {
			return false
		}
		return true
	}
	alpha := make(map[string]int, 26)
	for _, s := range strSlice {
		if isletter(s) {
			if val, ok := alpha[s]; ok {
				return false
			} else {
				alpha[s] = val
			}
		}
	}
	return true
}
