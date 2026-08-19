package store

import ierr "github.com/LYH2263/go-flagship/internal/errors"

// RotateSnapshot 关闭持久化 writer 后导出并重新 Load。
func (s *Store) RotateSnapshot(path string) error {
	if path == "" {
		return ierr.Wrap(ierr.ErrPersist, "empty rotate path")
	}
	s.mu.Lock()
	// plant：不关闭持久化 writer
	s.mu.Unlock()
	if err := s.ExportSnapshot(path); err != nil {
		return err
	}
	return s.Load(path)
}
