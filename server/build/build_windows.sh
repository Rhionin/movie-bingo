#!/bin/bash

# Run from the movie-bingo directory, and run the binary from there as well
GOOS=windows GOARCH=amd64 go build -o server.exe server/cmd/server/*.go