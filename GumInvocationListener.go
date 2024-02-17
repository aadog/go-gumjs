package gum

/*
#include "invocationlistener.h"
extern void InvocationListenerCallbackOnEnter(void*,void*);
extern void InvocationListenerCallbackOnLeave(void*,void*);
*/
import "C"
import (
	"github.com/google/uuid"
	"sync"
	"unsafe"
)

type GumInvocationListener struct {
	ptr unsafe.Pointer
	Id  string
}

var mpGumInvocationListenerCallback sync.Map

//export InvocationListenerCallbackOnEnter
func InvocationListenerCallbackOnEnter(listener unsafe.Pointer, context unsafe.Pointer) {
	cid := (*C.GumInvocationListenerProxy)(listener).Id
	id := GoString((*C.char)(unsafe.Pointer(&cid)))
	iproxy, ok := mpGumInvocationListenerCallback.Load(id)
	if !ok {
		return
	}

	switch proxy := iproxy.(type) {
	case *InvocationListenerCallbacks:
		if proxy.OnEnter != nil {
			proxy.GetOnEnter()(InvocationContextWithPtr(context))
		}
	case InstructionProbeCallback:
		if proxy.GetOnEnter() != nil {
			proxy.GetOnEnter()(InvocationContextWithPtr(context))
		}
	}
}

//export InvocationListenerCallbackOnLeave
func InvocationListenerCallbackOnLeave(listener unsafe.Pointer, context unsafe.Pointer) {
	cid := (*C.GumInvocationListenerProxy)(listener).Id
	id := GoString((*C.char)(unsafe.Pointer(&cid)))
	iproxy, ok := mpGumInvocationListenerCallback.Load(id)
	if !ok {
		return
	}
	switch proxy := iproxy.(type) {
	case *InvocationListenerCallbacks:
		if proxy.GetOnLeave() != nil {
			proxy.GetOnLeave()(InvocationContextWithPtr(context))
		}
	case InstructionProbeCallback:
		if proxy.GetOnLeave() != nil {
			proxy.GetOnLeave()(InvocationContextWithPtr(context))
		}
	}
}
func init() {
	C.SetGumInvocationProxyCallback(C.InvocationListenerCallbackOnEnter, C.InvocationListenerCallbackOnLeave)
}

func (g *GumInvocationListener) Unref() {
	C.g_object_unref((C.gpointer)(g.ptr))
	mpGumInvocationListenerCallback.Delete(g.Id)
}

func NewListener(proxy IListenerCallback) *GumInvocationListener {
	id := uuid.New().String()
	mpGumInvocationListenerCallback.Store(id, proxy)
	return &GumInvocationListener{
		ptr: unsafe.Pointer((*C.GumInvocationListener)(C.CreateListener(CString(id)))),
		Id:  id,
	}
}
