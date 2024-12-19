#include "./icugo.h"
#include <string.h>
#include <stdlib.h>

IGLocale *ig_init_locale(const char *localeText)
{
    IGLocale *l = malloc(sizeof(IGLocale));
    l->source = localeText;
    struct DiplomatStringView locale_text_str = {
        localeText,
        strlen(localeText)};

    icu4x_Locale_from_string_mv1_result locale_result = icu4x_Locale_from_string_mv1(locale_text_str);
    if (!locale_result.is_ok)
    {
        return l;
    }
    l->locale = locale_result.ok;
    return l;
}

const char *ig_get_locale_language(IGLocale *l)
{
    DiplomatWrite *write = diplomat_buffer_write_create(0);
    icu4x_Locale_language_mv1(l->locale, write);
    return diplomat_buffer_write_get_bytes(write);
}

void ig_free_locale(IGLocale *IGLocale)
{
    icu4x_Locale_destroy_mv1(IGLocale->locale);
    free(IGLocale);
}