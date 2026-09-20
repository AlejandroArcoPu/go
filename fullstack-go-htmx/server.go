package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AlejandroArcoPu/go/fullstack-go-htmx/types"
)

const jsonContentType = "application/json"

type Store interface {
	GetCars() []types.Car
	CreateCar(car types.Car) error
	GetCar(vin string) (types.Car, error)
	DeleteCar(vin string) error
}

type CarServer struct {
	store Store
	http.Handler
}

func NewCarServer(store Store) *CarServer {
	c := new(CarServer)

	c.store = store

	server := http.NewServeMux()
	server.HandleFunc("/cars", http.HandlerFunc(c.carsHandler))
	server.HandleFunc("/car", http.HandlerFunc(c.carHandler))

	c.Handler = server

	return c
}

func (c *CarServer) carsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Add("Content-Type", jsonContentType)
		json.NewEncoder(w).Encode(c.store.GetCars())
	}
}

func (c *CarServer) carHandler(w http.ResponseWriter, r *http.Request) {
	vin := strings.TrimPrefix(r.URL.Path, "/car/")

	switch r.Method {
	case http.MethodGet:
		c.showCar(w, vin)
	case http.MethodPost:
		c.createCar(w, r)
	}

}

func (c *CarServer) showCar(w http.ResponseWriter, vin string) {
	w.Header().Add("Content-Type", jsonContentType)

	car, err := c.store.GetCar(vin)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(car)
}

func (c *CarServer) createCar(w http.ResponseWriter, r *http.Request) {
	car := types.Car{}
	json.NewDecoder(r.Body).Decode(&car)
	err := c.store.CreateCar(car)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}
