package fun_test

import (
	"testing"

	"github.com/martindrlik/fun"
)

func TestDup(t *testing.T) {
	v := "hello"
	v1, v2 := fun.Dup(v)
	if v1 != v2 {
		t.Errorf("%v and %v should be of same value", v1, v2)
	}
	if v1 != v {
		t.Errorf("%v should be %v", v1, v)
	}
}
