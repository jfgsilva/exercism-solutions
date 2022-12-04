package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
	sChange     string
}

const headerOutputFormat string = "%-10v | %-25v | %-6v\n"
const entryOutputFormat string = "%-10v | %-25v | %13v\n"

func (e *Entry) format(currency string, locale string) error {
	errDate := e.formatDate(locale)
	if errDate != nil {
		return errDate
	}
	e.formatDescription()
	e.formatChange(currency, locale)
	return nil
}

func (e Entry) String() string {
	return fmt.Sprintf(entryOutputFormat, e.Date, e.Description, e.sChange)
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

func (e *Entry) formatChange(currency string, locale string) error {
	cur := ""
	switch currency {
	case "USD":
		cur = "$"
	case "EUR":
		cur = "€"
	default:
		return errors.New("wrong currency input")
	}
	strChange := ""
	if e.Change < 0 {
		strChange = strconv.Itoa(e.Change)[1:]
	} else {
		strChange = strconv.Itoa(e.Change)
	}
	l := len(strChange)
	cents, hundreds, thousands, centSep, hundredSep := "", "", "", "", ""
	switch {
	case l < 2:
		hundreds = "0"
		cents = "0" + strChange
	case l == 2:
		cents = strChange[l-2:]
		hundreds = "0"

	case l <= 5:
		cents = strChange[l-2:]
		hundreds = strChange[:l-2]
	default:
		cents = strChange[l-2:]
		hundreds = strChange[l-5 : l-2]
		thousands = strChange[:l-5]
	}
	switch locale {
	case "nl-NL":
		centSep = ","
		if l > 5 {
			hundredSep = "."
		}
		if e.Change < 0 {
			e.sChange = fmt.Sprintf("%v %v%v%v%v%v-", cur, thousands, hundredSep, hundreds, centSep, cents)
		} else {
			e.sChange = fmt.Sprintf("%v %v%v%v%v%v ", cur, thousands, hundredSep, hundreds, centSep, cents)
		}
	case "en-US":
		centSep = "."
		if l > 5 {
			hundredSep = ","
		}
		if e.Change < 0 {
			e.sChange = fmt.Sprintf("(%v%v%v%v%v%v)", cur, thousands, hundredSep, hundreds, centSep, cents)
		} else {
			e.sChange = fmt.Sprintf("%v%v%v%v%v%v ", cur, thousands, hundredSep, hundreds, centSep, cents)
		}
	}
	return nil
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
		header = fmt.Sprintf(headerOutputFormat, "Datum", "Omschrijving", "Verandering")
	case "en-US":
		header = fmt.Sprintf(headerOutputFormat, "Date", "Description", "Change")
	}
	// initialize header
	s = header
	// fill ledger
	for i := 0; i < len(entriesCopy); i++ {
		err := entriesCopy[i].format(currency, locale)
		if err != nil {
			return "", err
		}
		s += fmt.Sprint(entriesCopy[i])
	}
	return s, nil
}
