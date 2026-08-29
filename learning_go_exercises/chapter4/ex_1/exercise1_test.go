package ex1

import (
	"reflect"
	"testing"
)

func TestCreateSlice(t *testing.T) {
	cases := []struct {
		Size        int
		Description string
		Want        []int
	}{
		{
			10,
			"slice of 10 should have from 0 to 9",
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			5,
			"slice of 5 should have from 0 to 4",
			[]int{0, 1, 2, 3, 4},
		},
		{
			2,
			"slice of size 2 should have from 0 to 1",
			[]int{0, 1},
		},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			got := CreateSlice(c.Size)
			if !reflect.DeepEqual(c.Want, got) {
				t.Errorf("got %v, want %v", got, c.Want)
			}
		})
	}

}
