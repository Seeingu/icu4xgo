//go:build !darwin
// +build !darwin

package icu4xgo

// TODO(build): signature update
func NewNumberFormatter(locale Locale) NumberFormatter {
	return func(locale Locale) NumberFormatter {
		return NewCNumberFormatter(locale.(*CLocale))
	}(locale)
}
