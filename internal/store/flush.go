package store

import (
	"encoding/json"

	ierr "github.com/LYH2263/go-flagship/internal/errors"
)

func (s *Store) Flush() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return ierr.ErrClosed
	}
	return s.flushLocked()
}

func (s *Store) flushLocked() error {
	if s.path == "" || !s.dirty {
		return nil
	}
	if s.writer == nil {
		s.writer = newFileWriter(s.path)
	}
	raw, err := json.MarshalIndent(s.snap, "", "  ")
	if err != nil {
		return ierr.WrapErr(ierr.ErrFlush, err)
	}
	if err := s.writer.WriteAll(raw); err != nil {
		return ierr.WrapErr(ierr.ErrPersist, err)
	}
	s.dirty = false
	return nil
}

func (s *Store) Sync() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return ierr.ErrClosed
	}
	if s.writer == nil {
		return nil
	}
	if err := s.writer.Sync(); err != nil {
		return ierr.WrapErr(ierr.ErrSync, err)
	}
	return nil
}
