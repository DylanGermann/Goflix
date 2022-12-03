package main

type server struct {
	store Store
}

func newServer() *server {
	s := &server{}
	return s
}
