package matrix

import (
	"errors"
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
// at most cols is as big as matrix (1 col n rows)
func (m Matrix) Cols() [][]int {
	cols := [][]int{}
	for r := 0; r <= len(m); r++ {
		var row []int
		for c := 0; c <= len(m); c++ {
			// we change the indexes here to return by cols
			val, ok := m[point{x: r, y: c}]
			if ok {
				row = append(row, val)
			} else {
				break
			}
		}
		if len(row) != 0 {
			cols = append(cols, row)
		}
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	cols := [][]int{}
	for r := 0; r <= len(m); r++ {
		var row []int
		for c := 0; c <= len(m); c++ {
			// we keep the indexes to return by row
			val, ok := m[point{x: c, y: r}]
			if ok {
				row = append(row, val)
			} else {
				break
			}
		}
		if len(row) != 0 {
			cols = append(cols, row)
		}
	}
	return cols
}

func (m Matrix) Set(row, col, val int) bool {
	_, ok := m[point{x: col, y: row}]
	if ok {
		m[point{x: col, y: row}] = val
		return true
	}
	return false
}
