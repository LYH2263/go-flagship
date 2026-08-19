package store

func (s *Store) FlagCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.snap == nil {
		return 0
	}
	return len(s.snap.Flags)
}
