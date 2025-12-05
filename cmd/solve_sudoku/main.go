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
	_ = puzzleStr
	fmt.Println("invalid puzzle")
	os.Exit(1)
}
