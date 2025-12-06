package sudoku

import "errors"

type BoardState string

var (
	ErrTooShort      = errors.New("too short")
	ErrTooLong       = errors.New("too long")
	ErrInvalid       = errors.New("invalid puzzle")
	ErrAlreadySolved = errors.New("already solved")
)

func New(puzzle string) (BoardState, error) {
	if len(puzzle) < 81 {
		return "", ErrTooShort
	}
	if len(puzzle) > 81 {
		return "", ErrTooLong
	}
	if !IsValidSudoku(puzzle) {
		return "", ErrInvalid
	}
	state := Check(puzzle)
	if !state.IsValid() {
		return "", ErrInvalid
	}
	if state.IsSolved() {
		return "", ErrAlreadySolved
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

type SudokuState struct {
	// A sudoku is valid if it is a string of exactly 81 characters that are either filled with valid digits '1' through
	// '9' or are empty ('0').
	isValid bool
	// A sudoku is complete if it is a string of exactly 81 characters with no empty digits ('0').
	isComplete bool
	// A sudoku is correct if it has no duplicated or invalid digits.
	isCorrect bool

	// A sudoku is solved if it is valid, complete, and correct.
}

func (s SudokuState) IsValid() bool {
	return s.isValid
}

func (s SudokuState) IsComplete() bool {
	return s.isComplete
}

func (s SudokuState) IsCorrect() bool {
	return s.isCorrect
}

func (s SudokuState) IsSolved() bool {
	return s.isValid && s.isComplete && s.isCorrect
}

func Check(puzzle string) SudokuState {
	if len(puzzle) != 81 {
		// wrong length -> can't be valid, complete, or solved
		return SudokuState{}
	}

	result := SudokuState{true, true, true}

	rows := [9][9]bool{}
	cols := [9][9]bool{}
	houses := [9][9]bool{}

	for i, c := range puzzle {
		row := cellRows[i]
		col := cellCols[i]
		house := cellHouses[i]

		if c == Empty {
			// ignore empties
			result.isComplete = false
			continue
		}

		if !(c >= '1' && c <= '9') {
			// invalid character
			result.isValid = false
			result.isCorrect = false
			continue
		}
		d := int(c) - '1'
		if rows[row][d] {
			result.isCorrect = false
		}
		rows[row][d] = true

		if cols[col][d] {
			result.isCorrect = false
		}
		cols[col][d] = true

		if houses[house][d] {
			result.isCorrect = false
		}
		houses[house][d] = true
	}

	return result
}
