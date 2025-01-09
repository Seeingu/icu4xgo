package icu4xgo

import (
	"strconv"

	"github.com/progrium/darwinkit/macos/foundation"
)

type NumberFormatterDarwin struct {
	NumberFormatter
	locale Locale
	ptr    foundation.NumberFormatter
}

var _ NumberFormatter = (*NumberFormatterDarwin)(nil)

func NewNumberFormatter(locale Locale) NumberFormatter {
	f := foundation.NewNumberFormatter()
	f.SetLocale(locale.(*LocaleDarwin).ptr)
	f.SetAllowsFloats(true)
	f.SetNumberStyle(foundation.NumberFormatterDecimalStyle)
	return &NumberFormatterDarwin{
		locale: locale,
		ptr:    f,
	}
}

func (n *NumberFormatterDarwin) Format(value string) (string, error) {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return "", NumberFormatFailedError
	}
	number := foundation.Number_NumberWithDouble(f)
	return n.ptr.StringFromNumber(number), nil
}
