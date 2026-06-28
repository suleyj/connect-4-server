package game

import (
	"github.com/suleyj/connect-4-server/board"
)

type increment struct {
	x, y int
}

func IsWin(g board.GameBoard, row, column int, val rune) bool {
	increments := []increment{
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
			if row+(increment.x*i) > g.Rows-1 || row+(increment.x*i) < 0 {
				win = false
				break
			}

			if column+(increment.y*i) > g.Columns-1 || column+(increment.y*i) < 0 {
				win = false
				break
			}

			if g.Board[row+(increment.x*i)][column+(increment.y*i)] != val {
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
