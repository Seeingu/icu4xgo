package icu4xgo

import (
	"testing"

	"github.com/Seeingu/icu4xgo/datetime"
	"github.com/stretchr/testify/assert"
)

func TestDateTimeFormatter(t *testing.T) {
	locale := NewCLocale("en-US")
	dt := NewDateTimeFormatter(locale)
	assert.Equal(t, datetime.Gregorian, dt.CalendarKind())
	dt.SetTimeZone()
	assert.Equal(t, "unk", dt.TimeZoneId())
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
	if assert.NoError(t, err) {
		assert.Equal(t, "February 1, 2021, 12:12:23 AM GMT+?", f)
	}
}
