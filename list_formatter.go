package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"

import (
	"runtime"
)

type ListFormatter struct {
	ptr *C.IGListFormatter
}

type ListInitType int

const (
	ListAnd ListInitType = iota
	ListOr
	ListUnit
)

type ListLength int

const (
	ListLengthWide ListLength = iota
	ListLengthShort
	ListLengthNarrow
)

func NewListFormatter(locale *Locale, initType ListInitType, listLength ListLength) *ListFormatter {
	l := &ListFormatter{
		ptr: C.ig_init_list_formatter(
			locale.ptr,
			C.IGListInitType(initType),
			C.ListLength(listLength),
		),
	}
	runtime.SetFinalizer(l, (*ListFormatter).free)
	return l
}

func (l *ListFormatter) Format(items []string) string {
	listView := arrayToDiplomatStrings(items)
	return C.GoString(C.ig_list_format(l.ptr, listView))
}

func (l *ListFormatter) free() {
	C.ig_free_list_formatter(l.ptr)
}
