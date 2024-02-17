package gum

/*
#include <frida-gumjs.h>
*/
import "C"
import (
	"unsafe"
)

type GError struct {
	Domain  uint32
	Code    int
	Message string
}

func (g *GError) Error() string {
	return g.Message
}

func ConvertGError(ptr unsafe.Pointer) *GError {
	if ptr == nil {
		return nil
	}
	cObj := (*C.GError)(ptr)
	gerr := &GError{
		Domain:  uint32(cObj.domain),
		Code:    int(cObj.code),
		Message: GoString(cObj.message),
	}
	return gerr
}

func ConvertGErrorAndFree(ptr unsafe.Pointer) *GError {
	gErr := ConvertGError(ptr)
	C.g_error_free((*C.GError)(ptr))
	return gErr
}
