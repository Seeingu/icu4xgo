package icu4xgo

import (
	"runtime"
	"testing"

	"github.com/Seeingu/icu4xgo/datetime"
	"github.com/stretchr/testify/assert"
)

func TestDateTimeFormatter(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("platform specific test")
	}
	locale := NewLocale("en-US")
	dt := NewDateTimeFormatter(locale)
	assert.Equal(t, datetime.Gregorian, dt.CalendarKind())
	assert.Equal(t, "GMT", dt.TimeZoneId())
	dt.SetDateTime(DateTimeArgs{
		year:       2021,
		month:      2,
		day:        1,
		hour:       0,
		minute:     12,
		second:     23,
		nanosecond: 23,
	})
	tz := "Australia/Sydney"
	dt.SetTimeZone(tz)
	assert.Equal(t, tz, dt.TimeZoneId())
	assert.Equal(t, 32, dt.DayOfYear())
	dt.SetFormatter(DateTimeLengthLong)
	f, err := dt.Format()
	assert.NoError(t, err)
	assert.Equal(t, "February 1, 2021 at 11:12:23 AM GMT+11", f)
}
