package icu4xgo

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

type CollatorDarwin struct {
	*CollatorGeneral
}

var _ Collator = (*CollatorDarwin)(nil)

func NewCollator(l Locale) Collator {
	c := &CollatorDarwin{
		CollatorGeneral: NewGeneralCollator(l),
	}
	return c
}

func (c *CollatorDarwin) compareOptionsLocale(a, b string, mask foundation.StringCompareOptions, locale objc.IObject) int {
	var s foundation.String
	other := b
	length := len(a)
	isNegative := false
	if len(a) < len(b) {
		s = foundation.NewStringWithString(b)
		isNegative = true
		length = len(b)
		other = a
	} else {
		s = foundation.NewStringWithString(a)
	}
	r := foundation.Range{
		Location: 0,
		Length:   uint64(length),
	}
	rv := objc.Call[foundation.ComparisonResult](s, objc.Sel("compare:options:range:locale:"), other, foundation.DiacriticInsensitiveSearch, r, locale)
	if isNegative {
		return -int(rv)
	}
	return int(rv)
}

func (c *CollatorDarwin) CompareMask(a, b string, mask CollatorMask) int {
	return c.compareOptionsLocale(a, b,
		foundation.StringCompareOptions(mask),
		c.locale.(*LocaleDarwin).ptr)
}
