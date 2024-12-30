package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"
import "runtime"

type NumberFormatter struct {
	ptr *C.IGFixedDecimalFormatter
}

type FixedDecimalGroupingStrategy int

const (
	FixedDecimalGroupingStrategyAuto FixedDecimalGroupingStrategy = iota
	FixedDecimalGroupingStrategyNever
	FixedDecimalGroupingStrategyAlways
	FixedDecimalGroupingStrategyMin2
)

func NewNumberFormatter(locale *Locale) *NumberFormatter {
	n := &NumberFormatter{
		ptr: C.ig_init_fixed_decimal_formatter(locale.ptr, C.FixedDecimalGroupingStrategy(FixedDecimalGroupingStrategyAuto)),
	}
	runtime.SetFinalizer(n, (*NumberFormatter).free)
	return n
}

func (n *NumberFormatter) Format(value string) string {
	return C.GoString(C.ig_fixed_decimal_format(n.ptr, C.CString(value)))
}

func (n *NumberFormatter) free() {
	C.ig_free_fixed_decimal_formatter(n.ptr)
}
