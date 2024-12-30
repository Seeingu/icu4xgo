#ifndef ICUGO_LIST_FORMATTER_H
#define ICUGO_LIST_FORMATTER_H

#include "./ig_locale.h"
#include "./ig_string.h"
#include "./icu4x/DataProvider.h"
#include "./icu4x/ListFormatter.h"

typedef struct IGListFormatter
{
    ListFormatter *formatter;
    bool is_ok;
} IGListFormatter;

typedef enum IGListInitType
{
    AND = 0,
    OR = 1,
    UNIT = 2,
} IGListInitType;

IGListFormatter *ig_init_list_formatter(
    const IGLocale *locale,
    IGListInitType init_type,
    ListLength length);

const char *ig_list_format(
    IGListFormatter *formatter,
    DiplomatStringsView list);

void ig_free_list_formatter(IGListFormatter *formatter);

#endif // ICUGO_LIST_FORMATTER_H