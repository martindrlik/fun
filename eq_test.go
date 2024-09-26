package fun_test

import (
	"maps"
	"testing"

	"github.com/martindrlik/fun"
)

func TestEqAlwaysTrue(t *testing.T) {
	t.Run("true for any argument", func(t *testing.T) {
		if !fun.EqAlwaysTrue(false, true) {
			t.Error("should always return true")
		}
		if !fun.EqAlwaysTrue(3.14, "pi") {
			t.Error("should always return true")
		}
	})
	t.Run("used by maps.EqualFunc", func(t *testing.T) {
		foo := map[string]int{"hello": 1}
		bar := map[string]bool{"hello": true}
		if !maps.EqualFunc(foo, bar, fun.EqAlwaysTrue) {
			t.Errorf("expected map %+v to be equal to %+v as its value is ignored", foo, bar)
		}
	})
}

func TestEq(t *testing.T) {
	if !t.Run("test dup", TestDup) {
		t.Fatalf("required dup does not work as expected")
	}
	if !fun.Eq(fun.Dup(1)) {
		t.Error("expected 1 and 1 to be equal")
	}
	if !fun.Eq(fun.Dup("hello")) {
		t.Error(`expected "hello" and "hello" to be equal`)
	}
	if fun.Eq(2, 1) {
		t.Error("2 is not equal to 1")
	}
	if fun.Eq[any](1, "1") {
		t.Error(`1 is not equal to "1"`)
	}
}
