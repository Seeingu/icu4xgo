package icu4xgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollator(t *testing.T) {
	t.Run("Compare in en-GB", func(t *testing.T) {
		c := NewCollator(NewLocale("en-GB"))
		defer c.Free()
		assert.Equal(t, c.Compare("abc", "abd"), -1)
		assert.Equal(t, c.Compare("a", "ä"), 0)
		assert.Equal(t, c.Compare("z", "ä"), 1)
	})

	t.Run("Compare in sv", func(t *testing.T) {
		c := NewCollator(NewLocale("sv"))
		defer c.Free()
		assert.Equal(t, c.Compare("a", "ä"), -1)
		assert.Equal(t, c.Compare("z", "ä"), -1)
	})
}
