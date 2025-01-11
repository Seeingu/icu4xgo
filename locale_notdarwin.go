//go:build !darwin
// +build !darwin

package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"

import (
	"fmt"
	"runtime"
)

type CLocale struct {
	Locale
	ptr *C.IGLocale
}

var _ Locale = (*CLocale)(nil)

func NewLocale(s string) Locale {
	l := &CLocale{
		ptr: C.ig_init_locale(C.CString(s)),
	}
	runtime.SetFinalizer(l, func(l *CLocale) {
		l.free()
	})
	return l
}

func (l *CLocale) Language() string {
	return C.GoString(C.ig_get_locale_language(l.ptr))
}

func (l *CLocale) Script() string {
	return C.GoString(C.ig_get_locale_script(l.ptr))
}

func (l *CLocale) Region() string {
	return C.GoString(C.ig_get_locale_region(l.ptr))
}

func (l *CLocale) BaseName() string {
	return C.GoString(C.ig_get_locale_basename(l.ptr))
}

func (l *CLocale) Extension(name string) (s string, err error) {
	r := C.ig_get_locale_extension(l.ptr, C.CString(name))
	if r == nil {
		return "", fmt.Errorf("extension %s not found", name)
	}
	return C.GoString(r), nil
}

func (l *CLocale) HourCycle() string {
	hc, err := l.Extension("hc")
	if err != nil {
		return ""
	}
	return hc
}

func (l *CLocale) Calendar() string {
	ca, err := l.Extension("ca")
	if err != nil {
		return ""
	}
	return ca
}

func (l *CLocale) CaseFirst() string {
	kf, err := l.Extension("kf")
	if err != nil {
		return ""
	}
	return kf
}

func (l *CLocale) Collation() string {
	co, err := l.Extension("co")
	if err != nil {
		return ""
	}
	return co
}

func (l *CLocale) NumberingSystem() string {
	ns, err := l.Extension("nu")
	if err != nil {
		return ""
	}
	return ns
}

func (l *CLocale) Numeric() bool {
	nu, err := l.Extension("kn")
	if err != nil {
		return false
	}
	return nu == "true"
}

func (l *CLocale) free() {
	C.ig_free_locale(l.ptr)
}
