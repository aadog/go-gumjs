package gumjs

/*
#include <frida-gumjs.h>
*/
import "C"
import (
	"github.com/samber/mo"
	"unsafe"
)

type ScriptBackEnd struct {
	ptr unsafe.Pointer
}

func (s *ScriptBackEnd) CreateSync(name string, script string) mo.Result[*Script] {
	var gerr *C.GError
	sc := C.gum_script_backend_create_sync((*C.GumScriptBackend)(s.ptr), CString(name), CString(script), nil, nil, &gerr)
	if gerr != nil {
		return mo.Err[*Script](ConvertGErrorAndFree(unsafe.Pointer(gerr)))
	}
	return mo.Ok(&Script{ptr: unsafe.Pointer(sc)})
}

func QjsBackEnd() *ScriptBackEnd {
	return &ScriptBackEnd{
		ptr: unsafe.Pointer(C.gum_script_backend_obtain_v8()),
	}
}
