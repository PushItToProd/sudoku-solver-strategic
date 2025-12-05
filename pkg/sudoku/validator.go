package sudoku

import "errors"

type BoardState string

func New(puzzle string) (BoardState, error) {
	if len(puzzle) < 81 {
		return "", errors.New("too short")
	}
	if len(puzzle) > 81 {
		return "", errors.New("too long")
	}
	if !IsValidSudoku(puzzle) {
		return "", errors.New("invalid sudoku")
	}
	return BoardState(puzzle), nil
}

// cellHouses is a lookup table mapping cell indices to house numbers.
var cellHouses = [81]int{
	0, 0, 0, 1, 1, 1, 2, 2, 2,
	0, 0, 0, 1, 1, 1, 2, 2, 2,
	0, 0, 0, 1, 1, 1, 2, 2, 2,

	3, 3, 3, 4, 4, 4, 5, 5, 5,
	3, 3, 3, 4, 4, 4, 5, 5, 5,
	3, 3, 3, 4, 4, 4, 5, 5, 5,

	6, 6, 6, 7, 7, 7, 8, 8, 8,
	6, 6, 6, 7, 7, 7, 8, 8, 8,
	6, 6, 6, 7, 7, 7, 8, 8, 8,
}

var cellRows = [81]int{
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 4, 4, 4, 4, 4, 4, 4, 4,
	5, 5, 5, 5, 5, 5, 5, 5, 5,
	6, 6, 6, 6, 6, 6, 6, 6, 6,
	7, 7, 7, 7, 7, 7, 7, 7, 7,
	8, 8, 8, 8, 8, 8, 8, 8, 8,
}

var cellCols = [81]int{
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
}

const Empty = '0'

func IsValidSudoku(puzzle string) bool {
	if len(puzzle) != 81 {
		return false
	}

	rows := [9][9]bool{}
	cols := [9][9]bool{}
	houses := [9][9]bool{}

	for i, c := range puzzle {
		row := cellRows[i]
		col := cellCols[i]
		house := cellHouses[i]

		if c == Empty {
			// ignore empties
			continue
		}

		if !(c >= '1' && c <= '9') {
			// invalid character
			return false
		}
		d := int(c) - '1'
		if rows[row][d] {
			return false
		}
		rows[row][d] = true

		if cols[col][d] {
			return false
		}
		cols[col][d] = true

		if houses[house][d] {
			return false
		}
		houses[house][d] = true
	}

	return true
}
