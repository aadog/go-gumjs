package gumjs

/*
#include <frida-gumjs.h>
extern void GoOnMessageCallBack(void* message,void* data,void* user_data);
*/
import "C"
import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"sync"
	"unsafe"
)

type OnMessageCallback func(message string, data []byte)

var onMessageCallBacks sync.Map
var ErrContextCancelled = errors.New("context cancelled")
var rpcCalls = sync.Map{}

//export GoOnMessageCallBack
func GoOnMessageCallBack(message *C.void, data *C.void, user_data *C.void) {
	var d []byte
	if data != nil {
		outlen := C.ulong(0)
		r := C.g_bytes_get_data((*C.GBytes)(unsafe.Pointer(data)), &outlen)
		if outlen > 0 {
			d = CBytesToGoBytes(unsafe.Pointer(r), int(outlen))
		}
	}
	fn, ld := onMessageCallBacks.Load(uintptr(unsafe.Pointer(user_data)))
	if ld {
		gomessage := GoString((*C.char)(unsafe.Pointer(message)))
		godata := d
		if strings.Contains(gomessage, "[\"frida:rpc\"") && strings.Contains(gomessage, "\"type\":\"send\"") {
			rpcID, ret, err := getRPCIDFromMessage(gomessage)
			if err != nil {
				panic(err)
			}
			callerCh, ok := rpcCalls.Load(rpcID)
			if !ok {
				panic("rpc-id not found")
			}
			ch := callerCh.(chan any)
			ch <- ret
			close(ch)
		} else {
			fn.(OnMessageCallback)(gomessage, godata)
		}
	}
}

type Script struct {
	ptr        unsafe.Pointer
	hasHandler bool
}

func (s *Script) On(fn OnMessageCallback) {
	s.hasHandler = true
	// hijack message to handle rpc calls

	addr := reflect.ValueOf(fn).Pointer()
	onMessageCallBacks.Store(addr, fn)

	C.gum_script_set_message_handler((*C.GumScript)(s.ptr), (*[0]byte)(unsafe.Pointer(C.GoOnMessageCallBack)), (C.gpointer)(unsafe.Pointer(addr)), (*[0]byte)(unsafe.Pointer(uintptr(0))))
}
func (s *Script) Load() {
	if !s.hasHandler {
		s.On(func(message string, data []byte) {

		})
	}
	C.gum_script_load_sync((*C.GumScript)(s.ptr), nil)
}

func (s *Script) Unload() {
	C.gum_script_unload_sync((*C.GumScript)(s.ptr), nil)
}
func (s *Script) Unref() {
	C.g_object_unref((C.gpointer)(s.ptr))
}

func getRPCIDFromMessage(message string) (string, any, error) {
	unmarshalled := make(map[string]any)
	if err := json.Unmarshal([]byte(message), &unmarshalled); err != nil {
		return "", nil, err
	}

	var rpcID string
	var ret any

	loopMap := func(mp map[string]any) {
		for _, v := range mp {
			if reflect.ValueOf(v).Kind() == reflect.Slice {
				slc := v.([]any)
				rpcID = slc[1].(string)
				ret = slc[3]

			}
		}
	}
	loopMap(unmarshalled)

	return rpcID, ret, nil
}
