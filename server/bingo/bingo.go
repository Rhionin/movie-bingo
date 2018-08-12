package bingo

import "math/rand"

type (
	// Game instance of a bingo game
	Game struct {
		// Events events that can happen in a movie
		Events []string

		// Boards the collection of boards used by players of the game
		Boards []Board
	}

	// Board a bingo board
	Board struct {
		// Player the player using the board
		Player string
		// Cells the matrix of cells on a bingo board
		Cells []Cell
		// Color the color of filled cells
		Color string
	}

	// Cell a cell on a board
	Cell struct {
		// Text the text in the cell
		Text string
		// Filled whether the cell is filled
		Filled bool
	}
)

// NewGame create a new bingo game
func NewGame() Game {

	events := getEvents()

	game := Game{
		Events: events,
		Boards: []Board{},
	}

	return game
}

func getEvents() []string {
	return []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
	}
}

// NewBoard creates a new board
func NewBoard(player string, color string, events []string) Board {
	cells := []Cell{}
	for _, evt := range events {
		cells = append(cells, Cell{
			Text:   evt,
			Filled: false,
		})
	}

	rand.Shuffle(len(cells), func(i, j int) {
		cells[i], cells[j] = cells[j], cells[i]
	})

	return Board{
		Player: player,
		Color:  color,
		Cells:  cells,
	}
}
