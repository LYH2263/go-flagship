package flagship

func (s *Ship) Stats() Stats {
	s.mu.Lock()
	defer s.mu.Unlock()
	return Stats{
		Flags:     s.table.Len(),
		Upserts:   s.upserts,
		Evaluates: s.evaluates,
		OnCount:   s.onCount,
		OffCount:  s.offCount,
		Closed:    s.closed,
	}
}
