package protein

import (
	"errors"
)

var translater = map[string]string{"AUG": "Methionine", "UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine", "UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine", "UGU": "Cysteine", "UGC": "Cysteine", "UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP"}

var ErrStop = errors.New("a stop codon has been reached")
var ErrInvalidBase = errors.New("an invalid base has been fed")

func FromRNA(rna string) (protein []string, err error) {
	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		aa, err := FromCodon(codon)
		switch err {
		case ErrStop:
			return protein, nil
		case ErrInvalidBase:
			return protein, ErrInvalidBase
		}
		if err != nil {
			return protein, nil
		}
		protein = append(protein, aa)
	}
	return protein, nil
}

func FromCodon(codon string) (string, error) {

	val, ok := translater[codon]
	if ok {
		if val == "STOP" {
			return "", ErrStop
		}
		return val, nil
	} else {
		return "", ErrInvalidBase
	}

}
