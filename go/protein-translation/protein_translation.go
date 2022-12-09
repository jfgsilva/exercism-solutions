package protein

import (
	"errors"
	"fmt"
)

var ErrStop = errors.New("a stop codon has been reached")
var ErrInvalidBase = errors.New("an invalid base has been fed")

func FromRNA(rna string) ([]string, error) {
	fmt.Println(rna)
	return []string{}, ErrStop
}

func FromCodon(codon string) (string, error) {
	fmt.Println(codon)
	return "", ErrInvalidBase
}
