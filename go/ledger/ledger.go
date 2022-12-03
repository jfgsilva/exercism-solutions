package ledger

import (
	"errors"
	"fmt"
	"sort"
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

func (e Entry) String() string {
	return fmt.Sprintf(outputFormat, e.Date, e.Description, e.Change)
}

func (e *Entry) formatDate(locale string) {
	entryDateLayout := "2006-01-02"
	enUsLayout := "01/02/2006"
	nlNlLayout := "02-01-2006"
	date, err := time.Parse(entryDateLayout, e.Date)
	if err != nil {
		strErr := fmt.Sprintf("date has wrong format %v", e.Date)
		panic(strErr)
	}
	if locale == "en-US" {
		e.Date = date.Format(enUsLayout)
	} else {
		e.Date = date.Format(nlNlLayout)
	}
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
	// fill ledger
	for i := 0; i < len(entriesCopy); i++ {
		entriesCopy[i].format(currency, locale)
		s += fmt.Sprint(entriesCopy[i])
	}
	return s, nil
}
