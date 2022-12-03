package main

type Store interface {
	Open() error
	Close() error
}
