package flagship_test

import (
	"testing"

	"github.com/LYH2263/go-flagship"
)

func TestBug04_NilHasherRolloutNoPanic(t *testing.T) {
	s := flagship.New(flagship.WithHasher(nil), flagship.WithRollout(nil))
	defer s.Close()
	if err := s.Upsert(flagship.FlagDef{Key: "x", Enabled: true, Percent: 100}); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if rec := recover(); rec != nil {
			t.Fatalf("Evaluate panicked on nil hasher/rollout: %v", rec)
		}
	}()
	dec, err := s.Evaluate("x", flagship.EvalAttrs{"user": "u1"})
	if err != nil {
		t.Fatal(err)
	}
	if !dec.On {
		t.Fatalf("expected on with defaults, got %+v", dec)
	}
}
