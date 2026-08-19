package flagship

// Close 先 Flush/Sync，再 Close store，最后置 nil。
func (s *Ship) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return nil
	}
	s.closed = true

	var first error
	if s.store != nil {
		if err := s.store.Flush(); err != nil && first == nil {
			first = err
		}
		if err := s.store.Sync(); err != nil && first == nil {
			first = err
		}
		if err := s.store.Close(); err != nil && first == nil {
			first = err
		}
	}
	// nil-after-close
	s.store = nil
	return first
}

func (s *Ship) Closed() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.closed
}
