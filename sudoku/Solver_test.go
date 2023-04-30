package sudoku

import (
	"testing"
)

var board001 = [][]int{
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

func TestIsInRow(t *testing.T) {
	type args struct {
		board Board
		row   int
		n     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"in row 1",
			args{
				board: board001,
				row:   0,
				n:     8,
			},
			true,
		},
		{
			"not in row 1",
			args{
				board: board001,
				row:   0,
				n:     2,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInRow(tt.args.board, tt.args.row, tt.args.n); got != tt.want {
				t.Errorf("isInRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInColumn(t *testing.T) {
	type args struct {
		board  Board
		column int
		n      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"in column 1",
			args{
				board:  board001,
				column: 0,
				n:      2,
			},
			true,
		},
		{
			"not in column 1",
			args{
				board:  board001,
				column: 0,
				n:      3,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInColumn(tt.args.board, tt.args.column, tt.args.n); got != tt.want {
				t.Errorf("isInColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInBox(t *testing.T) {
	type args struct {
		board  Board
		row    int
		column int
		n      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"in box 1",
			args{
				board:  board001,
				row:    1,
				column: 1,
				n:      6,
			},
			true,
		},
		{
			"not in box 1",
			args{
				board:  board001,
				row:    1,
				column: 1,
				n:      3,
			},
			false,
		},
		{
			"in box 9",
			args{
				board:  board001,
				row:    8,
				column: 8,
				n:      8,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInBox(tt.args.board, tt.args.row, tt.args.column, tt.args.n); got != tt.want {
				t.Errorf("isInBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidProposition(t *testing.T) {
	type args struct {
		board  Board
		row    int
		column int
		n      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"in box 5",
			args{
				board:  board001,
				row:    4,
				column: 3,
				n:      2,
			},
			false,
		},
		{
			"in box 2",
			args{
				board:  board001,
				row:    1,
				column: 5,
				n:      1,
			},
			false,
		},
		{
			"in box 2",
			args{
				board:  board001,
				row:    5,
				column: 8,
				n:      6,
			},
			false,
		},
		{
			"not in box 9",
			args{
				board:  board001,
				row:    5,
				column: 8,
				n:      1,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidProposition(tt.args.board, tt.args.row, tt.args.column, tt.args.n); got != tt.want {
				t.Errorf("isValidProposition() = %v, want %v", got, tt.want)
			}
		})
	}
}
