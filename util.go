package vk

// #include <stdlib.h>
// #include "vk.h"
import "C"

import (
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

func externStrings(ss []string) (**C.char, func()) {
	if len(ss) == 0 {
		return nil, func() {}
	}
	var ptrs []uptr

	ptr := allocn(len(ss), C.size_t(unsafe.Sizeof(uintptr(0))))
	ptrs = append(ptrs, ptr)
	slice := (*[math.MaxInt32]*C.char)(ptr)[:len(ss)]
	for i, s := range ss {
		slice[i] = C.CString(s)
		ptrs = append(ptrs, uptr(slice[i]))
	}
	return (**C.char)(ptr), func() {
		for _, ptr := range ptrs {
			free(ptr)
		}
	}
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
