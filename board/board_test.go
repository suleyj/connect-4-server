package board

import ( 
	"testing"
	"slices"
)

func TestCreateBoard(t *testing.T) {
	gameBoard := GameBoard{}
	gameBoard.CreateBoard()

	testBoard := [][]rune{
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},	
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},	
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},	
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},	
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},	
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},	
	}

	if !slices.EqualFunc(gameBoard.Board, testBoard, slices.Equal) {
		t.Error("Create board error")
	}
}

func TestPlaceOnBoard(t *testing.T){
	gameBoard := GameBoard{}
	gameBoard.CreateBoard()

	gameBoard.PlaceOnBoard(1, 'R')

	if gameBoard.Board[5][0] != 'R' {
		t.Error("Place on board error")
	}
}
