package gumjs

/*
#include <frida-gumjs.h>
*/
import "C"
import (
	"fmt"
	"github.com/aadog/go-ffi"
	"github.com/samber/lo"
	"regexp"
	"strings"
	"unsafe"
)

var Process ProcessStruct

type ProcessStruct struct {
}

func (p ProcessStruct) Id() uint {
	return uint(C.gum_process_get_id())
}

func (p ProcessStruct) GetCurrentThreadId() uint {
	return uint(C.gum_process_get_current_thread_id())
}
func (p ProcessStruct) IsDebuggerAttached() bool {
	return bool(lo.If(C.gum_process_is_debugger_attached() == 1, true).Else(false))
}
func (p ProcessStruct) GetNativeOs() GumOS {
	return GumOS(C.gum_process_get_native_os())
}
func (p ProcessStruct) GetMainModule() *ModuleDetails {
	ptr := C.gum_process_get_main_module()
	return ModuleDetailsWithPtr(unsafe.Pointer(ptr))
}

func (p ProcessStruct) EnumerateModules() []*ModuleDetails {
	details := make([]*ModuleDetails, 0)
	var cb = ffi.NewNativeCallback(func(ptr ffi.NativePointer, ptr2 ffi.NativePointer) bool {
		details = append(details, ModuleDetailsWithPtr(ptr.Ptr()))
		return true
	}, ffi.TBool, []ffi.ArgTypeName{ffi.TPointer, ffi.TPointer})
	defer cb.Free()
	C.gum_process_enumerate_modules((*[0]byte)(unsafe.Pointer(cb.Ptr())), (C.gpointer)(unsafe.Pointer(uintptr(0))))
	return details
}
func (p ProcessStruct) EnumerateThreads() []*ThreadDetails {
	details := make([]*ThreadDetails, 0)
	var cb = ffi.NewNativeCallback(func(ptr ffi.NativePointer, ptr2 ffi.NativePointer) bool {
		details = append(details, ThreadDetailsWithPtr(ptr.Ptr()))
		return true
	}, ffi.TBool, []ffi.ArgTypeName{ffi.TPointer, ffi.TPointer})
	defer cb.Free()
	C.gum_process_enumerate_threads((*[0]byte)(unsafe.Pointer(cb)), (C.gpointer)(unsafe.Pointer(uintptr(0))))
	return details
}
func (p ProcessStruct) EnumerateMallocRanges() []*RangeDetails {
	details := make([]*RangeDetails, 0)
	var cb = ffi.NewNativeCallback(func(ptr ffi.NativePointer, ptr2 ffi.NativePointer) int {
		details = append(details, RangeDetailsWithPtr(ptr.Ptr()))
		return 1
	}, ffi.Tint, []ffi.ArgTypeName{ffi.TPointer, ffi.TPointer})
	defer cb.Free()
	C.gum_process_enumerate_malloc_ranges((*[0]byte)(unsafe.Pointer(cb.Ptr())), (C.gpointer)(unsafe.Pointer(uintptr(0))))
	return details
}

func (p ProcessStruct) FindModuleByName(moduleName string) *ModuleDetails {
	modules := p.EnumerateModules()
	for _, module := range modules {
		fmt.Println(module.Name)
		if module.Name == moduleName {
			return module
		}
		if strings.TrimSuffix(moduleName, ".so") == moduleName {
			return module
		}
		matched, _ := regexp.MatchString(moduleName, module.Name)
		if matched {
			return module
		}
	}
	return nil
}
