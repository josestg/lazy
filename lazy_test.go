package lazy_test

import (
	"errors"
	"github.com/josestg/lazy"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("when no error", func(t *testing.T) {
		loader := lazy.New(func() (int, error) {
			return 123, nil
		})

		if loader.Loaded() {
			t.Errorf("got %t, want %t", loader.Loaded(), false)
		}

		got := loader.Value()
		if got != 123 {
			t.Errorf("got %d, want %d", got, 123)
		}

		if !loader.Loaded() {
			t.Errorf("got %t, want %t", loader.Loaded(), true)
		}

		err := loader.Error()
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
	})

	t.Run("when error", func(t *testing.T) {
		loader := lazy.New(func() (int, error) {
			return 123, errors.New("error")
		})

		if loader.Loaded() {
			t.Errorf("got %t, want %t", loader.Loaded(), false)
		}

		got := loader.Value()
		if got != 123 {
			t.Errorf("got %d, want %d", got, 123)
		}

		if !loader.Loaded() {
			t.Errorf("got %t, want %t", loader.Loaded(), true)
		}

		err := loader.Error()
		if err == nil {
			t.Errorf("got %v, want %v", err, errors.New("error"))
		}
	})
}
