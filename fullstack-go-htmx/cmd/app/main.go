package main

import (
	"log"
	"net/http"

	server "github.com/AlejandroArcoPu/go/fullstack-go-htmx"
	"github.com/AlejandroArcoPu/go/fullstack-go-htmx/store"
	"github.com/AlejandroArcoPu/go/fullstack-go-htmx/types"
)

func main() {
	cars := []types.Car{{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}}
	store := store.NewInMemoryCarStore(cars)
	server := server.NewCarServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("server can't start, %v", err)
	}
}
