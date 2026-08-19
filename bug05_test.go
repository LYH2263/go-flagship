package flagship_test

import (
	"errors"
	"testing"

	"github.com/LYH2263/go-flagship"
)

func TestBug05_InvalidRuleWrapsSentinel(t *testing.T) {
	s := flagship.New(flagship.WithStrictRules(true))
	defer s.Close()
	err := s.Upsert(flagship.FlagDef{
		Key:     "dark",
		Enabled: true,
		Percent: 10,
		Rules:   []flagship.Rule{{Attr: "", Op: flagship.OpEQ, Values: []string{"x"}}},
	})
	if err == nil {
		t.Fatal("expected invalid rule error")
	}
	if !errors.Is(err, flagship.ErrInvalidRule) {
		t.Fatalf("Upsert must errors.Is ErrInvalidRule, got %v", err)
	}
}
