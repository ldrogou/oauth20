package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ldrogou/goauth20/store"
)

func main() {
	fmt.Println("OAuth RCA")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s)\n", err)
		os.Exit(1)
	}
}

func run() error {
	srv := newServer()
	srv.store = &store.DbStore{}

	err := srv.store.Open()
	if err != nil {
		return err
	}
	defer srv.store.Close()

	http.HandleFunc("/", srv.serveHTTP)

	port := 8080
	log.Printf("servering http port %v", port)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}
