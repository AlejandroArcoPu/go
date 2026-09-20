package store_test

import (
	"reflect"
	"testing"

	"github.com/AlejandroArcoPu/go/fullstack-go-htmx/store"
	"github.com/AlejandroArcoPu/go/fullstack-go-htmx/types"
)

func TestInMemoryStore(t *testing.T) {
	t.Run("get cars", func(t *testing.T) {
		want := []types.Car{{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}}
		store := store.NewInMemoryCarStore(want)

		got := store.GetCars()

		assertCars(t, got, want)
	})

	t.Run("create car", func(t *testing.T) {
		car := types.Car{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}
		store := store.NewInMemoryCarStore([]types.Car{})

		err := store.CreateCar(car)

		got := store.GetCars()
		want := []types.Car{{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}}

		assertNoError(t, err)
		assertCars(t, got, want)
	})

	t.Run("create a repeated car", func(t *testing.T) {
		cars := []types.Car{{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}}
		car := types.Car{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}
		store := store.NewInMemoryCarStore(cars)

		err := store.CreateCar(car)

		got := store.GetCars()

		assertError(t, err)
		assertCars(t, got, cars)
	})

	t.Run("get car by vin", func(t *testing.T) {
		cars := []types.Car{{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}}
		store := store.NewInMemoryCarStore(cars)

		got, err := store.GetCar("123")
		want := types.Car{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}

		assertNoError(t, err)
		assertCars(t, got, want)
	})

	t.Run("get car that doesn't exist", func(t *testing.T) {
		cars := []types.Car{}
		store := store.NewInMemoryCarStore(cars)

		_, err := store.GetCar("123")
		assertError(t, err)
	})

	t.Run("delete car by vin", func(t *testing.T) {
		cars := []types.Car{{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}}
		store := store.NewInMemoryCarStore(cars)

		err := store.DeleteCar("123")

		got := store.GetCars()
		want := []types.Car{}

		assertNoError(t, err)
		assertCars(t, got, want)
	})

	t.Run("delete car that doesn't exist", func(t *testing.T) {
		cars := []types.Car{}
		store := store.NewInMemoryCarStore(cars)

		err := store.DeleteCar("123")

		assertError(t, err)
	})
}

func assertCars[T any](t testing.TB, got, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("didn't expect an error but got one, %v", err)
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("expected an error but didn't have any")
	}
}
