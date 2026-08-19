package flagship_test

import (
	"context"
	"errors"
	"testing"

	"github.com/LYH2263/go-flagship"
)

func TestBug08_EvaluateAllHonorsContext(t *testing.T) {
	s := flagship.New()
	defer s.Close()
	for _, k := range []string{"a", "b", "c"} {
		if err := s.Upsert(flagship.FlagDef{Key: k, Enabled: true, Percent: 100}); err != nil {
			t.Fatal(err)
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := s.EvaluateAll(ctx, []string{"a", "b", "c"}, flagship.EvalAttrs{"user": "u1"})
	if err == nil {
		t.Fatal("EvaluateAll ignored canceled ctx")
	}
	if !errors.Is(err, flagship.ErrCanceled) && !errors.Is(err, context.Canceled) {
		t.Fatalf("want canceled, got %v", err)
	}
}
