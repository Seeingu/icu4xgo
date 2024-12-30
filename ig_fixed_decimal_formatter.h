
#ifndef ICUGO_FIXED_DECIMAL_FORMATTER_H
#define ICUGO_FIXED_DECIMAL_FORMATTER_H

#include "./ig_locale.h"
#include "./ig_string.h"
#include "./icu4x/FixedDecimal.h"
#include "./icu4x/FixedDecimalFormatter.h"
#include "./icu4x/DataProvider.h"

#include <stdlib.h>
#include <string.h>

typedef struct IGFixedDecimalFormatter
{
    FixedDecimalFormatter *formatter;
    bool is_ok;
} IGFixedDecimalFormatter;

IGFixedDecimalFormatter *ig_init_fixed_decimal_formatter(
    const IGLocale *locale,
    FixedDecimalGroupingStrategy grouping_strategy);

const char *ig_fixed_decimal_format(IGFixedDecimalFormatter *formatter, const char *s);

void ig_free_fixed_decimal_formatter(IGFixedDecimalFormatter *formatter);

#endif // ICUGO_FIXED_DECIMAL_FORMATTER_H