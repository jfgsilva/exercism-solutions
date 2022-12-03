package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

const outputFormat string = "%-10v | %-25v | %-6v\n"

func (e *Entry) format(currency string, locale string) {
	e.formatDate(locale)
	e.formatDescription()
	e.formatChange(currency, locale)
}

func (e *Entry) formatDate(locale string) error {
	entryDateLayout := "2006-01-02"
	enUsLayout := "01/02/2006"
	nlNlLayout := "02-01-2006"
	date, err := time.Parse(entryDateLayout, e.Date)
	if err != nil {
		strErr := fmt.Sprintf("date has wrong format %v", e.Date)
		return errors.New(strErr)
	}
	if locale == "en-US" {
		e.Date = date.Format(enUsLayout)
	} else {
		e.Date = date.Format(nlNlLayout)
	}
	return nil
}

func (e *Entry) formatDescription() {
	if len(e.Description) > 25 {
		e.Description = e.Description[:22] + "..."
	}
}

func (e *Entry) formatChange(currency string, locale string) {

}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var entriesCopy []Entry
	// creates a copy of entries
	entriesCopy = append(entriesCopy, entries...)
	// refactored error handling
	switch {
	case !(currency == "USD" || currency == "EUR"):
		return "", errors.New("wrong currency")
	case !(locale == "en-US" || locale == "nl-NL"):
		return "", errors.New("wrong locale")
	}
	//refactored entriesCopy sorting: ascending date asc, description asc, change asc
	sort.Slice(entriesCopy, func(i, j int) bool {
		switch {
		case entries[i].Date < entries[j].Date:
			return true
		case entries[i].Date == entries[j].Date:
			if entries[i].Description < entries[j].Description {
				return true
			} else {
				return entries[i].Change < entries[j].Change
			}
		default:
			return false
		}
	})
	// refactored ledger header

	var s, header string

	switch locale {
	case "nl-NL":
		header = fmt.Sprintf(outputFormat, "Datum", "Omschrijving", "Verandering")
	case "en-US":
		header = fmt.Sprintf(outputFormat, "Date", "Description", "Change")
	}
	// initialize header
	s = header
	// Parallelism, always a great idea
	co := make(chan struct {
		i int
		s string
		e error
	})
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			//fmt.Println("before:", entry)
			//entry.format(currency, locale)
			//fmt.Println("after:", entry)
			if len(entry.Date) != 10 {
				//fmt.Println("error 1")
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("error 1")}
			}
			d1, _, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
			// if d2 != '-' {
			// 	fmt.Println("error 2")
			// 	co <- struct {
			// 		i int
			// 		s string
			// 		e error
			// 	}{e: errors.New("error 2")}
			// }
			if d4 != '-' {
				//fmt.Println("error 3")
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("error 3")}
			}
			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			} else {
				de = de + strings.Repeat(" ", 25-len(de))
			}
			var d string
			if locale == "nl-NL" {
				d = d5 + "-" + d3 + "-" + d1
			} else if locale == "en-US" {
				d = d3 + "/" + d5 + "/" + d1
			}
			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			var a string
			if locale == "nl-NL" {
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				}
				// } else {
				// 	co <- struct {
				// 		i int
				// 		s string
				// 		e error
				// 	}{e: errors.New("error 4")}
				// }
				a += " "
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + "."
				}
				a = a[:len(a)-1]
				a += ","
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += "-"
				} else {
					a += " "
				}
			} else if locale == "en-US" {
				if negative {
					a += "("
				}
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				}
				// solved with error handling added at the start
				// } else {
				// 	co <- struct {
				// 		i int
				// 		s string
				// 		e error
				// 	}{e: errors.New("error 5")}
				// }
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + ","
				}
				a = a[:len(a)-1]
				a += "."
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += ")"
				} else {
					a += " "
				}
			}
			// } else {
			// 	co <- struct {
			// 		i int
			// 		s string
			// 		e error
			// 	}{e: errors.New("error 6")}
			// }
			var al int
			for range a {
				al++
			}
			co <- struct {
				i int
				s string
				e error
			}{i: i, s: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
				strings.Repeat(" ", 13-al) + a + "\n"}
		}(i, et)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}
	for i := 0; i < len(entriesCopy); i++ {
		s += ss[i]
	}
	return s, nil
}
