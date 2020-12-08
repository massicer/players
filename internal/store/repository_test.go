package store

import (
	"testing"
)

func TestGetPlayerScore(t *testing.T) {

    t.Run("returns score for non existing player", func(t *testing.T) {
		store := InMemoryPlayerStore{}
		got := store.GetPlayerScore("something")
		want := 0

		if got != want {
			t.Fatalf("Got %d, want %d", got, want)
		}
	})
}