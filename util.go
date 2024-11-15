package vk

// #include <stdlib.h>
// #include "vk.h"
import "C"

import (
	"bytes"
	"unsafe"
)

func externString(a *allocator, s string) *C.char {
	m := C.CString(s)
	a.allocs = append(a.allocs, unsafe.Pointer(m))
	return m
}

func externStrings(a *allocator, ss []string) **C.char {
	cstrings := make([]*C.char, len(ss))
	for i, s := range ss {
		cstrings[i] = externString(a, s)
	}
	return pinSlice(a, cstrings)
}

func result2error(res Result) error {
	if res == Success {
		return nil
	}
	return res
}

func vkBool(b bool) C.VkBool32 {
	if b {
		return C.VK_TRUE
	}
	return C.VK_FALSE
}

func str(x []C.char) string {
	v := *(*[]byte)(unsafe.Pointer(&x))
	idx := bytes.IndexByte(v, 0)
	return string(v[:idx])
}
