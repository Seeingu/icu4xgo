//go:build !darwin
// +build !darwin

package icu4xgo

func NewPluralRules(l Locale, t PluralRulesType) PluralRules {
	return NewCPluralRules(l.(*CLocale), t)
}
