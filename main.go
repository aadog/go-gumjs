package gumjs

import (
	"runtime"
)

func init() {
	runtime.LockOSThread()
}
func main() {
	InitEmbedded()
}
