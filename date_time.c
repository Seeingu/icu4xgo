#include "./ig_date_time.h"
#include "./ig_diplomat_string.h"
#include <stdio.h>

IGDateTime *ig_init_datetime_from_codes(IGDateTimeCalendarCodes codes)
{
    icu4x_DateTime_from_iso_in_calendar_mv1_result result = icu4x_DateTime_from_iso_in_calendar_mv1(
        codes.year,
        codes.month,
        codes.day,
        codes.hour,
        codes.minute,
        codes.second,
        codes.nanosecond,
        codes.calendar->calendar);
    IGDateTime *dt = malloc(sizeof(IGDateTime));
    if (!result.is_ok)
    {
        return dt;
    }
    dt->datetime = result.ok;
    return dt;
}

const char *ig_datetime_format(IGZonedDateTimeFormatter *formatter, IGDateTime *datetime, IGTimeZoneInfo *timezone)
{
    IGStringWriter *w = ig_init_string_writer();
    icu4x_ZonedDateTimeFormatter_format_datetime_with_custom_time_zone_mv1_result result =
        icu4x_ZonedDateTimeFormatter_format_datetime_with_custom_time_zone_mv1(
            formatter->formatter,
            datetime->datetime,
            timezone->timezone,
            w->write);
    if (!result.is_ok)
    {
        char err[sizeof(int)];
        sprintf(err, "format err: %d", result.err);
        return err;
    }
    return ig_string_writer_to_string(w);
}

void ig_free_datetime(IGDateTime *datetime)
{
    icu4x_DateTime_destroy_mv1(datetime->datetime);
    free(datetime);
}

IGZonedDateTimeFormatter *ig_init_zoned_date_time_formatter(const IGLocale *locale, DateTimeLength length)
{
    icu4x_ZonedDateTimeFormatter_create_with_length_mv1_result result = icu4x_ZonedDateTimeFormatter_create_with_length_mv1(
        icu4x_DataProvider_compiled_mv1(),
        locale->locale,
        length);
    IGZonedDateTimeFormatter *formatter = malloc(sizeof(IGZonedDateTimeFormatter));
    if (!result.is_ok)
    {
        return formatter;
    }
    formatter->formatter = result.ok;
    return formatter;
}

void ig_free_zoned_date_time_formatter(IGZonedDateTimeFormatter *formatter)
{
    icu4x_ZonedDateTimeFormatter_destroy_mv1(formatter->formatter);
    free(formatter);
}

IGTimeZoneInfo *ig_init_time_zone_info(void)
{
    IGTimeZoneInfo *tz = malloc(sizeof(IGTimeZoneInfo));
    TimeZoneInfo *timezone = icu4x_TimeZoneInfo_unknown_mv1();
    icu4x_TimeZoneInfo_set_standard_time_mv1(timezone);
    tz->timezone = timezone;
    return tz;
}

const char *ig_time_zone_info_id(IGTimeZoneInfo *tz)
{
    DiplomatWrite *write = diplomat_buffer_write_create(0);
    icu4x_TimeZoneInfo_time_zone_id_mv1(tz->timezone, write);
    return diplomat_buffer_write_get_bytes(write);
}

void ig_free_time_zone_info(IGTimeZoneInfo *timezone)
{
    icu4x_TimeZoneInfo_destroy_mv1(timezone->timezone);
    free(timezone);
}

IGCalendar *ig_init_calendar(const IGLocale *locale)
{
    DataProvider *dataProvider = icu4x_DataProvider_compiled_mv1();
    icu4x_Calendar_create_for_locale_mv1_result result = icu4x_Calendar_create_for_locale_mv1(dataProvider, locale->locale);
    IGCalendar *calendar = malloc(sizeof(IGCalendar));
    if (!result.is_ok)
    {
        return calendar;
    }
    calendar->calendar = result.ok;
    return calendar;
}

AnyCalendarKind ig_calendar_kind(IGCalendar *c)
{
    return icu4x_Calendar_kind_mv1(c->calendar);
}

void ig_free_calendar(IGCalendar *calendar)
{
    icu4x_Calendar_destroy_mv1(calendar->calendar);
    free(calendar);
}