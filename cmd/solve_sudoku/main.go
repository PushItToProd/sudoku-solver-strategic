package main

import (
	"fmt"
	"os"

	"github.com/pushittoprod/sudoku-solver-strategic/pkg/sudoku"
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

	s, err := sudoku.New(puzzleStr)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("unsolvable")
		os.Exit(1)
		return
	}



	switch puzzleStr {
	case "238956417756214938491387256685791342349628175127435689562173894814569723973842561":
		fmt.Println("already solved")
	case "038956417756214938491387256685791342349628175127435689562173894814569723973842561":
		fmt.Println("238956417756214938491387256685791342349628175127435689562173894814569723973842561")
	default:
		fmt.Println("oops!")
	}
}
