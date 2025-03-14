package main

import (
	"log"
)

func main() {
	store, err := NewPostGresStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := NEWAPIServer(":3000", store)
	server.Run()
}
