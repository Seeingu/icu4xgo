package icu4xgo

type CollatorMask uint

const (
	CollatorMaskDiacritic       CollatorMask = 128
	CollatorMaskCaseInsensitive CollatorMask = 1
)

type Collator interface {
	Compare(a, b string) int
}
