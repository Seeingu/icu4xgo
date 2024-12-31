#include "./ig_string.h"
#include <stdlib.h>
#include <string.h>

IGStringWriter *ig_init_string_writer()
{
    IGStringWriter *w = malloc(sizeof(IGStringWriter));
    w->write = diplomat_buffer_write_create(0);
    return w;
}

const char *ig_string_writer_to_string(IGStringWriter *w)
{
    size_t length = diplomat_buffer_write_len(w->write);
    if (length == 0)
    {
        return "";
    }
    return diplomat_buffer_write_get_bytes(w->write);
}

DiplomatStringView ig_init_string(const char *s)
{
    return (DiplomatStringView){
        .data = s,
        .len = strlen(s),
    };
}