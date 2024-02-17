package gumjs

/*
#include <frida-gumjs.h>
*/
import "C"
import (
	"github.com/aadog/go-ffi"
	"unsafe"
)

type RangeDetails struct {
	Base       ffi.NativePointer
	Size       int64
	Protection GumPageProtection
}

func RangeDetailsWithPtr(ptr unsafe.Pointer) *RangeDetails {
	details := (*C.GumMallocRangeDetails)(ptr)
	return &RangeDetails{
		Base: ffi.Ptr(uintptr(details._range.base_address)),
		Size: int64(details._range.size),
		//Protection: details._range.protection,
	}
}
