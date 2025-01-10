// go:build !darwin
//  +build !darwin

#include "./ig_locale.h"
#include "./ig_string.h"
#include <string.h>
#include <stdlib.h>

IGLocale *ig_init_locale(const char *localeText)
{
    IGLocale *l = malloc(sizeof(IGLocale));
    l->source = localeText;
    struct DiplomatStringView locale_text_str = ig_init_string(localeText);

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
    IGStringWriter *w = ig_init_string_writer();
    icu4x_Locale_language_mv1(l->locale, w->write);
    return ig_string_writer_to_string(w);
}

const char *ig_get_locale_basename(IGLocale *IGLocale)
{
    IGStringWriter *w = ig_init_string_writer();
    icu4x_Locale_basename_mv1(IGLocale->locale, w->write);
    return ig_string_writer_to_string(w);
}

const char *ig_get_locale_script(IGLocale *IGLocale)
{
    IGStringWriter *w = ig_init_string_writer();
    icu4x_Locale_script_mv1(IGLocale->locale, w->write);
    return ig_string_writer_to_string(w);
}

const char *ig_get_locale_region(IGLocale *IGLocale)
{
    IGStringWriter *w = ig_init_string_writer();
    icu4x_Locale_region_mv1(IGLocale->locale, w->write);
    return ig_string_writer_to_string(w);
}

const char *ig_get_locale_extension(IGLocale *IGLocale, const char *extension)
{
    IGStringWriter *w = ig_init_string_writer();
    struct DiplomatStringView extension_str = ig_init_string(extension);
    icu4x_Locale_get_unicode_extension_mv1_result result = icu4x_Locale_get_unicode_extension_mv1(IGLocale->locale, extension_str, w->write);
    if (!result.is_ok)
    {
        return NULL;
    }
    return ig_string_writer_to_string(w);
}

void ig_free_locale(IGLocale *l)
{
    icu4x_Locale_destroy_mv1(l->locale);
    free(l);
}