package gumjs

/*
#include <frida-gumjs.h>
*/
import "C"
import (
	"unsafe"
)

type ThreadDetails struct {
	Id         uint
	State      GumThreadState
	CpuContext ICpuContext
}

func ThreadDetailsWithPtr(ptr unsafe.Pointer) *ThreadDetails {
	th := (*C.GumThreadDetails)(ptr)

	return &ThreadDetails{
		Id:         uint(th.id),
		State:      GumThreadState(th.state),
		CpuContext: ICpuContextWithPtr(th.cpu_context),
	}
}
