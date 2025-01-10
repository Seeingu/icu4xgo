//go:build !darwin
// +build !darwin

package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"

import (
	"runtime"
)

type CListFormatter struct {
	ListFormatter
	ptr *C.IGListFormatter
}

var _ ListFormatter = (*CListFormatter)(nil)

func newCListFormatter(locale *CLocale, initType ListInitType, listLength ListLength) *CListFormatter {
	l := &CListFormatter{
		ptr: C.ig_init_list_formatter(
			locale.ptr,
			C.IGListInitType(initType),
			C.ListLength(listLength),
		),
	}
	runtime.SetFinalizer(l, (*CListFormatter).free)
	return l
}

func (l *CListFormatter) Format(items []string) string {
	listView := arrayToDiplomatStrings(items)
	return C.GoString(C.ig_list_format(l.ptr, listView))
}

func (l *CListFormatter) free() {
	C.ig_free_list_formatter(l.ptr)
}

func NewListFormatter(locale Locale, initType ListInitType, listLength ListLength) ListFormatter {
	return newCListFormatter(locale.(*CLocale), initType, listLength)
}
