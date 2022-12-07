// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

const (
	body   string = "For want of a %v the %v was lost."
	ending string = "And all for the want of a %v."
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	l := len(rhyme)
	if l == 0 {
		return []string{}
	}
	var s = make([]string, l)
	for i := 1; i < l; i++ {
		s[i-1] = fmt.Sprintf(body, rhyme[i-1], rhyme[i])
	}
	s[l-1] = fmt.Sprintf(ending, rhyme[0])
	return s
}
