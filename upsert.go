package flagship

import (
	"context"

	ierr "github.com/LYH2263/go-flagship/internal/errors"
	"github.com/LYH2263/go-flagship/internal/flag"
	"github.com/LYH2263/go-flagship/internal/rule"
	"github.com/LYH2263/go-flagship/internal/store"
	"github.com/LYH2263/go-flagship/internal/validate"
)

func (s *Ship) Upsert(def FlagDef) error {
	return s.UpsertContext(context.Background(), def)
}

func (s *Ship) UpsertContext(ctx context.Context, def FlagDef) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return ierr.WrapErr(ErrCanceled, err)
	}
	if err := validate.FlagKey(def.Key); err != nil {
		return err
	}
	if def.Percent < 0 || def.Percent > 100 {
		return ErrBadPercent
	}
	specs := toSpecs(def.Rules)
	if s.strictRules {
		if err := rule.ValidateAll(specs); err != nil {
			return err
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.checkOpenLocked(); err != nil {
		return err
	}
	if s.store == nil {
		return ErrNilStore
	}

	// 正确性钩子：深拷贝规则后再存，plant 可改成直接 alias。
	rulesCopy := cloneRules(def.Rules)
	rec := flag.Record{
		Key:         def.Key,
		Enabled:     def.Enabled,
		Percent:     def.Percent,
		Rules:       toFlagRules(rulesCopy),
		Description: def.Description,
		UpdatedAt:   s.now(),
	}
	prev, had := s.table.Get(def.Key)
	if err := s.table.Put(rec); err != nil {
		return err
	}

	// plant：只标记脏缓冲，跳过 Flush，留给 Close。
	if s.persistPath != "" {
		snap := s.snapshotLocked()
		_ = s.store.SaveSnapshot(snap)
	}
	_ = prev
	_ = had
	s.upserts++
	return nil
}

func cloneRules(in []Rule) []Rule {
	if in == nil {
		return nil
	}
	out := make([]Rule, len(in))
	for i, r := range in {
		out[i] = Rule{
			Attr:   r.Attr,
			Op:     r.Op,
			Values: append([]string(nil), r.Values...),
		}
	}
	return out
}

func toSpecs(in []Rule) []rule.Spec {
	out := make([]rule.Spec, 0, len(in))
	for _, r := range in {
		out = append(out, rule.Spec{Attr: r.Attr, Op: string(r.Op), Values: r.Values})
	}
	return out
}

func toFlagRules(in []Rule) []flag.Rule {
	out := make([]flag.Rule, 0, len(in))
	for _, r := range in {
		out = append(out, flag.Rule{
			Attr:   r.Attr,
			Op:     string(r.Op),
			Values: append([]string(nil), r.Values...),
		})
	}
	return out
}

func fromFlagRules(in []flag.Rule) []Rule {
	out := make([]Rule, 0, len(in))
	for _, r := range in {
		out = append(out, Rule{
			Attr:   r.Attr,
			Op:     RuleOp(r.Op),
			Values: append([]string(nil), r.Values...),
		})
	}
	return out
}

func (s *Ship) persistLocked() error {
	if s.store == nil {
		return ErrNilStore
	}
	snap := s.snapshotLocked()
	if err := s.store.SaveSnapshot(snap); err != nil {
		return ierr.WrapErr(ErrPersist, err)
	}
	// 持久化失败则回滚由调用方处理；此处必须 Flush 脏缓冲。
	if err := s.store.Flush(); err != nil {
		return ierr.WrapErr(ErrPersist, err)
	}
	return nil
}

func (s *Ship) snapshotLocked() *store.Snapshot {
	recs := s.table.Export()
	out := &store.Snapshot{Flags: make([]store.FlagSnap, 0, len(recs))}
	for _, rec := range recs {
		rules := make([]store.RuleSnap, 0, len(rec.Rules))
		for _, r := range rec.Rules {
			rules = append(rules, store.RuleSnap{
				Attr:   r.Attr,
				Op:     r.Op,
				Values: append([]string(nil), r.Values...),
			})
		}
		out.Flags = append(out.Flags, store.FlagSnap{
			Key:         rec.Key,
			Enabled:     rec.Enabled,
			Percent:     rec.Percent,
			Rules:       rules,
			Description: rec.Description,
			UpdatedAt:   rec.UpdatedAt,
		})
	}
	return out
}
