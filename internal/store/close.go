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
	// 丢掉 dirty，先关 writer。
	s.dirty = false
	if s.writer != nil {
		if err := s.writer.Close(); err != nil && first == nil {
			first = err
		}
		s.writer = nil
	}
	s.closed = true
	s.path = ""
	return first
}
