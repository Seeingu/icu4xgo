package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"
import "runtime"

type PluralRules struct {
	ptr *C.IGPluralRules
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

func NewPluralRules(l *Locale, t PluralRulesType) *PluralRules {
	var p *PluralRules
	if t == Cardinal {
		p = &PluralRules{
			ptr: C.ig_init_cardinal_plural_rules(l.ptr),
		}
	} else {
		p = &PluralRules{
			ptr: C.ig_init_ordinal_plural_rules(l.ptr),
		}
	}
	runtime.SetFinalizer(p, func(p *PluralRules) {
		p.free()
	})
	return p
}

func (p *PluralRules) free() {
	C.ig_free_plural_rules(p.ptr)
}

func (p *PluralRules) Select(n int) PluralCategory {
	category := C.ig_plural_rules_get_category_int(p.ptr, C.int(n))
	return PluralCategory(category)
}

func (p *PluralRules) Categories() []PluralCategory {
	var categories []PluralCategory
	cc := C.ig_plural_rules_get_categories(p.ptr)
	if cc.zero {
		categories = append(categories, PluralCategoryZero)
	}
	if cc.one {
		categories = append(categories, PluralCategoryOne)
	}
	if cc.two {
		categories = append(categories, PluralCategoryTwo)
	}
	if cc.few {
		categories = append(categories, PluralCategoryFew)
	}
	if cc.many {
		categories = append(categories, PluralCategoryMany)
	}
	if cc.other {
		categories = append(categories, PluralCategoryOther)
	}
	return categories
}
