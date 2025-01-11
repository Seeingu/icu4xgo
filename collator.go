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

// CollatorGeneral is a general collator support compare
// strings without options
type CollatorGeneral struct {
	Collator
	locale  Locale
	collate *collate.Collator
}

var _ Collator = (*CollatorGeneral)(nil)

func NewGeneralCollator(l Locale) *CollatorGeneral {
	cl := collate.New(l.ToGoLanguage())
	c := &CollatorGeneral{
		locale:  l,
		collate: cl,
	}
	return c
}

func (c *CollatorGeneral) Compare(a, b string) int {
	return c.collate.Compare([]byte(a), []byte(b))
}
