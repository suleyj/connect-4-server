package main

import (
	"fmt"

	"github.com/suleyj/connect-4-server/board"
	"github.com/suleyj/connect-4-server/display"
	"github.com/suleyj/connect-4-server/game"
	//    tea "charm.land/bubbletea/v2"
)

func main() {
	gameBoard := board.GameBoard{}
	gameBoard.CreateBoard()

	playerChoice := 0

	playerOneTurn := true

	win := false

	fmt.Println("Connect 4 Server")
	fmt.Println()

	for win == false {

		fmt.Print("Drop a piece into a column 1 - 7: ")
		fmt.Scan(&playerChoice)
		fmt.Println()

		playerVal := 'R'

		if !playerOneTurn {
			playerVal = 'Y'
		}

		row, col, err := gameBoard.PlaceOnBoard(playerChoice, playerVal)
		if err != nil {
			fmt.Println(err)
			continue
		}

		display.DisplayBoard(gameBoard.Board)

		if game.IsWin(gameBoard, row, col, playerVal) {
			win = true
		} else {
			playerOneTurn = !playerOneTurn
		}

	}

	fmt.Println()

	if playerOneTurn {
		fmt.Println("PLayer 1 is the Winner")
	} else {
		fmt.Println("PLayer 2 is the Winner")
	}

}
