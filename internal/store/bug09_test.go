package store

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBug09_SnapshotRotateClosesWriters(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "snap.json")
	s := New(path, nil)
	defer s.Close()
	if err := s.SaveSnapshot(&Snapshot{Flags: []FlagSnap{{Key: "a", Enabled: true, Percent: 50}}}); err != nil {
		t.Fatal(err)
	}
	s.MarkDirty()
	for i := 0; i < 3; i++ {
		if err := s.RotateSnapshot(path); err != nil {
			t.Fatalf("rotate %d: %v", i, err)
		}
	}
	replacement := []byte(`{"flags":[]}`)
	if err := os.WriteFile(path, replacement, 0o644); err != nil {
		t.Fatalf("cannot replace snapshot after rotate (writers not closed): %v", err)
	}
	if err := s.Load(path); err != nil {
		t.Fatal(err)
	}
	s.mu.Lock()
	n := 0
	if s.snap != nil {
		n = len(s.snap.Flags)
	}
	s.mu.Unlock()
	if n != 0 {
		t.Fatalf("reload after replace still has flags=%d (stale/locked snapshot)", n)
	}
}
