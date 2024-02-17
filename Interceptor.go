package gum

/*
#include <frida-gumjs.h>
*/
import "C"
import (
	"github.com/aadog/go-ffi"
	"sync"
	"unsafe"
)

var Interceptor InterceptorStruct

var interceptorPtr unsafe.Pointer
var interceptorInitFunc = sync.OnceFunc(func() {
	interceptorPtr = unsafe.Pointer(C.gum_interceptor_obtain())
})

type InterceptorStruct struct {
}

func (g InterceptorStruct) Replace(functionAddress ffi.NativePointer, replacementFunction ffi.NativePointer, replacementData ffi.NativePointer) {

	var originalFunction unsafe.Pointer
	C.gum_interceptor_replace(
		(*C.GumInterceptor)(interceptorPtr),
		(C.gpointer)(functionAddress.Ptr()),
		(C.gpointer)(replacementFunction.Ptr()),
		(C.gpointer)(replacementData.Ptr()),
		(*C.gpointer)(&originalFunction),
	)
}
func (i InterceptorStruct) Flush() {
	C.gum_interceptor_flush((*C.GumInterceptor)(interceptorPtr))
}
func (i InterceptorStruct) Attach(functionAddress ffi.NativePointer, listenerCallback IListenerCallback) *GumInvocationListener {
	l := NewListener(listenerCallback)
	C.gum_interceptor_attach(
		(*C.GumInterceptor)(interceptorPtr),
		(C.gpointer)(functionAddress.Ptr()),
		(*C.GumInvocationListener)(l.ptr),
		(C.gpointer)(uintptr(0)),
	)
	return l
}
