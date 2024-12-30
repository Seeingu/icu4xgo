package icu4xgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluralRules(t *testing.T) {
	t.Run("Cardinal in ar", func(t *testing.T) {
		l := NewLocale("ar")
		p := NewPluralRules(l, Cardinal)
		categories := p.Categories()
		assert.Equal(t, len(categories), 6)
		assert.Equal(t, categories[0], PluralCategoryZero)
		assert.Equal(t, categories[5], PluralCategoryOther)
	})

	t.Run("Ordinal", func(t *testing.T) {
		l := NewLocale("en-US")
		p := NewPluralRules(l, Ordinal)
		assert.Equal(t, p.Select(0), PluralCategoryOther)
		assert.Equal(t, p.Select(1), PluralCategoryOne)
		assert.Equal(t, p.Select(2), PluralCategoryTwo)
		assert.Equal(t, p.Select(21), PluralCategoryOne)
	})
}
