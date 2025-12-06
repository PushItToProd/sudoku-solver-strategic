package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/pushittoprod/sudoku-solver-strategic/pkg/sudoku"
	sudokubruteforce "github.com/pushittoprod/sudoku-solver-strategic/pkg/sudoku_bruteforce"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("sudoku")
		os.Exit(1)
		return
	}

	puzzleStr := args[0]
	solveSudoku(puzzleStr)
}

func solveSudoku(puzzleStr string) {
	_, err := sudoku.New(puzzleStr)
	if errors.Is(err, sudoku.ErrAlreadySolved) {
		// Puzzle is already solved, just print it as is
		fmt.Println("already solved")
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Simple bruteforce solver for now
	board, err := sudokubruteforce.FromString(puzzleStr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	sudokubruteforce.DoSolveSudoku(board)
	fmt.Println(sudokubruteforce.ToFlatString(board))
}
