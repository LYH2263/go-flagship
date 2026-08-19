package store

import (
	"os"
	"path/filepath"
)

func BackupFile(src, dstDir string) error {
	raw, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		return err
	}
	dst := filepath.Join(dstDir, filepath.Base(src)+".bak")
	return os.WriteFile(dst, raw, 0o644)
}
