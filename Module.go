package gumjs

/*
#include <frida-gumjs.h>
*/
import "C"
import (
	"errors"
	"github.com/aadog/go-ffi"
	"github.com/samber/lo"
	"github.com/samber/mo"
	"unsafe"
)

var Module ModuleStruct

type ModuleStruct struct {
}

func (m ModuleStruct) GetBaseAddress(name string) ffi.NativePointer {
	cModule_name := CString(name)
	return ffi.Ptr(uintptr(C.gum_module_find_base_address(cModule_name)))
}
func (m ModuleStruct) ModuleLoad(moduleName string) mo.Result[bool] {
	cModule_name := CString(moduleName)
	var err *C.GError
	b := C.gum_module_load(cModule_name, &err)
	if err != nil {
		return mo.Err[bool](ConvertGErrorAndFree(unsafe.Pointer(err)))
	}
	if b == 0 {
		return mo.Err[bool](errors.New("load error"))
	}
	return mo.Ok(lo.If(b < 1, true).Else(false))
}
func (m ModuleStruct) EnsureInitialized(moduleName string) bool {
	cModule_name := CString(moduleName)
	b := C.gum_module_ensure_initialized(cModule_name)
	return lo.If(b < 1, true).Else(false)
}
func (m ModuleStruct) FindExportByName(moduleName *string, SymbolName string) ffi.NativePointer {
	var cModule_name *C.char = nil
	if moduleName != nil {
		cModule_name = CString(*moduleName)
	}
	cSymbolName := CString(SymbolName)
	return ffi.Ptr(uintptr(C.gum_module_find_export_by_name(cModule_name, cSymbolName)))
}
func (m ModuleStruct) FindSymbolByName(moduleName *string, SymbolName string) ffi.NativePointer {
	var cModule_name *C.char = nil
	if moduleName != nil {
		cModule_name = C.CString(*moduleName)
	}
	cSymbolName := C.CString(SymbolName)
	return ffi.Ptr(uintptr(C.gum_module_find_symbol_by_name(cModule_name, cSymbolName)))
}
