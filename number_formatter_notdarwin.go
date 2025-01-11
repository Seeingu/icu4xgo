//go:build !darwin

package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"
import "runtime"

func NewNumberFormatter(locale Locale) NumberFormatter {
	return func(locale Locale) NumberFormatter {
		return newCNumberFormatter(locale.(*CLocale))
	}(locale)
}

type CNumberFormatter struct {
	NumberFormatter
	ptr *C.IGFixedDecimalFormatter
}

var _ NumberFormatter = (*CNumberFormatter)(nil)

type FixedDecimalGroupingStrategy int

const (
	FixedDecimalGroupingStrategyAuto FixedDecimalGroupingStrategy = iota
	FixedDecimalGroupingStrategyNever
	FixedDecimalGroupingStrategyAlways
	FixedDecimalGroupingStrategyMin2
)

func newCNumberFormatter(locale *CLocale) *CNumberFormatter {
	n := &CNumberFormatter{
		ptr: C.ig_init_fixed_decimal_formatter(locale.ptr, C.FixedDecimalGroupingStrategy(FixedDecimalGroupingStrategyAuto)),
	}
	runtime.SetFinalizer(n, (*CNumberFormatter).free)
	return n
}

func (n *CNumberFormatter) Format(value string) (string, error) {
	return C.GoString(C.ig_fixed_decimal_format(n.ptr, C.CString(value))), nil
}

func (n *CNumberFormatter) free() {
	C.ig_free_fixed_decimal_formatter(n.ptr)
}
