package flag

import (
	"sort"

	ierr "github.com/LYH2263/go-flagship/internal/errors"
)

type Table struct {
	max int
	by  map[string]Record
}

func NewTable(max int) *Table {
	if max < 1 {
		max = 1
	}
	return &Table{max: max, by: make(map[string]Record)}
}

func (t *Table) Len() int { return len(t.by) }

func (t *Table) Get(key string) (Record, bool) {
	r, ok := t.by[key]
	return r, ok
}

func (t *Table) Put(rec Record) error {
	if _, ok := t.by[rec.Key]; !ok && len(t.by) >= t.max {
		return ierr.ErrTooMany
	}
	t.by[rec.Key] = cloneRecord(rec)
	return nil
}

func (t *Table) Delete(key string) bool {
	if _, ok := t.by[key]; !ok {
		return false
	}
	delete(t.by, key)
	return true
}

func (t *Table) Export() []Record {
	keys := make([]string, 0, len(t.by))
	for k := range t.by {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]Record, 0, len(keys))
	for _, k := range keys {
		out = append(out, t.by[k])
	}
	return out
}

func cloneRecord(r Record) Record {
	rules := make([]Rule, len(r.Rules))
	for i, rule := range r.Rules {
		vals := append([]string(nil), rule.Values...)
		rules[i] = Rule{Attr: rule.Attr, Op: rule.Op, Values: vals}
	}
	return Record{
		Key:         r.Key,
		Enabled:     r.Enabled,
		Percent:     r.Percent,
		Rules:       rules,
		Description: r.Description,
		UpdatedAt:   r.UpdatedAt,
	}
}
