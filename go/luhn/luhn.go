package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid determine whether or not the given number (as string parameter) is valid per the Luhn formula
// return true for valid number, false otherwise
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	length := len(input)
	if length <= 1 {
		return false
	}
	sum := 0
	var digit int
	for i, r := range input {
		if !unicode.IsDigit(r) {
			return false
		}
		digit, _ = strconv.Atoi(string(r))
		if i%2 == length%2 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
