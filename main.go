package gumjs

/*
#include "go-gum.h"
*/
import "C"
import (
	"fmt"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	InitEmbedded()
	//C.testC()
	//defer DeinitEmbedded()

	Interceptor.Attach(Module.FindExportByName(nil, "open"), &InvocationListenerCallbacks{
		OnEnter: func(context *InvocationContext) {
			//fmt.Println("call onenter:", context.GetNthArgumentPtr(0).ReadCString(-1))
		},
		OnLeave: func(context *InvocationContext) {
			fmt.Println("call onleave:", context.GetNthArgumentPtr(0).ReadCString(-1))
		},
	})

	//threads := Process.EnumerateThreads()
	//for _, details := range threads {
	//	fmt.Println(fmt.Sprintf("thread:%d:%v", details.Id, details.State))
	//}
	//fmt.Println("xxx")
	fmt.Println(Module.ModuleLoad("libdl.so").MustGet())
	fmt.Println(fmt.Sprintf("os:%d", Process.GetNativeOs()))
	fmt.Println(fmt.Sprintf("IsDebuggerAttached:%v", Process.IsDebuggerAttached()))
	fmt.Println(fmt.Sprintf("processId:%v", Process.Id()))
	fmt.Println(fmt.Sprintf("threadId:%v", Process.GetCurrentThreadId()))
	mainModule := Process.FindModuleByName("libpthread-2.31.so")
	fmt.Println(mainModule)
	//fmt.Println(fmt.Sprintf("mainModule:%s path:%s base:%v size:%v", mainModule.Name, mainModule.Path, mainModule.Base, mainModule.Size))
	//fmt.Println(fmt.Sprintf("getBase:%v", Module.GetBaseAddress("libdl.so")))
	//fmt.Println(fmt.Sprintf("getBase:%v", Module.GetBaseAddress("libdl.so")))
	//for _, module := range Process.EnumerateModules() {
	//	fmt.Println(module.Path)
	//}

	//for _, details := range Process.EnumerateMallocRanges() {
	//	fmt.Println(fmt.Sprintf("thread:%d:%d ,%v", details.Base, details.Size, details.Protection))
	//}

	//for _, details := range Process.EnumerateModules() {
	//	fmt.Println(details.Name)
	//}
	//for _, details := range Process.EnumerateMallocRanges() {
	//	fmt.Println(details.Base)
	//}
	//
	//for _, details := range mainModule.EnumerateSymbols() {
	//	fmt.Println(fmt.Sprintf("symbol:%s", details.Name))
	//}

}
