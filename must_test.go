package fun_test

import (
	"errors"
	"testing"

	"github.com/martindrlik/fun"
)

func TestMust(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		v := fun.Must(func() (int, error) {
			return 1, nil
		}())
		if v != 1 {
			t.Errorf("expected 1, got %v", v)
		}
	})
	t.Run("has error", func(t *testing.T) {
		defer func() {
			p := recover().(error)
			if p.Error() != "some error" {
				t.Errorf("expected panic with error \"some error\", got %v", p)
			}
		}()
		fun.Must(func() (int, error) {
			return 0, errors.New("some error")
		}())
	})
}
