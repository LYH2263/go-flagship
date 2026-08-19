package flagship_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/LYH2263/go-flagship"
)

func TestBug06_PersistFailureRollsBack(t *testing.T) {
	dir := t.TempDir()
	bad := filepath.Join(dir, "not-a-file-dir")
	if err := os.MkdirAll(bad, 0o755); err != nil {
		t.Fatal(err)
	}
	s := flagship.New(flagship.WithPersistPath(bad))
	defer s.Close()
	err := s.Upsert(flagship.FlagDef{Key: "x", Enabled: true, Percent: 100})
	if err == nil {
		t.Fatal("expected persist error")
	}
	if !errors.Is(err, flagship.ErrPersist) {
		t.Fatalf("want ErrPersist, got %v", err)
	}
	list, _ := s.List()
	if len(list) != 0 {
		t.Fatalf("partial upsert left flags=%v", list)
	}
}
