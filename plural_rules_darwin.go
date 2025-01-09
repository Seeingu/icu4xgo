package icu4xgo

type PluralRulesDarwin struct {
	PluralRules
	locale Locale
}

var _ PluralRules = (*PluralRulesDarwin)(nil)

func NewPluralRules(l Locale, t PluralRulesType) PluralRules {
	p := &PluralRulesDarwin{
		locale: l,
	}
	return p
}

// TODO(darwin): Implement plural rules for Darwin
func (p *PluralRulesDarwin) Categories() []PluralCategory {
	return []PluralCategory{}
}

// TODO(darwin): Implement plural rules for Darwin
func (p *PluralRulesDarwin) Select(n int) PluralCategory {
	return PluralCategoryOther
}
