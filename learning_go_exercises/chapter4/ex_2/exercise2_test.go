package ex2

import (
	"bytes"
	"testing"
)

func TestDivisible(t *testing.T) {
	cases := []struct {
		Number      int
		Description string
		Want        string
	}{
		{
			2,
			"divisible of 2 should print two",
			"Two!",
		},
		{
			3,
			"divisible of 3 should print three",
			"Three!",
		},
		{
			5,
			"divisible of 5 should print never mind",
			"Never mind",
		},
		{
			6,
			"divisible of 6 should print six",
			"Six!",
		},
		{
			18,
			"divisible of 6 should print six",
			"Six!",
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := &bytes.Buffer{}
			Divisible(test.Number, got)

			if got.String() != test.Want {
				t.Errorf("got %s, want %s", got, test.Want)
			}
		})
	}
}
