package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	port := getenvOrDefault("PORT", "8018")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Running server on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func getenvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
