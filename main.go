package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	srv := newServer()
	srv.store = &dbStore{}
	err := srv.store.Open()
	if err != nil {
		return err
	}
	movies, err := srv.store.GetMovies()
	if err != nil {
		return err
	}
	fmt.Printf("movies=%v\n", movies)

	defer srv.store.Close()
	return nil
}
