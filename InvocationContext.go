package gum

/*
#include <frida-gumjs.h>
*/
import "C"
import (
	"github.com/aadog/go-ffi"
	"unsafe"
)

type InvocationContext struct {
	ptr unsafe.Pointer
}

func (g *InvocationContext) GetNthArgumentPtr(n uint) ffi.NativePointer {
	return ffi.Ptr(unsafe.Pointer(C.gum_invocation_context_get_nth_argument((*C.GumInvocationContext)(g.ptr), C.uint(n))))
}

func (g *InvocationContext) GetFunction() ffi.NativePointer {
	return ffi.Ptr(uintptr((*C.GumInvocationContext)(g.ptr).function))
}

func (g *InvocationContext) ReplaceNthArgumentPtr(n uint, value ffi.NativePointer) {
	C.gum_invocation_context_replace_nth_argument((*C.GumInvocationContext)(g.ptr), C.uint(n), (C.gpointer)(value.Ptr()))
}

func (g *InvocationContext) GetReturnValuePtr() ffi.NativePointer {
	return ffi.Ptr(uintptr(C.gum_invocation_context_get_return_value((*C.GumInvocationContext)(g.ptr))))
}

func (g *InvocationContext) GetThreadId() uint {
	return uint(C.gum_invocation_context_get_thread_id((*C.GumInvocationContext)(g.ptr)))
}
func (g *InvocationContext) GetListenerThreadDataPtr(size int64) ffi.NativePointer {
	return ffi.Ptr(uintptr(C.gum_invocation_context_get_listener_thread_data((*C.GumInvocationContext)(g.ptr), CULong(uint64(size)))))
}

func (g *InvocationContext) GetListenerFunctionDataPtr() ffi.NativePointer {
	return ffi.Ptr(uintptr(C.gum_invocation_context_get_listener_function_data((*C.GumInvocationContext)(g.ptr))))
}
func (g *InvocationContext) GetListenerInvocationDataPtr(size int64) ffi.NativePointer {
	return ffi.Ptr(uintptr(C.gum_invocation_context_get_listener_invocation_data((*C.GumInvocationContext)(g.ptr), CULong(uint64(size)))))
}

func (g *InvocationContext) GetReplacementDataPtr() ffi.NativePointer {
	return ffi.Ptr(uintptr(C.gum_invocation_context_get_replacement_data((*C.GumInvocationContext)(g.ptr))))
}

func (g *InvocationContext) GetCpuContext() ffi.NativePointer {
	return ffi.Ptr(uintptr(unsafe.Pointer((*C.GumInvocationContext)(g.ptr).cpu_context)))
}

func InvocationContextWithPtr(ptr unsafe.Pointer) *InvocationContext {
	return &InvocationContext{ptr: ptr}
}
