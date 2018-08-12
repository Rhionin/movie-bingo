#!/bin/bash

# Run from the movie-bingo/server directory
GOOS=windows GOARCH=amd64 go build cmd/server/*.go