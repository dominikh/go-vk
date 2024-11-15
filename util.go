package vk

// #include <stdlib.h>
// #include "vk.h"
import "C"

import (
	"bytes"
	"unsafe"

	"honnef.co/go/safeish"
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

func externFloat32(a *allocator, vs []float32) *C.float {
	if len(vs) == 0 {
		return nil
	}
	ptr := allocn[C.float](a, len(vs))
	s := unsafe.Slice(ptr, len(vs))
	copy(s, safeish.SliceCast[[]C.float](vs))
	return ptr
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
