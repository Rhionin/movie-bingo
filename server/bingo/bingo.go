package bingo

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/satori/go.uuid"
)

var (
	// ErrNotFound returned when a requested resource is not found
	ErrNotFound = fmt.Errorf("Not found")

	games = make(map[string]Game)
)

func init() {
	if os.Getenv("CREATE_MOCK_GAME") == "true" {
		game := NewGame()
		fmt.Println("Mock game created: " + game.ID)
		game.NewBoard("Bruce Wayne", "Black")
	}
}

type (
	// Game instance of a bingo game
	Game struct {
		// ID The game ID
		ID string
		// Events events that can happen in a movie
		Events []string
		// Boards the collection of boards used by players of the game
		Boards map[uuid.UUID]Board
	}

	// Board a bingo board
	Board struct {
		// ID The board ID
		ID uuid.UUID
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
	id := newGameID()
	events := getEvents()

	game := Game{
		ID:     id,
		Events: events,
		Boards: make(map[uuid.UUID]Board),
	}

	// Add game to in-memory game registry
	games[id] = game

	return game
}

// GetGame gets a bingo game by ID
func GetGame(id string) (Game, error) {
	game, ok := games[id]
	if !ok {
		return game, ErrNotFound
	}

	return game, nil
}

// NewBoard creates a new board
func (g Game) NewBoard(player string, color string) Board {
	cells := []Cell{}
	for _, evt := range g.Events {
		cells = append(cells, Cell{
			Text:   evt,
			Filled: false,
		})
	}

	rand.Shuffle(len(cells), func(i, j int) {
		cells[i], cells[j] = cells[j], cells[i]
	})

	boardID := uuid.NewV4()
	board := Board{
		ID:     boardID,
		Player: player,
		Color:  color,
		Cells:  cells,
	}

	g.Boards[boardID] = board

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

func newGameID() string {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 4)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
