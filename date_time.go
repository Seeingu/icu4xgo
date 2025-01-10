package icu4xgo

import (
	"github.com/Seeingu/icu4xgo/datetime"
)

type DateTimeFormatter interface {
	CalendarKind() datetime.AnyCalendarKind
	TimeZoneId() string
	SetDateTime(args DateTimeArgs) DateTimeFormatter
	SetTimeZone(tz string) DateTimeFormatter
	SetFormatter(length DateTimeLength) DateTimeFormatter
	Format() (s string, err error)
	DayOfYear() int
}

type DateTimeArgs struct {
	year, month, day, hour, minute, second, nanosecond int
}

type DateTimeLength int

const (
	DateTimeLengthLong DateTimeLength = iota
	DateTimeLengthMedium
	DateTimeLengthShort
)
