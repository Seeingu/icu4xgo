package icu4xgo

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocale(t *testing.T) {
	// TODO(build): use NewLocale when darwin migration is complete
	var newLocale func(id string) Locale
	if runtime.GOOS == "darwin" {
		newLocale = NewLocaleMac
	} else {
		newLocale = func(id string) Locale {
			return NewCLocale(id)
		}
	}
	t.Run("Locale", func(t *testing.T) {
		locale := newLocale("zh-Hans-CN-u-ca-chinese-hc-h12")
		assert.Equal(t, "zh", locale.Language())
		assert.Equal(t, "Hans", locale.Script())
		assert.Equal(t, "CN", locale.Region())
		assert.Equal(t, "zh-Hans-CN", locale.BaseName())
		assert.Equal(t, "h12", locale.HourCycle())
		assert.Equal(t, "chinese", locale.Calendar())
	})
	t.Run("NumberingSystem", func(t *testing.T) {
		locale := newLocale("zh-Hans-CN-u-nu-hanidec")
		assert.Equal(t, "hanidec", locale.NumberingSystem())
	})
	t.Run("Numeric", func(t *testing.T) {
		locale := newLocale("zh-Hans-CN-u-nu-latn")
		assert.Equal(t, "latn", locale.NumberingSystem())
	})
	t.Run("Collation", func(t *testing.T) {
		locale := newLocale("zh-Hans-CN-u-co-pinyin")
		assert.Equal(t, "pinyin", locale.Collation())
	})
	t.Run("CaseFirst", func(t *testing.T) {
		locale := newLocale("zh-Hans-CN-u-kf-upper")
		assert.Equal(t, "upper", locale.CaseFirst())
	})
}
