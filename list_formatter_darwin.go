package icu4xgo

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

type ListFormatterDarwin struct {
	ListFormatter
	locale Locale
	ptr    *foundation.ListFormatter
}

var _ ListFormatter = (*ListFormatterDarwin)(nil)

func NewListFormatter(locale Locale, initType ListInitType, listLength ListLength) ListFormatter {
	ptr := foundation.NewListFormatter()
	ptr.SetLocale(locale.(*LocaleDarwin).ptr)
	l := &ListFormatterDarwin{
		locale: locale,
		ptr:    &ptr,
	}
	return l
}

// FIXME: only support join by and
func (l *ListFormatterDarwin) Format(items []string) string {
	itemsArr := make([]objc.IObject, len(items))
	for i, item := range items {
		itemsArr[i] = foundation.NewStringWithString(item)
	}
	return l.ptr.StringFromItems(itemsArr)
}
