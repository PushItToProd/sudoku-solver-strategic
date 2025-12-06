package main

import (
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
	if len(puzzleStr) != 81 {
		fmt.Println("invalid puzzle")
		os.Exit(1)
	}

	_, err := sudoku.New(puzzleStr)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("unsolvable")
		os.Exit(1)
		return
	}

	state := sudoku.Check(puzzleStr)
	if state.IsSolved() {
		fmt.Println("already solved")
		return
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
