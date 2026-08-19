package flagship_test

import (
	"testing"

	"github.com/LYH2263/go-flagship"
)

func TestBug02_SnapshotListSharedRules(t *testing.T) {
	s := flagship.New()
	defer s.Close()
	if err := s.Upsert(flagship.FlagDef{
		Key: "feat", Enabled: true, Percent: 100,
		Rules: []flagship.Rule{{Attr: "org", Op: flagship.OpEQ, Values: []string{"acme"}}},
	}); err != nil {
		t.Fatal(err)
	}
	list, err := s.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 1 || len(list[0].Rules) == 0 {
		t.Fatalf("bad list: %+v", list)
	}
	list[0].Rules[0].Values[0] = "mutated"
	snap, err := s.Snapshot()
	if err != nil {
		t.Fatal(err)
	}
	if snap[0].Rules[0].Values[0] == "mutated" {
		t.Fatal("List/Snapshot returned shared rule slices")
	}
	got, err := s.Get("feat")
	if err != nil {
		t.Fatal(err)
	}
	if got.Rules[0].Values[0] == "mutated" {
		t.Fatal("List mutation polluted store")
	}
}
