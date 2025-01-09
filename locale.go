package icu4xgo

type Locale interface {
	Language() string
	Script() string
	Region() string
	BaseName() string
	Extension(name string) (s string, err error)
	HourCycle() string
	Calendar() string
	CaseFirst() string
	Collation() string
	NumberingSystem() string
	Numeric() bool
}
