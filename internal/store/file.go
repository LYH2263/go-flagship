package store

import (
	"os"
	"path/filepath"
)

type fileWriter struct {
	path string
	f    *os.File
}

func newFileWriter(path string) *fileWriter {
	return &fileWriter{path: path}
}

func (w *fileWriter) ensure() error {
	if w.f != nil {
		return nil
	}
	dir := filepath.Dir(w.path)
	if dir != "." && dir != "" {
		_ = os.MkdirAll(dir, 0o755)
	}
	f, err := os.OpenFile(w.path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	w.f = f
	return nil
}

func (w *fileWriter) WriteAll(b []byte) error {
	if err := w.ensure(); err != nil {
		return err
	}
	if err := w.f.Truncate(0); err != nil {
		return err
	}
	if _, err := w.f.Seek(0, 0); err != nil {
		return err
	}
	if _, err := w.f.Write(b); err != nil {
		return err
	}
	return w.f.Sync()
}

func (w *fileWriter) Sync() error {
	if w.f == nil {
		return nil
	}
	return w.f.Sync()
}

func (w *fileWriter) Close() error {
	if w.f == nil {
		return nil
	}
	err := w.f.Close()
	w.f = nil
	return err
}
