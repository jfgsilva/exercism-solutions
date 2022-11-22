package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	counts := 0
	for _, val := range cb[file] {
		if val {
			counts += 1
		}
	}
	return counts
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	counts := 0
	if rank < 1 || rank > 8 {
		return 0
	}
	for _, val := range cb {
		if val[rank-1] {
			counts += 1
		}
	}
	return counts
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	squares := 0
	for _, val := range cb {
		squares += len(val)
	}
	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupied := 0
	for _, val := range cb {
		for _, squares := range val {
			if squares {
				occupied += 1
			}
		}
	}
	return occupied
}
