package store

import ierr "github.com/LYH2263/go-flagship/internal/errors"

// Close 先 Flush 脏缓冲，再关闭 writer。
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return nil
	}
	var first error
	// 先把 dirty 缓冲落盘，再关闭 writer；否则 writer 一关，未刷的 dirty 永久丢失。
	if err := s.flushLocked(); err != nil && first == nil {
		first = err
	}
	if s.writer != nil {
		if err := s.writer.Close(); err != nil && first == nil {
			first = ierr.WrapErr(ierr.ErrPersist, err)
		}
		s.writer = nil
	}
	s.closed = true
	s.dirty = false
	s.path = ""
	return first
}
