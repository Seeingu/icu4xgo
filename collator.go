package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"
import "runtime"

type Collator struct {
	ptr *C.IGCollator
}

func NewCollator(l *Locale) *Collator {
	options := CollatorOptions{}
	c := &Collator{
		ptr: C.ig_init_collator(l.ptr, options.ToC()),
	}
	runtime.SetFinalizer(c, func(c *Collator) {
		c.free()
	})
	return c
}

func NewCollatorWithOptions(l *Locale, options CollatorOptions) *Collator {
	return &Collator{
		ptr: C.ig_init_collator(l.ptr, options.ToC()),
	}
}

func (c *Collator) Compare(a, b string) int {
	result := C.ig_collator_compare(c.ptr, C.CString(a), C.CString(b))
	return int(result)
}

func (c *Collator) free() {
	C.ig_free_collator(c.ptr)
}
