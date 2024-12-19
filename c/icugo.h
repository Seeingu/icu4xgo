#ifndef ICUGO_H
#define ICUGO_H

#include "./include/Locale.h"
#include "./include/diplomat_runtime.h"

typedef struct IGLocale
{
    Locale *locale;
    const char *source;
} IGLocale;

const char *ig_get_locale_language(IGLocale *IGLocale);

IGLocale *ig_init_locale(const char *localeText);

void ig_free_locale(IGLocale *IGLocale);

#endif // ICUGO_H