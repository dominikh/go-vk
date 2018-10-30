// Copyright (c) 2015-2018 Khronos Group
// Copyright (c) 2018 Dominik Honnef

package vk

// #cgo pkg-config: vulkan
// #include "vk.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func init() {
	assertSameSize(unsafe.Sizeof(SurfaceCapabilities{}), unsafe.Sizeof(C.VkSurfaceCapabilitiesKHR{}))
	assertSameSize(unsafe.Sizeof(SurfaceFormatKHR{}), unsafe.Sizeof(C.VkSurfaceFormatKHR{}))
}

type SurfaceKHR struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSurfaceKHR)
	hnd C.VkSurfaceKHR
}

func (hnd SurfaceKHR) String() string {
	return fmt.Sprintf("VkSurfaceKHR(%#x)", hnd.hnd)
}

type SurfaceCapabilities struct {
	MinImageCount           uint32
	MaxImageCount           uint32
	CurrentExtent           Extent2D
	MinImageExtent          Extent2D
	MaxImageExtent          Extent2D
	MaxImageArrayLayers     uint32
	SupportedTransforms     SurfaceTransformFlagsKHR
	CurrentTransform        SurfaceTransformFlagsKHR
	SupportedCompositeAlpha CompositeAlphaFlagsKHR
	SupportedUsageFlags     ImageUsageFlags

	// must be kept identical to C struct
}

type SurfaceFormatKHR struct {
	Format     Format
	ColorSpace ColorSpaceKHR

	// must be kept identical to C struct
}

func (dev *PhysicalDevice) SurfaceSupportKHR(queueFamilyIndex uint32, surface SurfaceKHR) (bool, error) {
	var out C.VkBool32
	res := Result(C.domVkGetPhysicalDeviceSurfaceSupportKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceSupportKHR], dev.hnd, C.uint32_t(queueFamilyIndex), surface.hnd, &out))
	return out == C.VK_TRUE, result2error(res)
}

func (dev *PhysicalDevice) SurfaceCapabilitiesKHR(surface SurfaceKHR) (*SurfaceCapabilities, error) {
	var out SurfaceCapabilities
	res := Result(C.domVkGetPhysicalDeviceSurfaceCapabilitiesKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceCapabilitiesKHR], dev.hnd, surface.hnd, (*C.VkSurfaceCapabilitiesKHR)(uptr(&out))))
	return &out, result2error(res)
}

func (dev *PhysicalDevice) SurfaceFormatsKHR(surface SurfaceKHR) ([]SurfaceFormatKHR, error) {
	var count C.uint32_t
	res := Result(C.domVkGetPhysicalDeviceSurfaceFormatsKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceFormatsKHR], dev.hnd, surface.hnd, &count, nil))
	if res != Success {
		return nil, res
	}
	out := make([]SurfaceFormatKHR, count)
	res = Result(C.domVkGetPhysicalDeviceSurfaceFormatsKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceFormatsKHR], dev.hnd, surface.hnd, &count, (*C.VkSurfaceFormatKHR)(uptr(&out[0]))))
	return out, result2error(res)
}

func (dev *PhysicalDevice) SurfacePresentModesKHR(surface SurfaceKHR) ([]PresentModeKHR, error) {
	var count C.uint32_t
	res := Result(C.domVkGetPhysicalDeviceSurfacePresentModesKHR(dev.instance.fps[vkGetPhysicalDeviceSurfacePresentModesKHR], dev.hnd, surface.hnd, &count, nil))
	if res != Success {
		return nil, res
	}
	out := make([]PresentModeKHR, count)
	res = Result(C.domVkGetPhysicalDeviceSurfacePresentModesKHR(dev.instance.fps[vkGetPhysicalDeviceSurfacePresentModesKHR], dev.hnd, surface.hnd, &count, (*C.VkPresentModeKHR)(uptr(&out[0]))))
	return out, result2error(res)
}
