package vk

// #include <stdlib.h>
import "C"
import (
	"runtime"
	"unsafe"

	"honnef.co/go/safeish"
)

type allocator struct {
	allocs []unsafe.Pointer
	pinner runtime.Pinner
}

func (a *allocator) free() {
	for _, alloc := range a.allocs {
		C.free(alloc)
	}
	a.allocs = nil
	a.pinner.Unpin()
}

func pinAsCastedPtr[Dst ~*DstE, Src ~*SrcE, DstE, SrcE any](a *allocator, s Src) Dst {
	ptr := safeish.Cast[Dst](s)
	a.pinner.Pin(ptr)
	return ptr
}

func pinSliceAsCastedPtr[Dst ~*DstE, Src ~[]SrcE, DstE, SrcE any](a *allocator, s Src) Dst {
	ptr := safeish.SliceCastPtr[Dst](s)
	a.pinner.Pin(ptr)
	return ptr
}

func alloc[T any](a *allocator) *T {
	return allocn[T](a, 1)
}

func allocn[T any](a *allocator, nmemb int) *T {
	if nmemb == 0 {
		return nil
	}
	m := (*T)(C.calloc(C.size_t(nmemb), C.size_t(unsafe.Sizeof(*new(T)))))
	a.allocs = append(a.allocs, unsafe.Pointer(m))
	return m
}
