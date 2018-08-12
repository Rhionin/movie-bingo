package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rhionin/movie-bingo/server/bingo"
)

type (
	// API the api server
	API struct {
		Port string
	}
)

// RunServer runs the http server
func (api API) RunServer() error {
	http.HandleFunc("/api/game", GameHandler)

	fs := http.FileServer(http.Dir("Public/"))
	http.Handle("/Public/", http.StripPrefix("/Public/", fs))

	fmt.Printf("Running server on port %s\n", api.Port)
	http.ListenAndServe(":"+api.Port, nil)

	return nil
}

// GameHandler handler for getting the game definition from the api
func GameHandler(w http.ResponseWriter, r *http.Request) {
	game := bingo.NewGame()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(game)
}
