package undead

import "testing"

func TestHoard(t *testing.T) {
	h := NewHoard()
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
