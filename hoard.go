// Package undead provides a convenience wrapper for goroutineleak
// profiles.
//
// Currently requires GOEXPERIMENT=goroutineleakprofile.
package undead

import (
	"runtime/pprof"
	"slices"
	"strings"
	"time"
)

// Hoard is a goroutineleak profile handle.
type Hoard struct {
	prof *pprof.Profile
}

// NewHoard returns an undead hoard.
func NewHoard() *Hoard {
	return &Hoard{prof: pprof.Lookup("goroutineleak")}
}

// Zombies returns the details of leaked goroutines.
func (h *Hoard) Zombies() []string {
	if h.prof == nil {
		return nil
	}
	time.Sleep(2 * time.Second)
	var buf strings.Builder
	h.prof.WriteTo(&buf, 2)
	leaked := slices.DeleteFunc(strings.Split(buf.String(), "\n\n"), func(s string) bool {
		return !strings.Contains(s, "(leaked)")
	})
	if leaked == nil {
		leaked = []string{}
	}
	for i, l := range leaked {
		leaked[i] = strings.TrimSpace(l)
	}
	return leaked
}
