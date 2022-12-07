package romannumerals

import (
	"errors"
	"fmt"
)

var (
	// map used to convert into roman numeral
	tabula = map[int]string{3000: "MMM",
		2000: "MM",
		1000: "M",
		900:  "CM",
		800:  "DCCC",
		700:  "DCC",
		600:  "DC",
		500:  "D",
		400:  "CD",
		300:  "CCC",
		200:  "CC",
		100:  "C",
		90:   "XC", 80: "LXXX", 70: "LXX", 60: "LX", 50: "L", 40: "XL", 30: "XXX", 20: "XX", 10: "X",
		9: "IX", 8: "VIII", 7: "VII", 6: "VI", 5: "V", 4: "IV", 3: "III", 2: "II", 1: "I"}
)

func ToRomanNumeral(input int) (string, error) {
	// error handling
	if input > 3999 || input <= 0 {
		return "", errors.New("not supported by roman numeral")
	}
	// vars to hold different parcels of input so it can be translated into corresponding oman numeral
	var thousands, hundreds, tens, units int
	// int division is used to remove the remainder and multiplied by same value so we can use tabula translation
	thousands = input / 1000 * 1000
	hundreds = (input - thousands) / 100 * 100
	tens = (input - thousands - hundreds) / 10 * 10
	units = (input - thousands - hundreds - tens) / 1 * 1
	roman := ""
	// fmt.Println(input, thousands, hundreds, tens, units)
	if thousands != 0 {
		roman += tabula[thousands]
	}

	if hundreds != 0 {
		roman += tabula[hundreds]
	}
	if tens != 0 {
		roman += tabula[tens]
	}
	if units != 0 {
		roman += tabula[units]
	}

	fmt.Println(input, thousands, hundreds, tens, units, "| roman:", roman)
	return roman, nil
}
