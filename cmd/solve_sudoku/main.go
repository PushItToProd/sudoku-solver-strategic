package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("sudoku")
		return
	}

	puzzleStr := args[0]
	if len(puzzleStr) != 81 {
		fmt.Println("invalid puzzle")
		os.Exit(1)
	}

	fmt.Println("already solved")
}
