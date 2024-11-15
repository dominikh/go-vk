package vk

// #include <stdlib.h>
import "C"
import "unsafe"

type allocator struct {
	allocs []unsafe.Pointer
}

func (a *allocator) free() {
	for _, alloc := range a.allocs {
		C.free(alloc)
	}
	a.allocs = nil
}

func allocRaw(a *allocator, size uintptr) unsafe.Pointer {
	m := C.calloc(1, C.size_t(size))
	a.allocs = append(a.allocs, m)
	return m
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
