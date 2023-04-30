package sudoku

// Solve solve a sudoku grid with backtracking
func Solve(board Board) bool {
	Print(board)
	length := len(board)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if board[i][j] == 0 {
				for n := 1; n <= length; n++ {
					if isValidProposition(board, i, j, n) {
						board[i][j] = n
						if Solve(board) {
							return true
						} else {
							board[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

// isInRow check if proposed number is already set in the row
func isInRow(board Board, row int, n int) bool {
	for _, cur := range board[row] {
		if cur == n {
			return true
		}
	}
	return false
}

// isInColumn check if proposed number is already set in the column
func isInColumn(board Board, column int, n int) bool {
	for _, cur := range board {
		if cur[column] == n {
			return true
		}
	}
	return false
}

// isInBox check if proposed number is already set in the box
func isInBox(board Board, row int, column int, n int) bool {
	top := row - row%3
	left := column - column%3
	// fmt.Printf("\nisInbox %d,%d,%d\ntop : %d left : %d\n", row, column, n, top, left)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// fmt.Printf("%d,%d\n", top+i, left+j)
			if board[top+i][left+j] == n {
				return true
			}
		}
	}
	return false
}

// isValidProposition check if the number proposed fit
func isValidProposition(board Board, row int, column int, n int) bool {
	return !isInBox(board, row, column, n) &&
		!isInRow(board, row, n) &&
		!isInColumn(board, column, n)
}
