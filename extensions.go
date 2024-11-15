package vk

// #include "vk.h"
// #include <stdlib.h>
import "C"
import "unsafe"

type Extension interface {
	isExtension()
	externalize() unsafe.Pointer
	internalize(unsafe.Pointer)
}

type structHeader struct {
	Type StructureType
	Next unsafe.Pointer
}

func buildChain(exs []Extension) unsafe.Pointer {
	var next unsafe.Pointer
	for i := len(exs) - 1; i >= 0; i-- {
		cex := exs[i].externalize()
		(*structHeader)(cex).Next = next
		next = cex
	}
	return next
}

func internalizeChain(exs []Extension, chain unsafe.Pointer) {
	if chain == nil {
		return
	}

	for _, ex := range exs {
		ex.internalize(chain)
		next := (*structHeader)(chain).Next
		C.free(chain)
		chain = next
	}
}

type PhysicalDeviceIDProperties struct {
	DeviceUUID      [C.VK_UUID_SIZE]byte
	DriverUUID      [C.VK_UUID_SIZE]byte
	DeviceLUID      [C.VK_UUID_SIZE]byte
	DeviceNodeMask  uint32
	DeviceLUIDValid bool
}

func (*PhysicalDeviceIDProperties) isExtension() {}

func (prop *PhysicalDeviceIDProperties) externalize() unsafe.Pointer {
	cprop := alloc[C.VkPhysicalDeviceIDProperties]()
	cprop.sType = C.VkStructureType(StructureTypePhysicalDeviceIdProperties)
	return unsafe.Pointer(cprop)
}

func (prop *PhysicalDeviceIDProperties) internalize(ptr unsafe.Pointer) {
	cprop := (*C.VkPhysicalDeviceIDProperties)(ptr)
	copy(prop.DeviceUUID[:], (*[C.VK_UUID_SIZE]byte)(unsafe.Pointer(&cprop.deviceUUID))[:])
	copy(prop.DriverUUID[:], (*[C.VK_UUID_SIZE]byte)(unsafe.Pointer(&cprop.driverUUID))[:])
	copy(prop.DeviceLUID[:], (*[C.VK_UUID_SIZE]byte)(unsafe.Pointer(&cprop.deviceLUID))[:])
	prop.DeviceNodeMask = uint32(cprop.deviceNodeMask)
	prop.DeviceLUIDValid = cprop.deviceLUIDValid == C.VK_TRUE
}
