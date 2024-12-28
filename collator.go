package icu4xgo

//#cgo LDFLAGS: -L./lib -licu_capi -lm
//#include "./icu4xgo.h"
//#include <stdlib.h>
//#include <string.h>
import "C"

type Collator struct {
	ptr *C.IGCollator
}

func NewCollator(l *Locale) *Collator {
	return &Collator{
		ptr: C.ig_init_collator(l.ptr),
	}
}

func (c *Collator) Compare(a, b string) int {
	result := C.ig_collator_compare(c.ptr, C.CString(a), C.CString(b))
	return int(result)
}

func (c *Collator) Free() {
	C.ig_free_collator(c.ptr)
}
