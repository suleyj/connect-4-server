package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/suleyj/connect-4-server/board"
	"github.com/suleyj/connect-4-server/game"
)

type player struct {
	playerValue rune
	playerName  string
}

type model struct {
	cursor        int
	p1            player
	p2            player
	board         board.GameBoard
	isPlayer1Turn bool
	win           bool
	err           error
	winner        player
	draw          bool
}

func initialModel() model {
	gameBoard := board.GameBoard{}
	gameBoard.CreateBoard()

	return model{
		cursor:        0,
		p1:            player{'X', "Player 1"},
		p2:            player{'Y', "Player 2"},
		board:         gameBoard,
		isPlayer1Turn: true,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left", "h":
			if m.cursor > 0 {
				m.cursor--
			}
		case "right", "l":
			if m.cursor < m.gb.Columns-1 {
				m.cursor++
			}
		case "enter", "space":
			var playerVal rune
			if m.isPlayer1Turn {
				playerVal = m.p1.playerValue
			} else {
				playerVal = m.p2.playerValue
			}
			row, col, err := m.gb.PlaceOnBoard(m.cursor+1, playerVal)
			if err != nil {
				fmt.Println(err)
			}
			if game.IsWin(m.gb, row, col, playerVal) {
				m.win = true
				return m, tea.Quit
			} else {
				m.isPlayer1Turn = !m.isPlayer1Turn
			}
			m.cursor = 0
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	s := "\nConnect-4\n\n"

	var currentPlayer player

	if m.isPlayer1Turn {
		currentPlayer = m.p1
	} else {
		currentPlayer = m.p2
	}

	s += currentPlayer.playerName + "'s turn\n\n"

	renderedCursor := false

	// for i := m.gb.Rows - 1; i >= 0; i-- {
	for i := range m.gb.Rows {
		for k := range m.gb.Columns {
			if !m.win && k == m.cursor && (i+1 > m.gb.Rows-1 || m.gb.Board[i+1][k] != ' ') && renderedCursor == false {
				s += fmt.Sprintf("| %c ", currentPlayer.playerValue)
				renderedCursor = true
			} else {
				s += fmt.Sprintf("| %c ", m.gb.Board[i][k])
			}
		}
		s += fmt.Sprintln("|")
	}

	dashes := strings.Repeat("----", 7)
	s += fmt.Sprintf("%s-\n", dashes)

	s += "\nPress q to quit.\n"

	v := tea.NewView(s)
	v.WindowTitle = "Grocery List"

	return v
}

func main() {
	p := tea.NewProgram(initialModel())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	m := finalModel.(model)

	var winner player

	if m.isPlayer1Turn {
		winner = m.p1
	} else {
		winner = m.p2
	}

	if m.win {
		fmt.Println("\n" + winner.playerName + " is the winner")
	}
}
