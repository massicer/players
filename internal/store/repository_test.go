package store

import (
	"testing"
)

func TestGetPlayerScore(t *testing.T) {

    t.Run("returns static score", func(t *testing.T) {
		store := InMemoryPlayerStore{}
		got := store.GetPlayerScore("something")
		want := 123

		if got != want {
			t.Fatalf("Got %d, want %d", got, want)
		}
	})
}