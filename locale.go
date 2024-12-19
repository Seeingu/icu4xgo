package main

//#cgo LDFLAGS: -L./target/release -L./build/c -licu_capi -licu4xgo
//#include "lib.h"
//#include <stdlib.h>
//#include <string.h>
import "C"

type Locale struct {
	ptr *C.IGLocale
}

func NewLocale(s string) *Locale {
	return &Locale{
		ptr: C.ig_init_locale(C.CString(s)),
	}
}

func (l *Locale) GetLanguage() string {
	return C.GoString(C.ig_get_locale_language(l.ptr))
}

func (l *Locale) Free() {
	C.ig_free_locale(l.ptr)
}
