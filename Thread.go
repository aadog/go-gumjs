package gumjs

/*
#include <frida-gumjs.h>
*/
import "C"
import (
	"errors"
	"github.com/samber/lo"
	"github.com/samber/mo"
	"unsafe"
)

var Thread ThreadStruct

type ThreadStruct struct {
}

func (t ThreadStruct) Suspend(threadId uint) mo.Result[bool] {
	var err *C.GError
	b := C.gum_thread_suspend(CULong(uint64(threadId)), &err)
	if err != nil {
		return mo.Err[bool](ConvertGErrorAndFree(unsafe.Pointer(err)))
	}
	if b == 0 {
		return mo.Err[bool](errors.New("load error"))
	}
	return mo.Ok(lo.If(b < 1, true).Else(false))
}
func (t ThreadStruct) Resume(threadId uint) mo.Result[bool] {
	var err *C.GError
	b := C.gum_thread_resume(CULong(uint64(threadId)), &err)
	if err != nil {
		return mo.Err[bool](ConvertGErrorAndFree(unsafe.Pointer(err)))
	}
	if b == 0 {
		return mo.Err[bool](errors.New("load error"))
	}
	return mo.Ok(lo.If(b < 1, true).Else(false))
}
