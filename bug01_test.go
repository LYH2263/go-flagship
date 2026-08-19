package flagship_test

import (
	"testing"

	"github.com/LYH2263/go-flagship"
)

func TestBug01_UpsertRulesAttrsAlias(t *testing.T) {
	s := flagship.New()
	defer s.Close()
	vals := []string{"acme"}
	rules := []flagship.Rule{{Attr: "org", Op: flagship.OpEQ, Values: vals}}
	if err := s.Upsert(flagship.FlagDef{Key: "dark_mode", Enabled: true, Percent: 100, Rules: rules}); err != nil {
		t.Fatal(err)
	}
	vals[0] = "hacked"
	rules[0].Values[0] = "hacked"
	got, err := s.Get("dark_mode")
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Rules) == 0 || got.Rules[0].Values[0] == "hacked" {
		t.Fatalf("Upsert stored alias of caller rules/attrs: %+v", got.Rules)
	}
}
