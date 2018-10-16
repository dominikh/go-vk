package vk

import (
	"fmt"
)

func assertSameSize(a, b uintptr) {
	if a != b {
		panic(fmt.Sprintf("types aren't of identical size: %d != %d", a, b))
	}
}
