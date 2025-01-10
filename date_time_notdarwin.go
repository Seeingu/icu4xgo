//go:build !darwin
// +build !darwin

package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"

import (
	"errors"
	"runtime"

	"github.com/Seeingu/icu4xgo/datetime"
)

type CDateTimeFormatter struct {
	DateTimeFormatter
	tzPtr        *C.IGTimeZoneInfo
	datetimePtr  *C.IGDateTime
	formatterPtr *C.IGZonedDateTimeFormatter
	locale       *CLocale
	calendar     *C.IGCalendar
}

var _ DateTimeFormatter = (*CDateTimeFormatter)(nil)

func NewDateTimeFormatter(l Locale) DateTimeFormatter {
	f := &CDateTimeFormatter{}
	locale := l.(*CLocale)
	f.locale = locale
	f.calendar = C.ig_init_calendar(locale.ptr)
	runtime.SetFinalizer(f, (*CDateTimeFormatter).free)
	return f
}

func (f *CDateTimeFormatter) CalendarKind() datetime.AnyCalendarKind {
	kind := C.ig_calendar_kind(f.calendar)
	return datetime.AnyCalendarKind(kind)
}

func (f *CDateTimeFormatter) TimeZoneId() string {
	return C.GoString(C.ig_time_zone_info_id(f.tzPtr))
}

func (f *CDateTimeFormatter) SetDateTime(args DateTimeArgs) DateTimeFormatter {
	codes := C.IGDateTimeCalendarCodes{
		year:       C.int(args.year),
		month:      C.uint8_t(args.month),
		day:        C.uint8_t(args.day),
		hour:       C.uint8_t(args.hour),
		minute:     C.uint8_t(args.minute),
		second:     C.uint8_t(args.second),
		nanosecond: C.uint32_t(args.nanosecond),
		calendar:   f.calendar,
	}
	f.datetimePtr = C.ig_init_datetime_from_codes(codes)
	return f
}

func (f *CDateTimeFormatter) SetTimeZone() DateTimeFormatter {
	f.tzPtr = C.ig_init_time_zone_info()
	return f
}

func (f *CDateTimeFormatter) SetFormatter(length DateTimeLength) DateTimeFormatter {
	f.formatterPtr = C.ig_init_zoned_date_time_formatter(f.locale.ptr, C.DateTimeLength(length))
	return f
}

func (f *CDateTimeFormatter) Format() (s string, err error) {
	result := C.GoString(C.ig_datetime_format(f.formatterPtr, f.datetimePtr, f.tzPtr))
	if f.formatterPtr.is_ok {
		return result, nil
	}
	return "", errors.New(result)
}

func (f *CDateTimeFormatter) DayOfYear() int {
	return int(C.icu4x_DateTime_day_of_year_mv1(f.datetimePtr.datetime))
}

func (f *CDateTimeFormatter) free() {
	C.ig_free_calendar(f.calendar)
	C.ig_free_time_zone_info(f.tzPtr)
	C.ig_free_datetime(f.datetimePtr)
	C.ig_free_zoned_date_time_formatter(f.formatterPtr)
}
