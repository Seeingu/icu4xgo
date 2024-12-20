package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocle(t *testing.T) {
	t.Run("Locale", func(t *testing.T) {
		locale := NewLocale("zh-Hans-CN-u-ca-chinese-hc-h12")
		defer locale.Free()
		assert.Equal(t, "zh", locale.Language())
		assert.Equal(t, "Hans", locale.Script())
		assert.Equal(t, "CN", locale.Region())
		assert.Equal(t, "zh-Hans-CN", locale.BaseName())
		hourCycle, _ := locale.HourCycle()
		assert.Equal(t, "h12", hourCycle)
		ca, _ := locale.Calendar()
		assert.Equal(t, "chinese", ca)
	})
	t.Run("NumberingSystem", func(t *testing.T) {
		locale := NewLocale("zh-Hans-CN-u-nu-hanidec")
		defer locale.Free()
		n, _ := locale.NumberingSystem()
		assert.Equal(t, "hanidec", n)
	})
	t.Run("Numeric", func(t *testing.T) {
		locale := NewLocale("zh-Hans-CN-u-nu-latn")
		defer locale.Free()
		n, _ := locale.NumberingSystem()
		assert.Equal(t, "latn", n)
	})
	t.Run("Collation", func(t *testing.T) {
		locale := NewLocale("zh-Hans-CN-u-co-pinyin")
		defer locale.Free()
		c, _ := locale.Collation()
		assert.Equal(t, "pinyin", c)
	})
	t.Run("CaseFirst", func(t *testing.T) {
		locale := NewLocale("zh-Hans-CN-u-kf-upper")
		defer locale.Free()
		cf, _ := locale.CaseFirst()
		assert.Equal(t, "upper", cf)
	})
}
