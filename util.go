package vk

// #include <stdlib.h>
import "C"

import (
	"math"
	"unsafe"
)

func calloc(nmemb C.size_t, size C.size_t) unsafe.Pointer {
	if nmemb == 0 {
		return nil
	}
	return C.calloc(nmemb, size)
}

// ucopy copies data from src to dst,
// where dst must be a C pointer and src must be a pointer to a Go slice.
func ucopy(dst, src unsafe.Pointer, size uintptr) {
	elems := *(*int)(unsafe.Pointer(uintptr(src) + unsafe.Sizeof(uintptr(0))))
	if elems == 0 {
		return
	}
	// Access the slice's underlying data
	src = (*(*unsafe.Pointer)(src))
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
