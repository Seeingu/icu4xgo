package datetime

//go:generate stringer -type=AnyCalendarKind
type AnyCalendarKind int

const (
	Iso AnyCalendarKind = iota
	Gregorian
	Buddhist
	Japanese
	JapaneseExtended
	Ethiopian
	EthiopianAmeteAlem
	Indian
	Coptic
	Dangi
	Chinese
	Hebrew
	IslamicCivil
	IslamicObservational
	IslamicTabular
	IslamicUmmAlQura
	Persian
	Roc
)
