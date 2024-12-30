package icu4xgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateTimeFormatter(t *testing.T) {
	locale := NewLocale("en-US")
	defer locale.Free()
	dt := NewDateTimeFormatter(locale)
	defer dt.Free()
	assert.Equal(t, 1, dt.CalendarKind())
	dt.SetTimeZone()
	assert.Equal(t, "unk", dt.TimeZoneId())
	dt.SetDateTime()
	assert.Equal(t, 32, dt.DayOfYear())
	dt.SetFormatter(DateTimeLengthLong)
	assert.Equal(t, "February 1, 2021, 12:12:23 AM GMT+?", dt.Format())
}
