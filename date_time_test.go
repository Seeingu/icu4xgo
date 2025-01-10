package icu4xgo

import (
	"runtime"
	"testing"

	"github.com/Seeingu/icu4xgo/datetime"
	"github.com/stretchr/testify/assert"
)

func TestDateTimeFormatter(t *testing.T) {
	locale := NewLocale("en-US")
	dt := NewDateTimeFormatter(locale)
	assert.Equal(t, datetime.Gregorian, dt.CalendarKind())
	tz := "Australia/Sydney"
	dt.SetTimeZone(tz)
	assert.Equal(t, tz, dt.TimeZoneId())
	dt.SetDateTime(DateTimeArgs{
		year:       2021,
		month:      2,
		day:        1,
		hour:       0,
		minute:     12,
		second:     23,
		nanosecond: 23,
	})
	assert.Equal(t, 32, dt.DayOfYear())
	dt.SetFormatter(DateTimeLengthLong)
	f, err := dt.Format()
	assert.NoError(t, err)
	if runtime.GOOS == "darwin" {
		assert.Equal(t, "February 1, 2021 at 3:12:23 AM GMT+11", f)
	} else {
		// FIXME(linux)
		assert.Equal(t, "February 1, 2021, 12:12:23 AM GMT+?", f)
	}
}
