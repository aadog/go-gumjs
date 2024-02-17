package gum

/*
#include <frida-gumjs.h>
*/
import "C"

func InitEmbedded() {
	C.gum_init_embedded()
	interceptorInitFunc()
}
func DeinitEmbedded() {
	C.gum_deinit_embedded()
}
