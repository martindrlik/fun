package fun_test

import (
	"testing"

	"github.com/martindrlik/fun"
)

func TestTrue(t *testing.T) {
	if !fun.True(func() {}) {
		t.Error("expected True to return true")
	}
}
