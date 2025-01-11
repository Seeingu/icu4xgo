//go:build !darwin
// +build !darwin

package icu4xgo

// TODO(xxx): CompareMask
func NewCollator(l Locale) Collator {
	return NewGeneralCollator(l)
}
