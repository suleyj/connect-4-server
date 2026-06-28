package display

import (
	"fmt"
	"strings"
)

func DisplayBoard(board [][]rune) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("| %c ", cell)
		}
		fmt.Println("|")
	}
	dashes := strings.Repeat("----", 7)
	fmt.Printf("%s-\n", dashes)
}
