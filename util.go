package vk

// #include <stdlib.h>
// #include "vk.h"
import "C"

import (
	"bytes"
	"math"
	"reflect"
	"unsafe"

	"honnef.co/go/safeish"
)

const alignment = 8

func align(p uintptr) uintptr {
	return (p + alignment - 1) &^ (alignment - 1)
}

// ucopy copies data from src to dst,
// where dst must be a C pointer and src must be a pointer to a Go slice.
func ucopy(dst, src unsafe.Pointer, size uintptr) {
	elems := (*reflect.SliceHeader)(src).Len
	if elems == 0 {
		return
	}
	// Access the slice's underlying data
	src = unsafe.Pointer((*reflect.SliceHeader)(src).Data)
	copy(
		(*[math.MaxInt32]byte)(dst)[:uintptr(elems)*size],
		(*[math.MaxInt32]byte)(src)[:uintptr(elems)*size],
	)
}

func externString(a *allocator, s string) *C.char {
	m := C.CString(s)
	a.allocs = append(a.allocs, unsafe.Pointer(m))
	return m
}

func externStrings(a *allocator, ss []string) **C.char {
	size0 := C.sizeof_uintptr_t * uintptr(len(ss))
	var size1 uintptr
	for _, s := range ss {
		size1 += uintptr(len(s)) + 1
	}
	size1 = align(size1)
	size := size0 + size1
	mem := C.calloc(1, C.size_t(size))
	a.allocs = append(a.allocs, mem)
	arr := (*[math.MaxInt32]unsafe.Pointer)(mem)[:len(ss)]
	data := unsafe.Add(mem, size0)
	for i, s := range ss {
		arr[i] = data
		ucopy(data, unsafe.Pointer(&s), 1)
		data = unsafe.Add(data, len(s)+1)
	}
	return (**C.char)(mem)
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
