package gumjs

/*
#cgo LDFLAGS: -lfrida-gumjs -ldl -lm -lstdc++
#cgo CXXFLAGS: -std=c++11
#cgo CFLAGS: -I/usr/local/include/ -w
#cgo darwin LDFLAGS: -lbsm -framework Foundation -framework AppKit -lresolv -lpthread
#cgo android CFLAGS: -DANDROID -Wno-error=incompatible-function-pointer-types
#cgo linux,!android LDFLAGS: -lpthread
#cgo linux CFLAGS: -pthread
#include <frida-gumjs.h>
*/
import "C"
