package gumjs

/*
#include <frida-gumjs.h>
*/
import "C"
import "time"

func InitEmbedded() {
	C.gum_init_embedded()
	interceptorInitFunc()
}
func DeinitEmbedded() {
	C.gum_deinit_embedded()
}

func GLoop() {
	context := C.g_main_context_get_thread_default()
	for {
		if C.g_main_context_pending(context) == 0 {
			time.Sleep(time.Millisecond * 10)
		}
		C.g_main_context_iteration(context, C.FALSE)
	}
}
