package util

import (
	"reflect"
	"unsafe"
)

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b)) // #nosec
}

func StringToByte(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s)) // #nosec
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))  // #nosec
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return
}
