package board

import (
	"errors"
)

type GameBoard struct {
	Rows    int
	Columns int
	Board   [][]rune
}

func (g *GameBoard) CreateBoard() {

	g.Rows = 6
	g.Columns = 7
	g.Board = make([][]rune, g.Rows)

	for i := range g.Rows {
		g.Board[i] = make([]rune, g.Columns)

		for k := range g.Board[i] {
			g.Board[i][k] = ' '
		}
	}
}

func (g *GameBoard) PlaceOnBoard(col int, val rune) (int, int, error) {
	col = col - 1
	for i := g.Rows - 1; i >= 0; i-- {
		if g.Board[i][col] == ' ' {
			g.Board[i][col] = val
			return i, col, nil
		}
	}

	return 0, 0, errors.New("This column is full try another")
}
