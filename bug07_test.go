package flagship_test

import (
	"context"
	"errors"
	"testing"

	"github.com/LYH2263/go-flagship"
)

func TestBug07_EvaluateContextHonorsCancel(t *testing.T) {
	s := flagship.New()
	defer s.Close()
	if err := s.Upsert(flagship.FlagDef{Key: "x", Enabled: true, Percent: 100}); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := s.EvaluateContext(ctx, "x", flagship.EvalAttrs{"user": "u1"})
	if err == nil {
		t.Fatal("expected cancel error")
	}
	if !errors.Is(err, flagship.ErrCanceled) && !errors.Is(err, context.Canceled) {
		t.Fatalf("want canceled, got %v", err)
	}
}
