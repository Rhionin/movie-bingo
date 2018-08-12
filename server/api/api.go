package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rhionin/movie-bingo/server/bingo"
	"github.com/gorilla/mux"
)

type (
	// API the api server
	API struct {
		Port string
	}
)

var (
	game = bingo.NewGame()
)

// RunServer runs the http server
func (api API) RunServer() error {

	r := mux.NewRouter()
	r.HandleFunc("/api/game", GameHandler)
	r.HandleFunc("/api/board", BoardHandler)

	// fs := http.FileServer(http.Dir("Public/"))
	// http.Handle("/Public/", http.StripPrefix("/Public/", fs))

	fmt.Printf("Running server on port %s\n", api.Port)
	http.ListenAndServe(":"+api.Port, r)

	return nil
}

// GameHandler handler for getting the game definition from the api
func GameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(game)
}

// BoardHandler handler for getting a board definition from the api
func BoardHandler(w http.ResponseWriter, r *http.Request) {
	playerName := "Some player" // TODO Get from request
	color := "Some color"       // TODO Get from request or have the server choose

	board := bingo.NewBoard(playerName, color, game.Events)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(board)
}
