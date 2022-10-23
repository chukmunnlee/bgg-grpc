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
	svc "github.com/chukmunnlee/bgg-grpc/grpc/server"
	health "github.com/chukmunnlee/bgg-grpc/grpc/healthcheck"
)

func main() {

	var hc *health.HealthService

	dbFile := flag.String("database", "", "Sqlite3 database file")
	port := flag.Int("port", 50051, "Server's port")
	healthPort := flag.Int("healthPort", -1, "Health check port, -1 is disabled")
	intf := flag.String("interface", "127.0.0.1", "Interface to bind to")
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
		log.Println("Enabling reflection for BGG Server")
		reflection.Register(s)
	}

	connStr := fmt.Sprintf("%s:%d", *intf, *port)
	lis, err := net.Listen("tcp", connStr)
	if nil != err {
		log.Fatalf("Cannot create listener: %v\n", err)
	}
	defer lis.Close()

	log.Printf("Starting BGGService on interface %s", connStr)
	go func() {
		if err := s.Serve(lis); nil != err {
			log.Fatalf("Cannot start service: %v", err)
		}
	}()

	// If healthcheck is enabled
	if *healthPort > 0 {

		hc, err = health.New("127.0.0.1", *port)
		if nil != err {
			log.Fatalf("Cannot start healthcheck: %v", err)
		}

		err = hc.Start(*intf, *healthPort)
		if nil != err {
			log.Fatalf("Cannot start healthcheck: %v", err)
		}
		defer hc.Close()

		if *reflect {
		log.Println("Enabling reflection for healthcheck service")
			reflection.Register(hc.Server)
		}
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Printf("Shutting down server")
}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}
