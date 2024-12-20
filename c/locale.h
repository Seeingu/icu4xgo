#ifndef ICUGO_LOCALE_H
#define ICUGO_LOCALE_H

#include "./include/Locale.h"
#include "./include/diplomat_runtime.h"

typedef struct IGLocale
{
    Locale *locale;
    const char *source;
} IGLocale;

const char *ig_get_locale_language(IGLocale *IGLocale);

const char *ig_get_locale_basename(IGLocale *IGLocale);

const char *ig_get_locale_script(IGLocale *IGLocale);

const char *ig_get_locale_region(IGLocale *IGLocale);

/**
 * Returns a string by default. If it fails to get the extension,
 * it returns NULL.
 */
const char *ig_get_locale_extension(IGLocale *IGLocale, const char *extension);

IGLocale *ig_init_locale(const char *localeText);

void ig_free_locale(IGLocale *IGLocale);

#endif // ICUGO_LOCALE_H
