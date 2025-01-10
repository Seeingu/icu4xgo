// go:build !darwin
//  +build !darwin

#include "./ig_plural_rules.h"
#include <string.h>
#include <stdlib.h>

IGPluralRules *ig_init_cardinal_plural_rules(IGLocale *locale)
{
    IGPluralRules *pr = malloc(sizeof(IGPluralRules));
    icu4x_PluralRules_create_cardinal_mv1_result result = icu4x_PluralRules_create_cardinal_mv1(
        icu4x_DataProvider_compiled_mv1(),
        locale->locale);
    if (!result.is_ok)
    {
        return pr;
    }
    pr->plural_rules = result.ok;
    return pr;
}

IGPluralRules *ig_init_ordinal_plural_rules(IGLocale *locale)
{
    IGPluralRules *pr = malloc(sizeof(IGPluralRules));
    icu4x_PluralRules_create_ordinal_mv1_result result = icu4x_PluralRules_create_ordinal_mv1(
        icu4x_DataProvider_compiled_mv1(), locale->locale);
    if (!result.is_ok)
    {
        return pr;
    }
    pr->plural_rules = result.ok;
    return pr;
}

void ig_init_operands(IGPluralRules *pl, const char *s)
{
    struct DiplomatStringView s_str = ig_init_string(s);
    icu4x_PluralOperands_from_string_mv1_result result = icu4x_PluralOperands_from_string_mv1(s_str);
    if (result.is_ok)
    {
        pl->operands = result.ok;
    }
}

void ig_init_operands_int(IGPluralRules *pl, int n)
{
    const FixedDecimal *d = icu4x_FixedDecimal_from_int32_mv1(n);
    pl->operands = icu4x_PluralOperands_from_fixed_decimal_mv1(d);
}

PluralCategories ig_plural_rules_get_categories(IGPluralRules *pl)
{
    return icu4x_PluralRules_categories_mv1(pl->plural_rules);
}

PluralCategory ig_plural_rules_get_category(IGPluralRules *pl, const char *s)
{
    ig_init_operands(pl, s);
    return icu4x_PluralRules_category_for_mv1(pl->plural_rules, pl->operands);
}

PluralCategory ig_plural_rules_get_category_int(IGPluralRules *pl, int n)
{
    ig_init_operands_int(pl, n);
    return icu4x_PluralRules_category_for_mv1(pl->plural_rules, pl->operands);
}

void ig_free_plural_rules(IGPluralRules *pluralRules)
{
    icu4x_PluralRules_destroy_mv1(pluralRules->plural_rules);
    free(pluralRules);
}
