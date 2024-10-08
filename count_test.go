package fun_test

import (
	"slices"
	"testing"

	"github.com/martindrlik/fun"
)

func TestCount(t *testing.T) {
	count := fun.Count(slices.Values([]string{"foo", "bar", "baz"}))
	if count != 3 {
		t.Errorf("expected count to be 3, got %v", count)
	}
}
