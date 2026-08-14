package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("bad")
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	dictionary.Add("good", "very satisfactory, enjoyable, pleasant, or interesting")

	want := "very satisfactory, enjoyable, pleasant, or interesting"
	got, err := dictionary.Search("good")

	if err != nil {
		t.Fatalf("should find the added word: %s", "good")
	}

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
