package store

import (
	"sync"

	"github.com/LYH2263/go-flagship/internal/clock"
)

type Store struct {
	mu     sync.Mutex
	path   string
	clk    clock.Clock
	dirty  bool
	closed bool
	snap   *Snapshot
	writer *fileWriter
}

func New(path string, clk clock.Clock) *Store {
	if clk == nil {
		clk = clock.Real{}
	}
	s := &Store{path: path, clk: clk, snap: &Snapshot{}}
	if path != "" {
		s.writer = newFileWriter(path)
	}
	return s
}
