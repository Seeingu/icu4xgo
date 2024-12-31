
#ifndef ICUGO_PLURAL_RULES_H
#define ICUGO_PLURAL_RULES_H

#include "./ig_locale.h"
#include "./ig_string.h"
#include "./icu4x/PluralRules.h"
#include "./icu4x/PluralOperands.h"
#include "./icu4x/DataProvider.h"
#include "./icu4x/FixedDecimal.h"

typedef struct IGPluralRules
{
    PluralRules *plural_rules;
    PluralOperands *operands;
} IGPluralRules;

IGPluralRules *ig_init_cardinal_plural_rules(IGLocale *locale);
IGPluralRules *ig_init_ordinal_plural_rules(IGLocale *locale);

PluralCategory ig_plural_rules_get_category(IGPluralRules *pl, const char *s);

PluralCategory ig_plural_rules_get_category_int(IGPluralRules *pl, int n);

PluralCategories ig_plural_rules_get_categories(IGPluralRules *pl);

void ig_free_plural_rules(IGPluralRules *pluralRules);

#endif // ICUGO_PLURAL_RULES_H