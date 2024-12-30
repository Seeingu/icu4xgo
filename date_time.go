package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"
import "errors"

type DateTimeFormatter struct {
	tzPtr        *C.IGTimeZoneInfo
	datetimePtr  *C.IGDateTime
	formatterPtr *C.IGZonedDateTimeFormatter
	locale       *Locale
	calendar     *C.IGCalendar
}

func NewDateTimeFormatter(l *Locale) *DateTimeFormatter {
	f := &DateTimeFormatter{}
	f.locale = l
	f.calendar = C.ig_init_calendar(l.ptr)
	return f
}

func (f *DateTimeFormatter) CalendarKind() int {
	kind := C.ig_calendar_kind(f.calendar)
	return int(kind)
}

func (f *DateTimeFormatter) TimeZoneId() string {
	return C.GoString(C.ig_time_zone_info_id(f.tzPtr))
}

func (f *DateTimeFormatter) SetDateTime() *DateTimeFormatter {
	codes := C.IGDateTimeCalendarCodes{
		year:       2021,
		month:      2,
		day:        1,
		hour:       0,
		minute:     12,
		second:     23,
		nanosecond: 0,
		calendar:   f.calendar,
	}
	f.datetimePtr = C.ig_init_datetime_from_codes(codes)
	return f
}

func (f *DateTimeFormatter) SetTimeZone() *DateTimeFormatter {
	f.tzPtr = C.ig_init_time_zone_info()
	return f
}

type DateTimeLength int

const (
	DateTimeLengthLong DateTimeLength = iota
	DateTimeLengthMedium
	DateTimeLengthShort
)

func (f *DateTimeFormatter) SetFormatter(length DateTimeLength) *DateTimeFormatter {
	f.formatterPtr = C.ig_init_zoned_date_time_formatter(f.locale.ptr, C.DateTimeLength(length))
	return f
}

func (f *DateTimeFormatter) Format() (s string, err error) {
	result := C.GoString(C.ig_datetime_format(f.formatterPtr, f.datetimePtr, f.tzPtr))
	if f.formatterPtr.is_ok {
		return result, nil
	}
	return "", errors.New(result)
}

func (f *DateTimeFormatter) DayOfYear() int {
	return int(C.icu4x_DateTime_day_of_year_mv1(f.datetimePtr.datetime))
}

func (f *DateTimeFormatter) Free() {
	C.ig_free_time_zone_info(f.tzPtr)
	C.ig_free_datetime(f.datetimePtr)
	C.ig_free_zoned_date_time_formatter(f.formatterPtr)
}
