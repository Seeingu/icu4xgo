#include "./locale.h"
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

const char *ig_get_locale_basename(IGLocale *IGLocale)
{
    DiplomatWrite *write = diplomat_buffer_write_create(0);
    icu4x_Locale_basename_mv1(IGLocale->locale, write);
    return diplomat_buffer_write_get_bytes(write);
}

const char *ig_get_locale_script(IGLocale *IGLocale)
{
    DiplomatWrite *write = diplomat_buffer_write_create(0);
    icu4x_Locale_script_mv1(IGLocale->locale, write);
    return diplomat_buffer_write_get_bytes(write);
}

const char *ig_get_locale_region(IGLocale *IGLocale)
{
    DiplomatWrite *write = diplomat_buffer_write_create(0);
    icu4x_Locale_region_mv1(IGLocale->locale, write);
    return diplomat_buffer_write_get_bytes(write);
}

const char *ig_get_locale_extension(IGLocale *IGLocale, const char *extension)
{
    DiplomatWrite *write = diplomat_buffer_write_create(0);
    struct DiplomatStringView extension_str = {
        extension,
        strlen(extension)};
    icu4x_Locale_get_unicode_extension_mv1_result result = icu4x_Locale_get_unicode_extension_mv1(IGLocale->locale, extension_str, write);
    if (!result.is_ok)
    {
        return NULL;
    }
    return diplomat_buffer_write_get_bytes(write);
}

void ig_free_locale(IGLocale *IGLocale)
{
    icu4x_Locale_destroy_mv1(IGLocale->locale);
    free(IGLocale);
}