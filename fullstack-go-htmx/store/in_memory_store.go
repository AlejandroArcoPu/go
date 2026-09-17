package store

import (
	"errors"

	"github.com/AlejandroArcoPu/go/fullstack-go-htmx/types"
)

type InMemoryCarStore struct {
	cars []types.Car
}

var ErrorCarNotFound = errors.New("The car is not found")

func NewInMemoryCarStore(cars []types.Car) *InMemoryCarStore {
	return &InMemoryCarStore{cars}
}

func (i *InMemoryCarStore) GetCars() []types.Car {
	return i.cars
}

func (i *InMemoryCarStore) CreateCar(car types.Car) {
	i.cars = append(i.cars, car)
}

func (i *InMemoryCarStore) GetCar(vin string) (types.Car, error) {
	for _, car := range i.cars {
		if car.Vin == vin {
			return car, nil
		}
	}
	return types.Car{}, ErrorCarNotFound
}
