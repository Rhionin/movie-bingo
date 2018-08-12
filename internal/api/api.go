package api

import (
	"fmt"
	"net/http"
)

type (
	// API the api server
	API struct {
		Port string
	}
)

// RunServer runs the http server
func (api API) RunServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Running server on port %s\n", api.port)
	http.ListenAndServe(":"+api.port, nil)

	return nil
}
