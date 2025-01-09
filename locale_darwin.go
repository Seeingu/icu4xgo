package icu4xgo

import (
	"github.com/progrium/darwinkit/macos/foundation"
)

type LocaleDarwin struct {
	Locale
	ptr             foundation.Locale
	identifier      string
	hourCycle       string
	countryCode     string
	numberingSystem string
	caseFirst       string
}

var _ Locale = (*LocaleDarwin)(nil)

func NewLocaleMac(identifier string) Locale {
	locale := foundation.Locale_LocaleWithLocaleIdentifier(identifier)

	variants := foundation.Locale_ComponentsFromLocaleIdentifier(identifier)

	return &LocaleDarwin{
		ptr:             locale,
		hourCycle:       variants["hours"],
		countryCode:     variants[string(foundation.LocaleCountryCode)],
		numberingSystem: variants["numbers"],
		caseFirst:       variants["colcasefirst"],
	}
}

func (l *LocaleDarwin) Language() string {
	return l.ptr.LanguageCode()
}

func (l *LocaleDarwin) Script() string {
	return l.ptr.ScriptCode()
}

func (l *LocaleDarwin) Region() string {
	return l.countryCode
}

func (l *LocaleDarwin) BaseName() string {
	s := l.Language()
	if l.Script() != "" {
		s += "-" + l.Script()
	}
	if l.Region() != "" {
		s += "-" + l.Region()
	}
	return s
}

func (l *LocaleDarwin) Extension(name string) (s string, err error) {
	return "", nil
}

func (l *LocaleDarwin) HourCycle() string {
	return l.hourCycle
}

func (l *LocaleDarwin) Calendar() string {
	return l.ptr.CalendarIdentifier()
}

func (l *LocaleDarwin) CaseFirst() string {
	return l.caseFirst
}

func (l *LocaleDarwin) Collation() string {
	return l.ptr.CollationIdentifier()
}

func (l *LocaleDarwin) NumberingSystem() string {
	return l.numberingSystem
}

func (l *LocaleDarwin) Numeric() bool {
	return false
}
