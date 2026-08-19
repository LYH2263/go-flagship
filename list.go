package flagship

func (s *Ship) List() ([]FlagView, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.checkOpenLocked(); err != nil {
		return nil, err
	}
	recs := s.table.Export()
	out := make([]FlagView, 0, len(recs))
	for _, rec := range recs {
		out = append(out, FlagView{
			Key:         rec.Key,
			Enabled:     rec.Enabled,
			Percent:     rec.Percent,
			Rules:       cloneRules(fromFlagRules(rec.Rules)),
			Description: rec.Description,
			UpdatedAt:   rec.UpdatedAt,
		})
	}
	return out, nil
}

func (s *Ship) Get(key string) (FlagView, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.checkOpenLocked(); err != nil {
		return FlagView{}, err
	}
	rec, ok := s.table.Get(key)
	if !ok {
		return FlagView{}, ErrNotFound
	}
	return FlagView{
		Key:         rec.Key,
		Enabled:     rec.Enabled,
		Percent:     rec.Percent,
		Rules:       cloneRules(fromFlagRules(rec.Rules)),
		Description: rec.Description,
		UpdatedAt:   rec.UpdatedAt,
	}, nil
}

func (s *Ship) Snapshot() ([]FlagDef, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.checkOpenLocked(); err != nil {
		return nil, err
	}
	recs := s.table.Export()
	out := make([]FlagDef, 0, len(recs))
	for _, rec := range recs {
		out = append(out, FlagDef{
			Key:         rec.Key,
			Enabled:     rec.Enabled,
			Percent:     rec.Percent,
			Rules:       cloneRules(fromFlagRules(rec.Rules)),
			Description: rec.Description,
			UpdatedAt:   rec.UpdatedAt,
		})
	}
	return out, nil
}

func (s *Ship) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.checkOpenLocked(); err != nil {
		return err
	}
	if !s.table.Delete(key) {
		return ErrNotFound
	}
	if s.persistPath != "" {
		return s.persistLocked()
	}
	return nil
}
