package twelve

import "fmt"

func Verse(i int) string {
	days := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	verses := []string{" a Partridge in a Pear Tree.",
		" two Turtle Doves, and",
		" three French Hens,",
		" four Calling Birds,",
		" five Gold Rings,",
		" six Geese-a-Laying,",
		" seven Swans-a-Swimming,",
		" eight Maids-a-Milking,",
		" nine Ladies Dancing,",
		" ten Lords-a-Leaping,",
		" eleven Pipers Piping,",
		" twelve Drummers Drumming,",
	}
	s := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", days[i-1])
	v := ""
	for n := i - 1; n >= 0; n-- {
		v += verses[n]
	}
	return s + v
}

func Song() string {
	s := ""
	for i := 1; i <= 12; i++ {
		if i != 12 {
			s += Verse(i) + "\n"
		} else {
			s += Verse(i)
		}
	}
	return s
}
