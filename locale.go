package icu4xgo

//#cgo LDFLAGS: -L./target/release -L./build/c -licu4xgo -licu_capi -lm
//#include "./lib.h"
//#include <stdlib.h>
//#include <string.h>
import "C"
import "fmt"

type Locale struct {
	ptr *C.IGLocale
}

func NewLocale(s string) *Locale {
	return &Locale{
		ptr: C.ig_init_locale(C.CString(s)),
	}
}

func (l *Locale) Language() string {
	return C.GoString(C.ig_get_locale_language(l.ptr))
}

func (l *Locale) Script() string {
	return C.GoString(C.ig_get_locale_script(l.ptr))
}

func (l *Locale) Region() string {
	return C.GoString(C.ig_get_locale_region(l.ptr))
}

func (l *Locale) BaseName() string {
	return C.GoString(C.ig_get_locale_basename(l.ptr))
}

func (l *Locale) Extension(name string) (s string, err error) {
	r := C.ig_get_locale_extension(l.ptr, C.CString(name))
	if r == nil {
		return "", fmt.Errorf("extension %s not found", name)
	}
	return C.GoString(r), nil
}

func (l *Locale) HourCycle() (string, bool) {
	hc, err := l.Extension("hc")
	if err != nil {
		return "", false
	}
	return hc, true
}

func (l *Locale) Calendar() (string, bool) {
	ca, err := l.Extension("ca")
	if err != nil {
		return "", false
	}
	return ca, true
}

func (l *Locale) CaseFirst() (string, bool) {
	kf, err := l.Extension("kf")
	if err != nil {
		return "", false
	}
	return kf, true
}

func (l *Locale) Collation() (string, bool) {
	co, err := l.Extension("co")
	if err != nil {
		return "", false
	}
	return co, true
}

func (l *Locale) NumberingSystem() (string, bool) {
	ns, err := l.Extension("nu")
	if err != nil {
		return "", false
	}
	return ns, true
}

func (l *Locale) Numeric() (bool, bool) {
	nu, err := l.Extension("kn")
	if err != nil {
		return false, false
	}
	return nu == "true", true
}

func (l *Locale) Free() {
	C.ig_free_locale(l.ptr)
}
