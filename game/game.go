package game

import (
	"slices"

	"github.com/suleyj/connect-4-server/board"
)

type point struct {
	x int
	y int
}

type adjacentDirections struct {
	first  string
	second string
}

func IsDraw(g board.GameBoard) bool {
	for _, row := range g.Board {
		if slices.Contains(row, ' ') {
				return false
		}	
	}

	return true
}

func IsWin(g board.GameBoard, row, column int, playerVal rune) bool {

	directionPoints := map[string]point{
		"down":              {0, -1},
		"up":                {0, 1},
		"left":              {-1, 0},
		"right":             {1, 0},
		"diagonalUpRight":   {1, 1},
		"diagonalDownLeft":  {-1, -1},
		"diagonalUpLeft":    {-1, 1},
		"diagonalDownRight": {1, -1},
	}

	pairedDirections := []adjacentDirections{
		{"down", "up"},
		{"left", "right"},
		{"diagonalUpRight", "diagonalDownLeft"},
		{"diagonalUpLeft", "diagonalDownRight"},
	}

	for _, pair := range pairedDirections {
		chipCount := 1
		chipCount += countChips(g, row, column, directionPoints[pair.first], playerVal)
		chipCount += countChips(g, row, column, directionPoints[pair.second], playerVal)

		if chipCount >= 4 {
			return true
		}
	}

	return false

}

func outOfBounds(index int, size int) bool {

	if index > size-1 || index < 0 {
		return true
	}

	return false
}

// counts the number of matching chips in a given direction
func countChips(g board.GameBoard, row int, column int, p point, playerVal rune) int {
	chipCount := 0
	for i := 1; ; i++ {
		columnDirection := column + (p.x * i)
		rowDirection := row + (p.y * i)

		if outOfBounds(columnDirection, g.Columns) {
			break
		}

		if outOfBounds(rowDirection, g.Rows) {
			break
		}

		if g.Board[rowDirection][columnDirection] != playerVal {
			break
		}

		if g.Board[rowDirection][columnDirection] == playerVal {
			chipCount += 1
		}
	}

	return chipCount
}
