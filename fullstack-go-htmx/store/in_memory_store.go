package store

import (
	"errors"

	"github.com/AlejandroArcoPu/go/fullstack-go-htmx/types"
)

type InMemoryCarStore struct {
	cars []types.Car
}

var (
	ErrorCarNotFound     = errors.New("The car is not found")
	ErrorCarAlreadyExist = errors.New("The car already exist")
)

func NewInMemoryCarStore(cars []types.Car) *InMemoryCarStore {
	return &InMemoryCarStore{cars}
}

func (m *InMemoryCarStore) GetCars() []types.Car {
	return m.cars
}

func (m *InMemoryCarStore) CreateCar(new types.Car) error {
	car, _ := m.find(new.Vin)
	if car != nil {
		return ErrorCarAlreadyExist
	}
	m.cars = append(m.cars, new)
	return nil
}

func (m *InMemoryCarStore) GetCar(vin string) (types.Car, error) {
	car, _ := m.find(vin)
	if car == nil {
		return types.Car{}, ErrorCarNotFound
	}
	return *car, nil
}

func (m *InMemoryCarStore) DeleteCar(vin string) error {
	_, index := m.find(vin)
	if index == -1 {
		return ErrorCarNotFound
	}
	m.cars = append(m.cars[:index], m.cars[index+1:]...)
	return nil
}

func (m *InMemoryCarStore) find(vin string) (*types.Car, int) {
	for i, car := range m.cars {
		if car.Vin == vin {
			return &car, i
		}
	}
	return nil, -1
}
