package vk

// #include "vk.h"
import "C"

type Extension interface {
	isExtension()
	externalize() uptr
	internalize(uptr)
}

type structHeader struct {
	Type StructureType
	Next uptr
}

func buildChain(exs []Extension) uptr {
	if len(exs) == 0 {
		return nil
	}
	out := make([]uptr, len(exs))
	for i, ex := range exs {
		out[i] = ex.externalize()
	}
	for i, ptr := range out[:len(out)-1] {
		(*structHeader)(ptr).Next = out[i+1]
	}
	return out[0]
}

func internalizeChain(exs []Extension, chain uptr) {
	if chain == nil {
		return
	}

	for _, ex := range exs {
		ex.internalize(chain)
		next := (*structHeader)(chain).Next
		free(chain)
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

func (prop *PhysicalDeviceIDProperties) externalize() uptr {
	cprop := (*C.VkPhysicalDeviceIDProperties)(alloc(C.sizeof_VkPhysicalDeviceIDProperties))
	cprop.sType = C.VkStructureType(StructureTypePhysicalDeviceIdProperties)
	return uptr(cprop)
}

func (prop *PhysicalDeviceIDProperties) internalize(ptr uptr) {
	cprop := (*C.VkPhysicalDeviceIDProperties)(ptr)
	copy(prop.DeviceUUID[:], (*[C.VK_UUID_SIZE]byte)(uptr(&cprop.deviceUUID))[:])
	copy(prop.DriverUUID[:], (*[C.VK_UUID_SIZE]byte)(uptr(&cprop.driverUUID))[:])
	copy(prop.DeviceLUID[:], (*[C.VK_UUID_SIZE]byte)(uptr(&cprop.deviceLUID))[:])
	prop.DeviceNodeMask = uint32(cprop.deviceNodeMask)
	prop.DeviceLUIDValid = cprop.deviceLUIDValid == C.VK_TRUE
}
