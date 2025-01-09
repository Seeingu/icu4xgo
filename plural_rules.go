package icu4xgo

type PluralRules interface {
	Select(n int) PluralCategory
	Categories() []PluralCategory
}

type PluralRulesType int

const (
	Cardinal PluralRulesType = iota
	Ordinal
)

type PluralCategory int

const (
	PluralCategoryZero PluralCategory = iota
	PluralCategoryOne
	PluralCategoryTwo
	PluralCategoryFew
	PluralCategoryMany
	PluralCategoryOther
)
