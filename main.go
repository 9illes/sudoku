package main

import (
	"fmt"

	"github.com/9illes/sudoku/sudoku"
)

func main() {

	var board001 = sudoku.Board{
		{9, 8, 0, 0, 0, 0, 0, 0, 0},
		{2, 4, 1, 0, 0, 0, 0, 0, 0},
		{0, 7, 6, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 8, 9, 4, 0, 0, 0},
		{0, 0, 0, 0, 7, 6, 0, 0, 0},
		{0, 0, 0, 5, 3, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 3},
		{0, 0, 0, 0, 0, 0, 7, 2, 6},
		{0, 0, 0, 0, 0, 0, 5, 1, 9},
	}
	sudoku.Print(board001)
	if !sudoku.Solve(board001) {
		fmt.Println("Unsolvable grid !\n")
	}
	sudoku.Print(board001)
	fmt.Println()
}
