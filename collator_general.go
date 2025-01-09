package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"
import "runtime"

type CCollator struct {
	Collator
	ptr *C.IGCollator
}

var _ Collator = (*CCollator)(nil)

func NewCCollator(l *CLocale) *CCollator {
	options := CollatorOptions{}
	c := &CCollator{
		ptr: C.ig_init_collator(l.ptr, options.ToC()),
	}
	runtime.SetFinalizer(c, func(c *CCollator) {
		c.free()
	})
	return c
}

func NewCollatorWithOptions(l *CLocale, options CollatorOptions) *CCollator {
	return &CCollator{
		ptr: C.ig_init_collator(l.ptr, options.ToC()),
	}
}

func (c *CCollator) Compare(a, b string) int {
	result := C.ig_collator_compare(c.ptr, C.CString(a), C.CString(b))
	return int(result)
}

func (c *CCollator) free() {
	C.ig_free_collator(c.ptr)
}
