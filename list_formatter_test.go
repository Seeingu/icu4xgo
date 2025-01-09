package icu4xgo

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListFormatter(t *testing.T) {
	t.Run("ListFormat and wide in en-US", func(t *testing.T) {
		locale := NewLocale("en-US")
		lf := NewListFormatter(locale, ListAnd, ListLengthWide)
		f := lf.Format([]string{"apple", "banana", "cherry", "strawberry"})
		assert.Equal(t, "apple, banana, cherry, and strawberry", f)
		f = lf.Format([]string{"apple", "banana"})
		assert.Equal(t, "apple and banana", f)
		f = lf.Format([]string{"apple"})
		assert.Equal(t, "apple", f)
		f = lf.Format([]string{})
		assert.Equal(t, "", f)
	})
	t.Run("ListFormat and long in de", func(t *testing.T) {
		locale := NewLocale("de")
		lf := NewListFormatter(locale, ListAnd, ListLengthWide)
		f := lf.Format([]string{"Apfel", "Banane"})
		assert.Equal(t, "Apfel und Banane", f)
	})
	// FIXME: only support join by and
	if runtime.GOOS == "darwin" {
		return
	}
	t.Run("ListFormat or narrow in en-US", func(t *testing.T) {
		locale := NewLocale("en-US")
		lf := NewListFormatter(locale, ListOr, ListLengthNarrow)
		f := lf.Format([]string{"apple", "banana", "strawberry"})
		assert.Equal(t, "apple, banana, or strawberry", f)
		f = lf.Format([]string{"apple", "banana"})
		assert.Equal(t, "apple or banana", f)
	})
	t.Run("ListFormat unit short in en-US", func(t *testing.T) {
		locale := NewLocale("en-US")
		lf := NewListFormatter(locale, ListUnit, ListLengthShort)
		f := lf.Format([]string{"apple", "banana"})
		assert.Equal(t, "apple, banana", f)
	})
}
