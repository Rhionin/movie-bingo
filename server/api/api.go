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

// RunServer runs the http server
func (api API) RunServer() error {

	r := mux.NewRouter()
	r.HandleFunc("/api/games", CreateGameHandler).Methods("POST")
	r.HandleFunc("/api/games/{gameID}", GetGameHandler).Methods("GET")
	r.HandleFunc("/api/games/{gameID}/boards", CreateBoardHandler).Methods("POST")

	fs := http.FileServer(http.Dir("ui/dist/movie-bingo"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

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
	game, err := getGameFromRequest(w, r)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(game)
}

// CreateBoardHandler handler for creating a board in a bingo game
func CreateBoardHandler(w http.ResponseWriter, r *http.Request) {
	game, err := getGameFromRequest(w, r)
	if err != nil {
		return
	}

	// b, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	failWithError(w, err)
	// 	return
	// }

	playerName := "Some player" // TODO Get from request
	color := "Some color"       // TODO Get from request or have the server choose

	board := game.NewBoard(playerName, color)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(board)
}

func getGameFromRequest(w http.ResponseWriter, r *http.Request) (bingo.Game, error) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	game, err := bingo.GetGame(gameID)

	if err == bingo.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(struct{ Message string }{Message: "Not found"})
		return game, err
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return game, err
	}

	return game, err
}
