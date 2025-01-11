package icu4xgo

import (
	"testing"

	"golang.org/x/text/collate"

	"github.com/stretchr/testify/assert"
)

func TestCollator(t *testing.T) {
	t.Run("Compare in en-GB", func(t *testing.T) {
		c := NewCollator(NewLocale("en-GB"))
		assert.Equal(t, -1, c.Compare("abc", "abd"))
	})

	t.Run("Compare in sv", func(t *testing.T) {
		c := NewCollator(NewLocale("sv"))
		assert.Equal(t, -1, c.Compare("a", "ä"))
		assert.Equal(t, -1, c.Compare("z", "ä"))
	})

	t.Run("Compare in de ", func(t *testing.T) {
		c := NewCollator(NewLocale("de"))
		assert.Equal(t, -1, c.Compare("a", "ä"))
		assert.Equal(t, 1, c.Compare("z", "ä"))
		assert.Equal(t, -1, c.Compare("z", "Z"))
	})
	t.Run("Compare in de ignore case", func(t *testing.T) {
		c := NewCollatorWithOptions(NewLocale("de"), collate.IgnoreCase)
		assert.Equal(t, 0, c.Compare("z", "Z"))
		assert.Equal(t, 1, c.Compare("Z", "ä"))
	})
}
