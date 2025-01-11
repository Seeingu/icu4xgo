//go:build !darwin

package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"

func arrayToDiplomatStrings(items []string) C.DiplomatStringsView {
	if len(items) == 0 {
		return C.DiplomatStringsView{
			data: nil,
			len:  0,
		}
	}
	data := make([]C.DiplomatStringView, len(items))
	for i, item := range items {
		data[i] = C.DiplomatStringView{
			data: C.CString(item),
			len:  C.size_t(len(item)),
		}
	}
	return C.DiplomatStringsView{
		data: &data[0],
		len:  C.size_t(len(items)),
	}
}
