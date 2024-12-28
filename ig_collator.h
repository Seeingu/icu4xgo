#ifndef ICUGO_COLLATOR_H
#define ICUGO_COLLATOR_H

#include "./ig_locale.h"
#include "./icu4x/Collator.h"
#include "./icu4x/CollatorMaxVariable.h"
#include "./icu4x/DataProvider.h"

typedef struct IGCollator
{
    Collator *collator;
} IGCollator;

int8_t ig_collator_compare(IGCollator *c, const char *left, const char *right);

IGCollator *ig_init_collator(IGLocale *locale);

void ig_free_collator(IGCollator *collator);

#endif // ICUGO_COLLATOR_H