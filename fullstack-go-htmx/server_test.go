package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	server "github.com/AlejandroArcoPu/go/fullstack-go-htmx"
	"github.com/AlejandroArcoPu/go/fullstack-go-htmx/store"
	"github.com/AlejandroArcoPu/go/fullstack-go-htmx/types"
)

func TestGETCars(t *testing.T) {
	cars := []types.Car{{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}}

	request := newCarsRequest()
	response := httptest.NewRecorder()
	store := store.NewInMemoryCarStore(cars)

	server := server.NewCarServer(store)
	server.ServeHTTP(response, request)

	got := getCarsFromResponse(response.Body)

	assertContentType(t, response)
	assertCars(t, got, cars)
	assertResponseCode(t, response.Code, 200)
}

func TestGETCar(t *testing.T) {
	t.Run("found car", func(t *testing.T) {
		cars := []types.Car{{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}}

		request := newCarRequest("123")
		response := httptest.NewRecorder()
		store := store.NewInMemoryCarStore(cars)

		server := server.NewCarServer(store)
		server.ServeHTTP(response, request)

		got := getCarFromResponse(response.Body)
		want := types.Car{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}

		assertContentType(t, response)
		assertCar(t, got, want)
		assertResponseCode(t, response.Code, 200)
	})

	t.Run("not found car", func(t *testing.T) {
		cars := []types.Car{}

		request := newCarRequest("123")
		response := httptest.NewRecorder()
		store := store.NewInMemoryCarStore(cars)

		server := server.NewCarServer(store)
		server.ServeHTTP(response, request)

		assertContentType(t, response)
		assertResponseCode(t, response.Code, 404)
	})

}

func TestPOSTCar(t *testing.T) {
	t.Run("not found car", func(t *testing.T) {
		car := types.Car{Vin: "123", Brand: "Seat", Model: "Leon", Year: "2010"}
		cars := []types.Car{}

		request := newPostCarRequest(car)
		response := httptest.NewRecorder()
		store := store.NewInMemoryCarStore(cars)

		server := server.NewCarServer(store)
		server.ServeHTTP(response, request)

		assertResponseCode(t, response.Code, 201)
	})

	t.Run("found car", func(t *testing.T) {

	})
}

func newPostCarRequest(car types.Car) (request *http.Request) {
	request, _ = http.NewRequest(http.MethodPost, "/car", nil)
	return
}

func newCarRequest(vin string) (request *http.Request) {
	request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/car/%s", vin), nil)
	return
}

func newCarsRequest() (request *http.Request) {
	request, _ = http.NewRequest(http.MethodGet, "/cars", nil)
	return
}

func getCarsFromResponse(body *bytes.Buffer) (cars []types.Car) {
	json.NewDecoder(body).Decode(&cars)
	return
}

func getCarFromResponse(body *bytes.Buffer) (car types.Car) {
	json.NewDecoder(body).Decode(&car)
	return
}

func assertContentType(t testing.TB, response http.ResponseWriter) {
	t.Helper()

	if response.Header().Get("content-type") != "application/json" {
		t.Errorf("expected content-type json, but didn't get it")
	}
}

func assertResponseCode(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertCars(t testing.TB, got, want []types.Car) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertCar(t testing.TB, got, want types.Car) {
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
