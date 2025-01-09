//go:build !darwin
// +build !darwin

package icu4xgo

// TODO(build): signature update
func NewLocale(id string) Locale {
	return func(id string) Locale {
		return NewCLocale(id)
	}(id)
}
