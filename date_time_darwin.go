package icu4xgo

import (
	"github.com/Seeingu/icu4xgo/datetime"
	"github.com/progrium/darwinkit/macos/foundation"
)

type DateTimeFormatterDarwin struct {
	DateTimeFormatter
	locale     Locale
	formatter  *foundation.DateFormatter
	date       foundation.Date
	components foundation.DateComponents
}

var _ DateTimeFormatter = (*DateTimeFormatterDarwin)(nil)

func NewDateTimeFormatter(l Locale) DateTimeFormatter {
	f := foundation.NewDateFormatter()
	f.SetLocale(GetFoundationLocale(l))
	// Set default timezone to GMT
	f.SetTimeZone(foundation.NewTimeZoneWithName("GMT"))
	d := &DateTimeFormatterDarwin{
		locale:    l,
		formatter: &f,
	}
	return d
}

func (d *DateTimeFormatterDarwin) CalendarKind() datetime.AnyCalendarKind {
	id := d.formatter.Calendar().CalendarIdentifier()
	return foundationCalendarKindToAnyCalendarKind(id)
}

func (d *DateTimeFormatterDarwin) TimeZoneId() string {
	return d.formatter.TimeZone().Name()
}

func (d *DateTimeFormatterDarwin) SetDateTime(args DateTimeArgs) DateTimeFormatter {
	components := foundation.NewDateComponents()
	components.SetCalendar(d.formatter.Calendar())
	components.SetTimeZone(d.formatter.TimeZone())
	components.SetYear(args.year)
	components.SetMonth(args.month)
	components.SetDay(args.day)
	components.SetHour(args.hour)
	components.SetMinute(args.minute)
	components.SetSecond(args.second)
	components.SetNanosecond(args.nanosecond)
	d.components = components
	d.date = components.Date()
	return d
}

func (d *DateTimeFormatterDarwin) SetTimeZone(tz string) DateTimeFormatter {
	d.formatter.SetTimeZone(foundation.NewTimeZoneWithName(tz))
	return d
}

func (d *DateTimeFormatterDarwin) SetFormatter(length DateTimeLength) DateTimeFormatter {
	switch length {
	case DateTimeLengthLong:
		d.formatter.SetDateStyle(foundation.DateFormatterLongStyle)
		d.formatter.SetTimeStyle(foundation.DateFormatterLongStyle)
	case DateTimeLengthMedium:
		d.formatter.SetDateStyle(foundation.DateFormatterMediumStyle)
		d.formatter.SetTimeStyle(foundation.DateFormatterMediumStyle)
	case DateTimeLengthShort:
		d.formatter.SetDateStyle(foundation.DateFormatterShortStyle)
		d.formatter.SetTimeStyle(foundation.DateFormatterShortStyle)
	}
	return d
}

func (d *DateTimeFormatterDarwin) Format() (s string, err error) {
	return d.formatter.StringFromDate(d.date), nil
}

func (d *DateTimeFormatterDarwin) DayOfYear() int {
	day := d.formatter.Calendar().OrdinalityOfUnitInUnitForDate(
		foundation.CalendarUnitDay, foundation.CalendarUnitYear, d.date)
	return int(day)
}

// MARK: - utils

func foundationCalendarKindToAnyCalendarKind(id foundation.CalendarIdentifier) datetime.AnyCalendarKind {
	switch id {
	case foundation.CalendarIdentifierGregorian:
		return datetime.Gregorian
	case foundation.CalendarIdentifierBuddhist:
		return datetime.Buddhist
	case foundation.CalendarIdentifierChinese:
		return datetime.Chinese
	case foundation.CalendarIdentifierCoptic:
		return datetime.Coptic
	case foundation.CalendarIdentifierEthiopicAmeteAlem:
		return datetime.EthiopianAmeteAlem
	case foundation.CalendarIdentifierEthiopicAmeteMihret:
		return datetime.Ethiopian
	case foundation.CalendarIdentifierHebrew:
		return datetime.Hebrew
	case foundation.CalendarIdentifierIndian:
		return datetime.Indian
	case foundation.CalendarIdentifierIslamic:
		return datetime.IslamicObservational
	case foundation.CalendarIdentifierIslamicCivil:
		return datetime.IslamicCivil
	case foundation.CalendarIdentifierJapanese:
		return datetime.Japanese
	case foundation.CalendarIdentifierPersian:
		return datetime.Persian
	case foundation.CalendarIdentifierRepublicOfChina:
		return datetime.Roc
	case foundation.CalendarIdentifierIslamicTabular:
		return datetime.IslamicTabular
	case foundation.CalendarIdentifierIslamicUmmAlQura:
		return datetime.IslamicUmmAlQura
	default:
		return datetime.Unknown
	}
}
