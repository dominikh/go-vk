package vk

// #include <stdlib.h>
// #include "vk.h"
import "C"

import (
	"bytes"
	"math"
	"reflect"
	"unsafe"
)

type uptr = unsafe.Pointer

func alloc(size C.size_t) uptr {
	return C.calloc(1, size)
}

func allocn(nmemb int, size C.size_t) uptr {
	if nmemb == 0 {
		return nil
	}
	return C.calloc(C.size_t(nmemb), size)
}

func free(ptr uptr) {
	C.free(ptr)
}

const alignment = 8

func align(p uintptr) uintptr {
	return (p + alignment - 1) &^ (alignment - 1)
}

// ucopy copies data from src to dst,
// where dst must be a C pointer and src must be a pointer to a Go slice.
func ucopy(dst, src uptr, size uintptr) {
	elems := (*reflect.SliceHeader)(src).Len
	if elems == 0 {
		return
	}
	// Access the slice's underlying data
	src = uptr((*reflect.SliceHeader)(src).Data)
	copy(
		(*[math.MaxInt32]byte)(dst)[:uintptr(elems)*size],
		(*[math.MaxInt32]byte)(src)[:uintptr(elems)*size],
	)
}

func ucopy1(dst, src uptr, size uintptr) {
	copy(
		(*[math.MaxInt32]byte)(dst)[:size],
		(*[math.MaxInt32]byte)(src)[:size],
	)
}

func externStrings(ss []string) **C.char {
	size0 := C.sizeof_uintptr_t * uintptr(len(ss))
	var size1 uintptr
	for _, s := range ss {
		size1 += uintptr(len(s)) + 1
	}
	size1 = align(size1)
	size := size0 + size1
	mem := alloc(C.size_t(size))
	arr := (*[math.MaxInt32]uptr)(mem)[:len(ss)]
	data := uptr(uintptr(mem) + size0)
	for i, s := range ss {
		arr[i] = data
		ucopy(data, unsafe.Pointer(&s), 1)
		data = uptr(uintptr(data) + uintptr(len(s)) + 1)
	}
	return (**C.char)(mem)
}

func externFloat32(vs []float32) *C.float {
	if len(vs) == 0 {
		return nil
	}
	return (*C.float)(C.CBytes((*[math.MaxInt32]byte)(uptr(&vs[0]))[:uintptr(len(vs))*unsafe.Sizeof(float32(0))]))
}

func result2error(res Result) error {
	if res == Success {
		return nil
	}
	return res
}

func slice2ptr(slice uptr) uptr {
	return uptr((*reflect.SliceHeader)(slice).Data)
}

func vkBool(b bool) C.VkBool32 {
	if b {
		return C.VK_TRUE
	}
	return C.VK_FALSE
}

func str(x []C.char) string {
	v := *(*[]byte)(uptr(&x))
	idx := bytes.IndexByte(v, 0)
	return string(v[:idx])
}
