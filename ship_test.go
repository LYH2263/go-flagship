package flagship_test

import (
	"testing"

	"github.com/LYH2263/go-flagship"
)

func TestUpsertEvaluateHappy(t *testing.T) {
	s := flagship.New()
	defer s.Close()
	err := s.Upsert(flagship.FlagDef{
		Key:     "dark_mode",
		Enabled: true,
		Percent: 100,
		Rules: []flagship.Rule{
			{Attr: "org", Op: flagship.OpEQ, Values: []string{"acme"}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	dec, err := s.Evaluate("dark_mode", flagship.EvalAttrs{"org": "acme", "user": "u1"})
	if err != nil {
		t.Fatal(err)
	}
	if !dec.On {
		t.Fatalf("want on, got %+v", dec)
	}
}
