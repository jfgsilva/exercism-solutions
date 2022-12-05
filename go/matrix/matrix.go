package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix map[point]int

type point struct {
	x int
	y int
}

func New(s string) (Matrix, error) {
	rows := strings.Split(strings.Trim(s, " "), "\n")
	var m Matrix = Matrix{}
	fmt.Println("rows", rows)
	rowLength := 0
	for rowIndex, row := range rows {
		row = strings.TrimSpace(row)
		cols := strings.Split(row, " ")
		// cols is actually the row, and we check here for uneveness of length
		if rowLength == 0 {
			rowLength = len(cols)
		}
		if rowLength != len(cols) {
			return m, errors.New("uneven rows")
		}
		for colIndex, col := range cols {
			col = strings.TrimSpace(col)
			val, err := strconv.Atoi(col)
			if err != nil {
				return Matrix{}, err
			}
			p := point{x: colIndex,
				y: rowIndex,
			}
			m[p] = val
		}
	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	panic("not implemented")
}

func (m Matrix) Rows() [][]int {
	panic("not implemented")
}

func (m Matrix) Set(row, col, val int) bool {
	panic("not implemented")
}
