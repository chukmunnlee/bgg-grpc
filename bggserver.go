package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	db "github.com/chukmunnlee/bgg-grpc/data"
)

func main() {

	dbFile := flag.String("database", "", "sqlite3 database file")
	flag.Parse()

	if isNull(*dbFile) {
		log.Fatalf("Missing sqlite3 database file")
	}

	bggdb := db.New(*dbFile)
	if err := bggdb.Open(); nil != err {
		log.Fatalf("Cannot open sqlite3 file: %v", err)
	}
	defer bggdb.Close()

	games, err := bggdb.FindBoardgameByName(context.Background(), "carcassonne", 10, 0)
	if nil != err {
		log.Fatalf("Query error: %v", err)
	}

	for _, g := range *games {
		fmt.Printf("GameId: %d, Name: %s\n", g.GameId, g.Name)
	}
}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}
