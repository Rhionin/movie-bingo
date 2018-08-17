package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rhionin/movie-bingo/server/bingo"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type (
	// API the api server
	API struct {
		Port string
	}
)

// RunServer runs the http server
func (api API) RunServer() error {

	r := mux.NewRouter()
	r.HandleFunc("/api/games", CreateGameHandler).Methods("POST")
	r.HandleFunc("/api/games/{gameID}", GetGameHandler).Methods("GET")
	// r.HandleFunc("/api/board", BoardHandler)

	fs := http.FileServer(http.Dir("Public/"))
	r.PathPrefix("/Public/").Handler(http.StripPrefix("/Public/", fs))

	fmt.Printf("Running server on port %s\n", api.Port)
	http.ListenAndServe(":"+api.Port, r)

	return nil
}

// CreateGameHandler handler for creating a new game
func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	game := bingo.NewGame()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(game)
}

// GetGameHandler handler for getting a game by ID
func GetGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := uuid.FromStringOrNil(vars["gameID"])

	game, err := bingo.GetGame(gameID)
	if err == bingo.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(struct{ Message string }{Message: "Not found"})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(game)
}

// // BoardHandler handler for getting a board definition from the api
// func BoardHandler(w http.ResponseWriter, r *http.Request) {
// 	playerName := "Some player" // TODO Get from request
// 	color := "Some color"       // TODO Get from request or have the server choose

// 	board := bingo.NewBoard(playerName, color, games[0].Events)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(board)
// }
