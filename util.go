package vk

// #include <stdlib.h>
import "C"

import (
	"math"
	"reflect"
	"unsafe"
)

func alloc(size C.size_t) unsafe.Pointer {
	return C.calloc(1, size)
}

func allocn(nmemb int, size C.size_t) unsafe.Pointer {
	if nmemb == 0 {
		return nil
	}
	return C.calloc(C.size_t(nmemb), size)
}

func free(ptr unsafe.Pointer) {
	C.free(ptr)
}

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

func ucopy1(dst, src unsafe.Pointer, size uintptr) {
	copy(
		(*[math.MaxInt32]byte)(dst)[:size],
		(*[math.MaxInt32]byte)(src)[:size],
	)
}

func externStrings(ss []string) (**C.char, func()) {
	if len(ss) == 0 {
		return nil, func() {}
	}
	var ptrs []unsafe.Pointer

	ptr := allocn(len(ss), C.size_t(unsafe.Sizeof(uintptr(0))))
	ptrs = append(ptrs, ptr)
	slice := (*[math.MaxInt32]*C.char)(ptr)[:len(ss)]
	for i, s := range ss {
		slice[i] = C.CString(s)
		ptrs = append(ptrs, unsafe.Pointer(slice[i]))
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
	return (*C.float)(C.CBytes((*[math.MaxInt32]byte)(unsafe.Pointer(&vs[0]))[:uintptr(len(vs))*unsafe.Sizeof(float32(0))]))
}

func result2error(res Result) error {
	if res == Success {
		return nil
	}
	return res
}

func slice2ptr(slice unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer((*reflect.SliceHeader)(slice).Data)
}
