package gumjs

/*
#include <frida-gumjs.h>
*/
import "C"
import (
	"github.com/aadog/go-ffi"
	"unsafe"
)

type ExportDetails struct {
	Type    GumExportType
	Name    string
	Address ffi.NativePointer
}

func ExportDetailsWithPtr(ptr unsafe.Pointer) *ExportDetails {
	th := (*C.GumExportDetails)(ptr)

	return &ExportDetails{
		Name:    GoString(th.name),
		Address: ffi.Ptr(uintptr(unsafe.Pointer(uintptr(th.address)))),
		Type:    th._type,
	}
}
