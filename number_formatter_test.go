package icu4xgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberFormatter(t *testing.T) {
	t.Run("NumberFormat in en-US", func(t *testing.T) {
		locale := NewLocale("en-US")
		nf := NewNumberFormatter(locale)
		f, err := nf.Format("123456.789")
		assert.NoError(t, err)
		assert.Equal(t, "123,456.789", f)
	})
	t.Run("NumberFormat in de-DE", func(t *testing.T) {
		locale := NewLocale("de-DE")
		nf := NewNumberFormatter(locale)
		f, err := nf.Format("123456.789")
		assert.NoError(t, err)
		assert.Equal(t, "123.456,789", f)
	})
	t.Run("NumberFormat in ar-EG", func(t *testing.T) {
		locale := NewLocale("ar-EG")
		nf := NewNumberFormatter(locale)
		f, err := nf.Format("123456.789")
		assert.NoError(t, err)
		assert.Equal(t, "١٢٣٬٤٥٦٫٧٨٩", f)
	})
}
