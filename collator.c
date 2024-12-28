#include "./ig_collator.h"
#include <string.h>
#include <stdlib.h>

IGCollator *ig_init_collator(IGLocale *locale)
{
    IGCollator *c = malloc(sizeof(IGCollator));
    CollatorOptionsV1 options = {
        .strength = CollatorStrength_Secondary,
        .alternate_handling = CollatorAlternateHandling_NonIgnorable,
        .case_first = CollatorCaseFirst_Off,
        .max_variable = CollatorMaxVariable_Punctuation,
        .case_level = CollatorCaseLevel_Off,
        .numeric = CollatorNumeric_Off,
        .backward_second_level = CollatorBackwardSecondLevel_Off};

    DataProvider *dataProvider = icu4x_DataProvider_compiled_mv1();
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