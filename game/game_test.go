package game

import (
	"testing"

	"github.com/suleyj/connect-4-server/board"
)

func TestOutofBounds(t *testing.T) {
	gameBoard := board.GameBoard{}
	gameBoard.CreateBoard()

	tests := []struct {
		name  string
		index int
		size  int
		want  bool
	}{
		{"row in bounds", 0, gameBoard.Rows, false},
		{"row outbounds", 6, gameBoard.Rows, true},
		{"column inbounds", 6, gameBoard.Columns, false},
		{"column outbounds", 7, gameBoard.Columns, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := outOfBounds(tt.index, tt.size)
			if got != tt.want {
				t.Errorf("outofbounds(%d, %d) = %t; want %t", tt.index, tt.size, got, tt.want)
			}
		})
	}
}

func TestIsWin(t *testing.T) {
	gameBoard := board.GameBoard{}
	gameBoard.CreateBoard()

	playerVal := 'R'
	column := 1

	gameBoard.PlaceOnBoard(column, playerVal)
	gameBoard.PlaceOnBoard(column, playerVal)
	gameBoard.PlaceOnBoard(column, playerVal)
	row, col, err := gameBoard.PlaceOnBoard(column, playerVal)

	if err != nil {
		t.Fatal("PlaceOnBoard: %w", err)
	}
	if !IsWin(gameBoard, row, col, 'R') {
		t.Error("Is Win error")
	}
}

func TestIsWinFromMiddle(t *testing.T) {
	gameBoard := board.GameBoard{}

	gameBoard.Board = [][]rune{
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', 'Y', 'R', ' ', ' ', ' '},
		{' ', ' ', 'R', 'R', ' ', ' ', ' '},
		{' ', ' ', 'R', 'R', ' ', ' ', ' '},
		{'R', 'Y', 'Y', 'Y', ' ', ' ', ' '},
	}

	gameBoard.Rows = 6
	gameBoard.Columns = 7

	playerVal := 'R'
	column := 2

	row, col, err := gameBoard.PlaceOnBoard(column, playerVal)

	if err != nil {
		t.Fatal("PlaceOnBoard: %w", err)
	}
	if !IsWin(gameBoard, row, col, 'R') {
		t.Error("Is Win error")
	}

}
