package store

import ierr "github.com/LYH2263/go-flagship/internal/errors"

// RotateSnapshot 关闭持久化 writer 后导出并重新 Load。
func (s *Store) RotateSnapshot(path string) error {
	if path == "" {
		return ierr.Wrap(ierr.ErrPersist, "empty rotate path")
	}
	s.mu.Lock()
	// 关闭持久化 writer，否则 Rename 到目标路径会与打开的写句柄冲突。
	if s.writer != nil {
		_ = s.writer.Sync()
		_ = s.writer.Close()
		s.writer = nil
	}
	s.mu.Unlock()
	if err := s.ExportSnapshot(path); err != nil {
		return err
	}
	return s.Load(path)
}
