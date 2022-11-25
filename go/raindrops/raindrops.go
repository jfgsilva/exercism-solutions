package raindrops

import "strconv"

func Convert(number int) string {
	rd := ""
	if number%3 == 0 {
		rd = "Pling"
	}
	if number%5 == 0 {
		rd = rd + "Plang"
	}
	if number%7 == 0 {
		rd = rd + "Plong"
	}
	if rd == "" {
		return strconv.Itoa(number)
	}
	return rd
}
