//go:build !darwin
// +build !darwin

package icu4xgo

func NewListFormatter(locale Locale, initType ListInitType, listLength ListLength) ListFormatter {
	return NewCListFormatter(locale.(*CLocale), initType, listLength)
}
