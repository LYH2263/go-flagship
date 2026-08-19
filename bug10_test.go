package flagship_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/LYH2263/go-flagship"
)

func TestBug10_CloseFlushesWriterBeforeRelease(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "flags.json")
	s := flagship.New(flagship.WithPersistPath(path))
	if err := s.Upsert(flagship.FlagDef{Key: "dark_mode", Enabled: true, Percent: 50}); err != nil {
		t.Fatal(err)
	}
	if err := s.Close(); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(raw, []byte("dark_mode")) {
		t.Fatalf("persist file missing data after Close (dirty dropped): %q", raw)
	}
}
