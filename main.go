package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ldrogou/goauth20/routeserv"
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
	srv := routeserv.NewServer()
	srv.Store = &store.DbStore{}

	err := srv.Store.Open()
	if err != nil {
		return err
	}
	defer srv.Store.Close()

	http.HandleFunc("/", srv.ServeHTTP)

	port := 8090
	log.Printf("servering http port %v", port)
	err = http.ListenAndServe(":8090", nil)
	if err != nil {
		return err
	}

	return nil
}
