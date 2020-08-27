# Simple Math Quiz Game

## How to run?

1. go build main.go, then ./main 
2. go run main.go

## What flags to include?

This game takes 2 flags or "command line arguments"
1. --csv (a filename)
Ex. main.go --csv=problems.csv
2. --limit (an integer that determines how many seconds the game runs for)
Ex. main.go --limit=10

General Example:
go run main.go --csv=problems.csv --limit=10
This will run the go executable with the csv file problems.csv and set a time
limit of 10 seconds
