#ifndef ICUGO_UTILS_H
#define ICUGO_UTILS_H

typedef struct
{
    int code;
    const char *message;
} EnumMapping;

const char *ig_format_enum(int code, EnumMapping *mapping, int count);

#endif // ICUGO_UTILS_H