#include "./ig_fixed_decimal_formatter.h"

IGFixedDecimalFormatter *ig_init_fixed_decimal_formatter(
    const IGLocale *locale,
    FixedDecimalGroupingStrategy grouping_strategy)
{
    DataProvider *dataProvider = icu4x_DataProvider_compiled_mv1();
    IGFixedDecimalFormatter *formatter = malloc(sizeof(IGFixedDecimalFormatter));
    FixedDecimalGroupingStrategy_option strategyOptions = {.ok = grouping_strategy, .is_ok = true};
    icu4x_FixedDecimalFormatter_create_with_grouping_strategy_mv1_result result =
        icu4x_FixedDecimalFormatter_create_with_grouping_strategy_mv1(dataProvider, locale->locale, strategyOptions);
    formatter->is_ok = result.is_ok;
    if (!formatter->is_ok)
    {
        return formatter;
    }
    formatter->formatter = result.ok;
    return formatter;
}

// TODO: different types of fixed decimal formatting
const char *ig_fixed_decimal_format(IGFixedDecimalFormatter *formatter, const char *s)
{
    DiplomatStringView input = ig_init_string(s);
    icu4x_FixedDecimal_from_string_mv1_result result = icu4x_FixedDecimal_from_string_mv1(input);
    if (!result.is_ok)
    {
        return "";
    }
    // TODO: error check
    FixedDecimal *d = result.ok;
    IGStringWriter *w = ig_init_string_writer();
    icu4x_FixedDecimalFormatter_format_mv1(formatter->formatter, d, w->write);
    return ig_string_writer_to_string(w);
}

void ig_free_fixed_decimal_formatter(IGFixedDecimalFormatter *formatter)
{
    icu4x_FixedDecimalFormatter_destroy_mv1(formatter->formatter);
    free(formatter);
}