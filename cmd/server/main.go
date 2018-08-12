package main

import (
	"os"

	"github.com/Rhionin/movie-bingo/api"
)

func main() {

	port := getenvOrDefault("PORT", "8018")

	api := api.API{
		Port: port,
	}

	api.RunServer()
}

func getenvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
