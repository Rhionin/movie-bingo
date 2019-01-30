package main

import (
	"os"
	"time"

	"github.com/Rhionin/movie-bingo/server/api"
)

func main() {

	port := getenvOrDefault("PORT", "8018")

	api := api.API{
		Port: port,
	}

	go api.RunServer()
	time.Sleep(1 * time.Hour)
}

func getenvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
