package gum

/*
#include <frida-gumjs.h>
*/
import "C"
import (
	"github.com/aadog/go-ffi"
	"github.com/samber/lo"
	"unsafe"
	//"fmt"
)

type SymbolDetails struct {
	IsGlobal bool
	Type     GumSymbolType
	Section  ffi.NativePointer
	Name     string
	Address  ffi.NativePointer
	Size     int64
}

func SymbolDetailsWithNativePointer(ptr ffi.NativePointer) *SymbolDetails {
	th := (*C.GumSymbolDetails)(ptr.Ptr())
	//fmt.Println(th)
	sym := &SymbolDetails{
		Name:     GoString(th.name),
		Address:  ffi.Ptr(uintptr(unsafe.Pointer(uintptr(th.address)))),
		Size:     int64(th.size),
		Section:  ffi.Ptr(uintptr(unsafe.Pointer(th.section))),
		Type:     th._type,
		IsGlobal: bool(lo.If(th.is_global == 1, true).Else(false)),
	}
	return sym
}
