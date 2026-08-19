package flagship

import ierr "github.com/LYH2263/go-flagship/internal/errors"

func (s *Ship) Flush() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.checkOpenLocked(); err != nil {
		return err
	}
	if s.persistPath == "" {
		return nil
	}
	if err := s.store.Flush(); err != nil {
		return ierr.WrapErr(ErrPersist, err)
	}
	return nil
}

func (s *Ship) Sync() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.checkOpenLocked(); err != nil {
		return err
	}
	if s.store == nil {
		return ErrNilStore
	}
	if err := s.store.Sync(); err != nil {
		return ierr.WrapErr(ErrSync, err)
	}
	return nil
}
