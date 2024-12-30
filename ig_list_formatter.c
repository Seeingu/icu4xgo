#include "./ig_list_formatter.h"
#include <stdlib.h>

IGListFormatter *ig_init_list_formatter(
    const IGLocale *locale,
    IGListInitType init_type,
    ListLength length)
{
    DataProvider *dataProvider = icu4x_DataProvider_compiled_mv1();
    IGListFormatter *formatter = malloc(sizeof(IGListFormatter));
    ListFormatter *lf;

    if (init_type == AND)
    {
        icu4x_ListFormatter_create_and_with_length_mv1_result result =
            icu4x_ListFormatter_create_and_with_length_mv1(dataProvider, locale->locale, length);
        formatter->is_ok = result.is_ok;
        lf = result.ok;
    }
    else if (init_type == OR)
    {
        icu4x_ListFormatter_create_or_with_length_mv1_result result =
            icu4x_ListFormatter_create_or_with_length_mv1(dataProvider, locale->locale, length);
        formatter->is_ok = result.is_ok;
        lf = result.ok;
    }
    else if (init_type == UNIT)
    {
        icu4x_ListFormatter_create_unit_with_length_mv1_result result =
            icu4x_ListFormatter_create_unit_with_length_mv1(dataProvider, locale->locale, length);
        formatter->is_ok = result.is_ok;
        lf = result.ok;
    }
    else
    {
        // ignore
    }
    if (!formatter->is_ok)
    {
        return formatter;
    }
    formatter->formatter = lf;
    return formatter;
}

const char *ig_list_format(
    IGListFormatter *formatter,
    DiplomatStringsView list)
{
    IGStringWriter *w = ig_init_string_writer();
    icu4x_ListFormatter_format_utf8_mv1(formatter->formatter, list, w->write);
    return ig_string_writer_to_string(w);
}

void ig_free_list_formatter(IGListFormatter *formatter)
{
    icu4x_ListFormatter_destroy_mv1(formatter->formatter);
    free(formatter);
}
