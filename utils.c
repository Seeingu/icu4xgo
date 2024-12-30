#include "./ig_utils.h"

const char *ig_format_enum(int code, EnumMapping *mapping, int count)
{
    for (int i = 0; i < count; i++)
    {
        if (mapping[i].code == code)
        {
            return mapping[i].message;
        }
    }
    return "Unknown Error";
}