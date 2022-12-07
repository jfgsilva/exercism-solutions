package strand

const (
	A rune = 65
	C rune = 67
	G rune = 71
	T rune = 84
	U rune = 85
)

func ToRNA(dna string) string {
	var rna = make([]rune, len(dna))
	for i, letter := range dna {
		switch letter {
		case A:
			rna[i] = U
		case C:
			rna[i] = G
		case G:
			rna[i] = C
		case T:
			rna[i] = A
		}
	}
	return string(rna)
}
