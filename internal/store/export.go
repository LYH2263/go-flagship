package store

import (
	"encoding/json"
	"os"
	"path/filepath"

	ierr "github.com/LYH2263/go-flagship/internal/errors"
)

// SnapshotSession 导出会话：必须 Close 后才能安全替换目标快照。
type SnapshotSession struct {
	f    *os.File
	tmp  string
	dst  string
	done bool
}

func (s *Store) BeginSnapshotExport(dst string) (*SnapshotSession, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return nil, ierr.ErrClosed
	}
	if dst == "" {
		return nil, ierr.Wrap(ierr.ErrPersist, "empty export path")
	}
	raw, err := json.MarshalIndent(s.snap, "", "  ")
	if err != nil {
		return nil, ierr.WrapErr(ierr.ErrPersist, err)
	}
	dir := filepath.Dir(dst)
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, ierr.WrapErr(ierr.ErrPersist, err)
		}
	}
	tmp := dst + ".export-tmp"
	f, err := os.OpenFile(tmp, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o644)
	if err != nil {
		return nil, ierr.WrapErr(ierr.ErrPersist, err)
	}
	if _, err := f.Write(raw); err != nil {
		_ = f.Close()
		_ = os.Remove(tmp)
		return nil, ierr.WrapErr(ierr.ErrPersist, err)
	}
	return &SnapshotSession{f: f, tmp: tmp, dst: dst}, nil
}

func (sess *SnapshotSession) Close() error {
	if sess == nil || sess.done {
		return nil
	}
	sess.done = true
	// Sync+Close 写句柄后再 Rename，否则 Windows 报 sharing violation。
	if sess.f != nil {
		if err := sess.f.Sync(); err != nil {
			_ = sess.f.Close()
			sess.f = nil
			return ierr.WrapErr(ierr.ErrSync, err)
		}
		if err := sess.f.Close(); err != nil {
			sess.f = nil
			return ierr.WrapErr(ierr.ErrPersist, err)
		}
		sess.f = nil
	}
	_ = os.Remove(sess.dst)
	if err := os.Rename(sess.tmp, sess.dst); err != nil {
		return ierr.WrapErr(ierr.ErrPersist, err)
	}
	return nil
}

func (s *Store) ExportSnapshot(dst string) error {
	sess, err := s.BeginSnapshotExport(dst)
	if err != nil {
		return err
	}
	return sess.Close()
}
