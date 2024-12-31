#ifndef ICUGO_DIPLOAMAT_STRING_H
#define ICUGO_DIPLOAMAT_STRING_H

#include "./icu4x/diplomat_runtime.h"

typedef struct IGStringWriter
{
    DiplomatWrite *write;
} IGStringWriter;

IGStringWriter *ig_init_string_writer();
const char *ig_string_writer_to_string(IGStringWriter *view);

DiplomatStringView ig_init_string(const char *s);

#endif // ICUGO_DIPLOAMAT_STRING_H
