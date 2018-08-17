package bingo

import (
	"fmt"
	"math/rand"

	"github.com/satori/go.uuid"
)

var (
	// ErrNotFound returned when a requested resource is not found
	ErrNotFound = fmt.Errorf("Not found")

	games = make(map[uuid.UUID]Game)
)

type (
	// Game instance of a bingo game
	Game struct {
		// ID The game ID
		ID uuid.UUID

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
	// TODO get different types of games (disney, hallmark, marvel, conference...)
	id := uuid.NewV4()
	events := getEvents()

	game := Game{
		ID:     id,
		Events: events,
		Boards: []Board{},
	}

	// Add game to in-memory game registry
	games[id] = game

	return game
}

// GetGame gets a bingo game by ID
func GetGame(id uuid.UUID) (Game, error) {
	game, ok := games[id]
	if !ok {
		return game, ErrNotFound
	}

	return game, nil
}

// NewBoard creates a new board
func (g Game) NewBoard(player string, color string, events []string) Board {
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

	board := Board{
		Player: player,
		Color:  color,
		Cells:  cells,
	}

	g.Boards = append(g.Boards, board)

	return board
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
