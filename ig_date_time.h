#ifndef ICUGO_DATE_TIME_H
#define ICUGO_DATE_TIME_H

#include "./ig_locale.h"
#include "./icu4x/DateTime.h"
#include "./icu4x/TimeZoneInfo.h"
#include "./icu4x/ZonedDateTimeFormatter.h"
#include "./icu4x/DataProvider.h"
#include "./icu4x/Calendar.h"

#include <stdlib.h>

typedef struct IGDateTime
{
    DateTime *datetime;
} IGDateTime;

typedef struct IGZonedDateTimeFormatter
{
    ZonedDateTimeFormatter *formatter;
    bool is_ok;
} IGZonedDateTimeFormatter;

typedef struct IGTimeZoneInfo
{
    TimeZoneInfo *timezone;
} IGTimeZoneInfo;

typedef struct IGCalendar
{
    Calendar *calendar;
} IGCalendar;

typedef struct IGDateTimeCalendarCodes
{
    int32_t year;
    uint8_t month;
    uint8_t day;
    uint8_t hour;
    uint8_t minute;
    uint8_t second;
    uint32_t nanosecond;
    IGCalendar *calendar;
} IGDateTimeCalendarCodes;

IGDateTime *ig_init_datetime_from_codes(IGDateTimeCalendarCodes codes);

const char *ig_datetime_format(
    IGZonedDateTimeFormatter *formatter,
    IGDateTime *datetime,
    IGTimeZoneInfo *timezone);

IGZonedDateTimeFormatter *ig_init_zoned_date_time_formatter(const IGLocale *locale, DateTimeLength length);

IGTimeZoneInfo *ig_init_time_zone_info(void);

const char *ig_time_zone_info_id(IGTimeZoneInfo *tz);

IGCalendar *ig_init_calendar(const IGLocale *locale);

AnyCalendarKind ig_calendar_kind(IGCalendar *c);

void ig_free_calendar(IGCalendar *calendar);
void ig_free_zoned_date_time_formatter(IGZonedDateTimeFormatter *formatter);
void ig_free_datetime(IGDateTime *datetime);
void ig_free_time_zone_info(IGTimeZoneInfo *timezone);

#endif // ICUGO_DATE_TIME_H