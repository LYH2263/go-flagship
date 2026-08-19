package store

import ierr "github.com/LYH2263/go-flagship/internal/errors"

func (s *Store) SaveSnapshot(snap *Snapshot) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return ierr.ErrClosed
	}
	s.snap = cloneSnapshot(snap)
	s.dirty = true
	// 不在此处落盘；由 Flush / Close 刷脏缓冲。
	return nil
}

func cloneSnapshot(snap *Snapshot) *Snapshot {
	if snap == nil {
		return &Snapshot{}
	}
	out := &Snapshot{Flags: make([]FlagSnap, 0, len(snap.Flags))}
	for _, f := range snap.Flags {
		rules := make([]RuleSnap, 0, len(f.Rules))
		for _, r := range f.Rules {
			rules = append(rules, RuleSnap{
				Attr:   r.Attr,
				Op:     r.Op,
				Values: append([]string(nil), r.Values...),
			})
		}
		out.Flags = append(out.Flags, FlagSnap{
			Key:         f.Key,
			Enabled:     f.Enabled,
			Percent:     f.Percent,
			Rules:       rules,
			Description: f.Description,
			UpdatedAt:   f.UpdatedAt,
		})
	}
	return out
}
