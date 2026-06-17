package main

// These imports will be used later in the tutorial. If you save the file
// now, Go might complain they are unused, but that's fine.
// You may also need to run `go mod tidy` to download bubbletea and its
// dependencies.
import (
	"fmt"
	"strings"
	// "os"
	//    tea "charm.land/bubbletea/v2"
)

type GameBoard struct {
	rows    int
	columns int
	board   [][]int
}

func (g *GameBoard) initBoard() {

	g.board = make([][]int, g.rows)

	for i := range g.rows {
		g.board[i] = make([]int, g.columns)
	}
}

type Increment struct {
	x, y int
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("| %d ", cell)
		}
		fmt.Println("|")
	}
	fmt.Println(strings.Repeat("----", 7))
}

func (g *GameBoard) placeOnBoard(row int, col int, val int) {
	g.board[row][col] = val
	if g.isWin(row, col, val) {
		println("Winner")
	} else {
		println("Try again")
	}
}

func (g GameBoard) isWin(row, column, val int) bool {
	increments := []Increment{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	for _, increment := range increments {
		win := true
		for i := 1; i < 4; i++ {
			if row+(increment.x*i) > g.rows-1 || row+(increment.x*i) < 0 {
				win = false
				break
			}

			if column+(increment.y*i) > g.columns-1 || column+(increment.y*i) < 0 {
				win = false
				break
			}

			if g.board[row+(increment.x*i)][column+(increment.y*i)] != val {
				win = false
				break
			}

		}

		if win {
			return true
		}

	}

	return false
}

func main() {
	gameBoard := GameBoard{rows: 6, columns: 7}

	gameBoard.initBoard()

	fmt.Println("Connect 4 Server")
	printBoard(gameBoard.board)
	gameBoard.placeOnBoard(5, 0, 1)
	gameBoard.placeOnBoard(5, 1, 1)
	gameBoard.placeOnBoard(5, 2, 1)
	gameBoard.placeOnBoard(5, 3, 1)
	println()
	printBoard(gameBoard.board)
}
