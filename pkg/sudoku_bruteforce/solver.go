package sudokubruteforce

import (
	"fmt"
	"log/slog"
	"os"
)

var (
	LogLevel   = new(slog.LevelVar)
	LogHandler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: LogLevel})
)

func init() {
	// logLevel.Set(slog.LevelDebug)
	LogLevel.Set(slog.LevelWarn)
	slog.SetDefault(slog.New(LogHandler))
}

const EMPTY = '0'

func FromString(puzzle string) ([][]byte, error) {
	if len(puzzle) != 81 {
		return nil, fmt.Errorf("expected puzzle to be exactly 81 chars long, got one with length %d instead (puzzle=%q)", len(puzzle), puzzle)
	}

	board := make([][]byte, 9)
	for i := range 9 {
		board[i] = make([]byte, 9)
		for j := range 9 {
			board[i][j] = puzzle[i*9+j]
		}
	}

	return board, nil
}

func ToString(board [][]byte) string {
	s := ""
	for i, row := range board {
		if i > 0 {
			s += "\n"
		}
		for j, cell := range row {
			if j > 0 {
				s += " "
			}
			s += string(cell)
		}
	}
	return s
}

func ToFlatString(board [][]byte) string {
	s := ""
	for _, row := range board {
		s += string(row)
	}
	return s
}

var CellHouses = [9][9]int{
	{0, 0, 0, 1, 1, 1, 2, 2, 2},
	{0, 0, 0, 1, 1, 1, 2, 2, 2},
	{0, 0, 0, 1, 1, 1, 2, 2, 2},

	{3, 3, 3, 4, 4, 4, 5, 5, 5},
	{3, 3, 3, 4, 4, 4, 5, 5, 5},
	{3, 3, 3, 4, 4, 4, 5, 5, 5},

	{6, 6, 6, 7, 7, 7, 8, 8, 8},
	{6, 6, 6, 7, 7, 7, 8, 8, 8},
	{6, 6, 6, 7, 7, 7, 8, 8, 8},
}

const One = int('1')

func IsValidSudoku(board [][]byte) bool {
	// Allocating fixed-size arrays on the stack is more efficient than using slices.
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	houses := [9][9]bool{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			d := int(board[row][col]) - One
			if d < 0 {
				continue
			}
			if d > 8 {
				panic("invalid digit")
			}

			if rows[row][d] {
				return false
			}
			rows[row][d] = true

			if cols[col][d] {
				return false
			}
			cols[col][d] = true

			house := CellHouses[row][col]
			if houses[house][d] {
				return false
			}
			houses[house][d] = true
		}
	}
	return true
}

// SolveSudoku solves the Sudoku puzzle using a recursive brute-force algorithm. `board` must be a 9x9 grid with '0'
// representing empty cells. The solution is written back to the `board` slice. The function returns true if the puzzle
// was solved.
func SolveSudoku(board [][]byte, answers [][]byte, i int) bool {
	if i >= 81 {
		slog.Info("finished solving")
		return true
	}

	row := i / 9
	col := i % 9
	d := board[row][col]
	if d != EMPTY {
		slog.Info("cell already solved", "row", row, "col", col, "digit", rune(d))
		return SolveSudoku(board, answers, i+1)
	}
	slog.Info("solving cell", "row", row, "col", col)

	for d := byte('1'); d <= '9'; d++ {
		answers[row][col] = d
		if !IsValidSudoku(answers) {
			slog.Debug("rejecting digit", "cell", i, "digit", rune(d))
			continue
		}

		// recurse, incrementing i to move on to the next digit
		if SolveSudoku(board, answers, i+1) {
			slog.Debug("accepting digit", "cell", i, "digit", rune(d))
			return true
		}
	}

	// failed to solve -> backtrack
	slog.Debug("failed to solve", "cell", i)

	// We have to reset the cell value here or else the board will be left in an invalid state.
	answers[row][col] = EMPTY

	return false
}

func CopyBoard(dst [][]byte, src [][]byte) {
	for i := range 9 {
		copy(dst[i], src[i])
	}
}

// DoSolveSudoku solves the given sudoku and updates the given slice with the solution.
func DoSolveSudoku(board [][]byte) {
	answers := make([][]byte, 9)
	for i := range 9 {
		answers[i] = make([]byte, 9)
	}
	// Pre-populate the answers with the values from the board.
	CopyBoard(answers, board)

	SolveSudoku(board, answers, 0)
	CopyBoard(board, answers)
}
