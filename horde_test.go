package undead

import "testing"

func TestHorde(t *testing.T) {
	h := NewHorde()
	defer func() {
		z := h.Zombies()
		if len(z) != 1 {
			t.Fatalf("unexpected number of zombies: got=%d want=1", len(z))
		}
	}()

	go func() {
		select {}
	}()
}
