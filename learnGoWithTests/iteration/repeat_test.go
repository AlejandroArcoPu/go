package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 6)
	}
}
