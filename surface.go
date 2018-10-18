// Copyright (c) 2015-2018 Khronos Group
// Copyright (c) 2018 Dominik Honnef

package vk

// #cgo pkg-config: vulkan
// #cgo CFLAGS: -DVK_NO_PROTOTYPES
// #include <vulkan/vulkan_core.h>
//
// VkResult domVkGetPhysicalDeviceSurfaceSupportKHR(PFN_vkGetPhysicalDeviceSurfaceSupportKHR fp, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, VkSurfaceKHR surface, VkBool32* pSupported);
// VkResult domVkGetPhysicalDeviceSurfaceCapabilitiesKHR(PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR fp, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, VkSurfaceCapabilitiesKHR* pSurfaceCapabilities);
// VkResult domVkGetPhysicalDeviceSurfaceFormatsKHR(PFN_vkGetPhysicalDeviceSurfaceFormatsKHR fp, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pSurfaceFormatCount, VkSurfaceFormatKHR* pSurfaceFormats);
// VkResult domVkGetPhysicalDeviceSurfacePresentModesKHR(PFN_vkGetPhysicalDeviceSurfacePresentModesKHR fp, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pPresentModeCount, VkPresentModeKHR* pPresentModes);
import "C"
import "unsafe"

func init() {
	assertSameSize(unsafe.Sizeof(SurfaceCapabilities{}), unsafe.Sizeof(C.VkSurfaceCapabilitiesKHR{}))
	assertSameSize(unsafe.Sizeof(SurfaceFormatKHR{}), unsafe.Sizeof(C.VkSurfaceFormatKHR{}))
}

type SurfaceKHR struct {
	hnd C.VkSurfaceKHR
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
}

type SurfaceFormatKHR struct {
	Format     Format
	ColorSpace ColorSpaceKHR
}

func (dev *PhysicalDevice) SurfaceSupportKHR(queueFamilyIndex uint32, surface *SurfaceKHR) (bool, error) {
	var out C.VkBool32
	res := Result(C.domVkGetPhysicalDeviceSurfaceSupportKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceSupportKHR], dev.hnd, C.uint32_t(queueFamilyIndex), surface.hnd, &out))
	if res != Success {
		return false, res
	}
	return out == C.VK_TRUE, nil
}

func (dev *PhysicalDevice) SurfaceCapabilitiesKHR(surface *SurfaceKHR) (*SurfaceCapabilities, error) {
	var out SurfaceCapabilities
	res := Result(C.domVkGetPhysicalDeviceSurfaceCapabilitiesKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceCapabilitiesKHR], dev.hnd, surface.hnd, (*C.VkSurfaceCapabilitiesKHR)(unsafe.Pointer(&out))))
	if res != Success {
		return nil, res
	}
	return &out, nil
}

func (dev *PhysicalDevice) SurfaceFormatsKHR(surface *SurfaceKHR) ([]SurfaceFormatKHR, error) {
	var count C.uint32_t
	res := Result(C.domVkGetPhysicalDeviceSurfaceFormatsKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceFormatsKHR], dev.hnd, surface.hnd, &count, nil))
	if res != Success {
		return nil, res
	}
	out := make([]SurfaceFormatKHR, count)
	res = Result(C.domVkGetPhysicalDeviceSurfaceFormatsKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceFormatsKHR], dev.hnd, surface.hnd, &count, (*C.VkSurfaceFormatKHR)(unsafe.Pointer(&out[0]))))
	if res != Success {
		return nil, res
	}
	return out, nil
}

func (dev *PhysicalDevice) SurfacePresentModesKHR(surface *SurfaceKHR) ([]PresentModeKHR, error) {
	var count C.uint32_t
	res := Result(C.domVkGetPhysicalDeviceSurfacePresentModesKHR(dev.instance.fps[vkGetPhysicalDeviceSurfacePresentModesKHR], dev.hnd, surface.hnd, &count, nil))
	if res != Success {
		return nil, res
	}
	out := make([]PresentModeKHR, count)
	res = Result(C.domVkGetPhysicalDeviceSurfacePresentModesKHR(dev.instance.fps[vkGetPhysicalDeviceSurfacePresentModesKHR], dev.hnd, surface.hnd, &count, (*C.VkPresentModeKHR)(unsafe.Pointer(&out[0]))))
	if res != Success {
		return nil, res
	}
	return out, nil
}
