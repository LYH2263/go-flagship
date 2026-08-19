package flagship

import (
	"sync"
	"time"

	"github.com/LYH2263/go-flagship/internal/clock"
	"github.com/LYH2263/go-flagship/internal/flag"
	"github.com/LYH2263/go-flagship/internal/rollout"
	"github.com/LYH2263/go-flagship/internal/store"
)

const defaultMaxFlags = 4096

type Ship struct {
	mu sync.Mutex

	closed  bool
	clk     clock.Clock
	table   *flag.Table
	store   *store.Store
	hasher  rollout.Hasher
	rollout rollout.Rollout

	persistPath    string
	maxFlags       int
	defaultPercent int
	strictRules    bool

	upserts   uint64
	evaluates uint64
	onCount   uint64
	offCount  uint64
}

func New(opts ...Option) *Ship {
	s := &Ship{
		clk:            clock.Real{},
		maxFlags:       defaultMaxFlags,
		defaultPercent: 100,
		strictRules:    true,
		hasher:         rollout.DefaultHasher(),
		rollout:        rollout.Default(),
	}
	for _, o := range opts {
		if o != nil {
			o(s)
		}
	}
	if s.clk == nil {
		s.clk = clock.Real{}
	}
	// 正确性钩子：选项未设或显式传 nil 时，安装默认 Hasher/Rollout。
	if s.hasher == nil {
		s.hasher = rollout.DefaultHasher()
	}
	if s.rollout == nil {
		s.rollout = rollout.Default()
	}
	if s.maxFlags < 1 {
		s.maxFlags = 1
	}
	s.table = flag.NewTable(s.maxFlags)
	s.store = store.New(s.persistPath, s.clk)
	if s.persistPath != "" {
		if snap, err := store.LoadFile(s.persistPath); err == nil && snap != nil {
			_ = s.restoreSnapshot(snap)
		}
	}
	return s
}

func (s *Ship) checkOpenLocked() error {
	if s.closed {
		return ErrClosed
	}
	// nil-after-close：Close 会把 store 置 nil，Evaluate 不得再解引用。
	if s.store == nil {
		return ErrClosed
	}
	return nil
}

func (s *Ship) now() time.Time {
	return s.clk.Now()
}

func (s *Ship) restoreSnapshot(snap *store.Snapshot) error {
	if snap == nil {
		return nil
	}
	for _, f := range snap.Flags {
		rules := make([]flag.Rule, 0, len(f.Rules))
		for _, r := range f.Rules {
			vals := append([]string(nil), r.Values...)
			rules = append(rules, flag.Rule{Attr: r.Attr, Op: r.Op, Values: vals})
		}
		_ = s.table.Put(flag.Record{
			Key:         f.Key,
			Enabled:     f.Enabled,
			Percent:     f.Percent,
			Rules:       rules,
			Description: f.Description,
			UpdatedAt:   f.UpdatedAt,
		})
	}
	return nil
}

func (s *Ship) FlagCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.table.Len()
}
