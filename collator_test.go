package icu4xgo

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollator(t *testing.T) {
	t.Run("Compare in en-GB", func(t *testing.T) {
		c := NewCollator(NewLocale("en-GB"))
		assert.Equal(t, -1, c.Compare("abc", "abd"))
		// TODO(xxx): non universal comparison
		if runtime.GOOS == "darwin" {
			assert.Equal(t, -1, c.Compare("a", "ä"))
		} else {
			assert.Equal(t, 0, c.Compare("a", "ä"))
		}
		assert.Equal(t, 1, c.Compare("z", "ä"))
	})

	t.Run("Compare in sv", func(t *testing.T) {
		c := NewCollator(NewLocale("sv"))
		assert.Equal(t, -1, c.Compare("a", "ä"))
		assert.Equal(t, -1, c.Compare("z", "ä"))
	})
}
