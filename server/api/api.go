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

	fs := http.FileServer(http.Dir("Public/"))
	http.Handle("/Public/", http.StripPrefix("/Public/", fs))

	fmt.Printf("Running server on port %s\n", api.Port)
	http.ListenAndServe(":"+api.Port, nil)

	return nil
}
