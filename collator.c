#include "./ig_collator.h"
#include <string.h>
#include <stdlib.h>

CollatorOptionsV1 ig_options_to_icu4x(IGCollatorOptions igOptions)
{
    CollatorOptionsV1 options = {
        .strength = {
            .is_ok = true,
            .ok = igOptions.strength},
        .alternate_handling = {.is_ok = true, .ok = igOptions.alternate_handling},
        .case_first = {.is_ok = true, .ok = igOptions.case_first},
        .max_variable = {.is_ok = true, .ok = igOptions.max_variable},
        .case_level = {.is_ok = true, .ok = igOptions.case_level},
        .numeric = {.is_ok = true, .ok = igOptions.numeric},
        .backward_second_level = {.is_ok = true, .ok = igOptions.backward_second_level},
    };
    return options;
}

IGCollator *ig_init_collator(IGLocale *locale, IGCollatorOptions igOptions)
{
    IGCollator *c = malloc(sizeof(IGCollator));
    DataProvider *dataProvider = icu4x_DataProvider_compiled_mv1();
    CollatorOptionsV1 options = ig_options_to_icu4x(igOptions);
    icu4x_Collator_create_v1_mv1_result result = icu4x_Collator_create_v1_mv1(dataProvider, locale->locale, options);
    if (!result.is_ok)
    {
        return c;
    }
    c->collator = result.ok;
    return c;
}

void ig_free_collator(IGCollator *collator)
{
    icu4x_Collator_destroy_mv1(collator->collator);
    free(collator);
}

int8_t ig_collator_compare(IGCollator *c, const char *left, const char *right)
{
    DiplomatStringView left_str = {
        left,
        strlen(left)};
    DiplomatStringView right_str = {
        right,
        strlen(right)};
    return icu4x_Collator_compare_utf8_mv1(c->collator, left_str, right_str);
}