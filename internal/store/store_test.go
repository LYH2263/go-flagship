package store

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/LYH2263/go-flagship/internal/clock"
)

func TestSaveLoadRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "flags.json")
	s := New(path, clock.Real{})
	err := s.SaveSnapshot(&Snapshot{Flags: []FlagSnap{{Key: "x", Enabled: true, Percent: 50}}})
	if err != nil {
		t.Fatal(err)
	}
	if err := s.Close(); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(raw) == 0 {
		t.Fatal("empty file")
	}
}
