package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	db "github.com/chukmunnlee/bgg-grpc/data"
	svc "github.com/chukmunnlee/bgg-grpc/grpc"
)

func main() {

	dbFile := flag.String("database", "", "Sqlite3 database file")
	port := flag.Int("port", 50051, "Server's port")
	reflect := flag.Bool("reflect", false, "Enable reflection")
	flag.Parse()

	if isNull(*dbFile) {
		log.Fatalf("Missing sqlite3 database file")
	}

	if _, err := os.Stat(*dbFile); nil != err {
		log.Fatalf("Database does not exists: %s, %v", *dbFile, err)
	}

	bggdb := db.New(*dbFile)
	if err := bggdb.Open(); nil != err {
		log.Fatalf("Cannot open sqlite3 file: %v", err)
	}
	defer bggdb.Close()

	log.Printf("Using %s as database", *dbFile)

	s := grpc.NewServer()
	bggSvc := svc.New(bggdb)
	svc.RegisterBGGServiceServer(s, &bggSvc)

	if *reflect {
		log.Println("Enabling reflection")
		reflection.Register(s)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if nil != err {
		log.Fatalf("Cannot create listener: %v\n", err)
	}
	defer lis.Close()

	log.Printf("Starting BGGService on port %d", *port)
	go func() {
		if err := s.Serve(lis); nil != err {
			log.Fatalf("Cannot start service: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Printf("Shutting down server")
}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}
