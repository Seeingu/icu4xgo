//go:build !darwin
// +build !darwin

package icu4xgo

func NewCollator(l Locale) Collator {
	return NewCCollator(l.(*CLocale))
}
