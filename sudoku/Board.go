package sudoku

import (
	"fmt"
	"time"
)

type Board [][]int

func Print(board Board) {
	fmt.Print("\033[H\033[2J")
	for _, row := range board {
		fmt.Println(row)
	}

	time.Sleep(50 * time.Millisecond)
}
