package store

import (
	"encoding/json"
	"io"
	"os"
)

type fileSnapshotReader struct {
	f *os.File
}

func (r *fileSnapshotReader) Read(p []byte) (int, error) {
	return r.f.Read(p)
}

func (r *fileSnapshotReader) Close() error {
	if r.f == nil {
		return nil
	}
	err := r.f.Close()
	r.f = nil
	return err
}

// Load 打开快照文件；必须 Close Reader。
func (s *Store) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	r := &fileSnapshotReader{f: f}
	return s.loadFromReader(r)
}

func (s *Store) loadFromReader(r io.Reader) error {
	var snap Snapshot
	dec := json.NewDecoder(r)
	if err := dec.Decode(&snap); err != nil {
		if err == io.EOF {
			s.mu.Lock()
			s.snap = &Snapshot{}
			s.mu.Unlock()
			return nil
		}
		return err
	}
	s.mu.Lock()
	s.snap = cloneSnapshot(&snap)
	s.dirty = false
	s.mu.Unlock()
	return nil
}

func LoadFile(path string) (*Snapshot, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := &fileSnapshotReader{f: f}
	defer r.Close()
	var snap Snapshot
	if err := json.NewDecoder(r).Decode(&snap); err != nil {
		if err == io.EOF {
			return &Snapshot{}, nil
		}
		return nil, err
	}
	return cloneSnapshot(&snap), nil
}
