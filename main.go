package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/joho/godotenv"
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
			if m.cursor == 0 {
				m.cursor = m.board.Columns - 1
			} else if m.cursor > 0 {
				m.cursor--
			}
			log.Printf("Cursor: %d\n", m.cursor)
		case "right", "l":
			if m.cursor == m.board.Columns-1 {
				m.cursor = 0
			} else if m.cursor < m.board.Columns-1 {
				m.cursor++
			}
		case "enter", "space":
			m.err = nil
			var playerVal rune
			if m.isPlayer1Turn {
				playerVal = m.p1.playerValue
			} else {
				playerVal = m.p2.playerValue
			}
			row, col, err := m.board.PlaceOnBoard(m.cursor+1, playerVal)
			if err != nil {
				m.err = err
				break
			}
			if game.IsWin(m.board, row, col, playerVal) {
				m.win = true
				if playerVal == m.p1.playerValue {
					m.winner = m.p1
				} else {
					m.winner = m.p2
				}
				return m, tea.Quit
			} else if game.IsDraw(m.board) {
				m.draw = true
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

	// cursorRendered := false

	if !m.win {
		s += fmt.Sprintf(strings.Repeat("    ", m.cursor)+"  %c  \n", currentPlayer.playerValue)
	} else {
		s += "\n"
	}

	for i := range m.board.Rows {
		for k := range m.board.Columns {
			s += fmt.Sprintf("| %c ", m.board.Board[i][k])
		}
		s += fmt.Sprintln("|")
	}

	dashes := strings.Repeat("----", m.board.Columns)
	s += fmt.Sprintf("%s-\n", dashes)

	s += "\nPress q to quit.\n"
	if m.err != nil {
		s += "\n" + m.err.Error() + "\n"
	}
	v := tea.NewView(s)
	v.WindowTitle = "Connect 4"

	return v
}

func main() {
	godotenv.Load()
	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		if _, err := tea.LogToFile(logfilePath, "simple"); err != nil {
			log.Fatal(err)
		}
	}

	p := tea.NewProgram(initialModel())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	m := finalModel.(model)

	if m.win {
		fmt.Println("\n" + m.winner.playerName + " is the winner \n")
	}

	if m.draw {
		fmt.Println("\nThis game is a draw")
	}
}
