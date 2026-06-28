package game

import (
	"testing"

	"github.com/suleyj/connect-4-server/board"
)

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
