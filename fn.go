package gumjs

import "C"
import "unsafe"

func GoString(p *C.char) string {
	return C.GoString(p)
}
func CString(b string) *C.char {
	temp := []byte(b)
	utf8StrArr := make([]uint8, len(temp)+1) // +1是因为Lazarus中PChar为0结尾
	copy(utf8StrArr, temp)
	return (*C.char)(unsafe.Pointer(&utf8StrArr[0]))
}

func CBytesToGoBytes(ustr unsafe.Pointer, n int) []byte {
	return copyBytes(ustr, n)
}
func copyBytes(src unsafe.Pointer, strLen int) []byte {
	if strLen == 0 {
		return nil
	}
	str := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		str[i] = *(*byte)(unsafe.Pointer(uintptr(src) + uintptr(i)))
	}
	return str
}
