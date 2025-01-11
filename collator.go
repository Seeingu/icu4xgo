package icu4xgo

import "golang.org/x/text/collate"

type CollatorMask uint

const (
	CollatorMaskDiacritic       CollatorMask = 128
	CollatorMaskCaseInsensitive CollatorMask = 1
)

type Collator interface {
	Compare(a, b string) int
}

type CollatorGeneral struct {
	Collator
	locale  Locale
	collate *collate.Collator
}

var _ Collator = (*CollatorGeneral)(nil)

func NewCollator(l Locale) *CollatorGeneral {
	cl := collate.New(l.ToGoLanguage())
	c := &CollatorGeneral{
		locale:  l,
		collate: cl,
	}
	return c
}

func NewCollatorWithOptions(l Locale, option collate.Option) *CollatorGeneral {
	cl := collate.New(l.ToGoLanguage(), option)
	c := &CollatorGeneral{
		locale:  l,
		collate: cl,
	}
	return c
}

func (c *CollatorGeneral) Compare(a, b string) int {
	return c.collate.Compare([]byte(a), []byte(b))
}
