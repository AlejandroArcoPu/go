package server

import (
	"github.com/AlejandroArcoPu/go/fullstack-go-htmx/types"
)

type Store interface {
	GetCars() []types.Car
	CreateCar(car types.Car)
	GetCar(vin string) (types.Car, error)
	DeleteCar(vin string) error
}
