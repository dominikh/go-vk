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

type Surface struct {
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

func (dev *PhysicalDevice) SurfaceSupportKHR(queueFamilyIndex uint32, surface *Surface) (bool, Result) {
	var out C.VkBool32
	res := C.domVkGetPhysicalDeviceSurfaceSupportKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceSupportKHR], dev.hnd, C.uint32_t(queueFamilyIndex), surface.hnd, &out)
	return out == C.VK_TRUE, Result(res)
}

func (dev *PhysicalDevice) SurfaceCapabilitiesKHR(surface *Surface) (*SurfaceCapabilities, Result) {
	var out SurfaceCapabilities
	res := C.domVkGetPhysicalDeviceSurfaceCapabilitiesKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceCapabilitiesKHR], dev.hnd, surface.hnd, (*C.VkSurfaceCapabilitiesKHR)(unsafe.Pointer(&out)))
	return &out, Result(res)
}

func (dev *PhysicalDevice) SurfaceFormatsKHR(surface *Surface) ([]SurfaceFormatKHR, Result) {
	var count C.uint32_t
	res := Result(C.domVkGetPhysicalDeviceSurfaceFormatsKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceFormatsKHR], dev.hnd, surface.hnd, &count, nil))
	if res != Success {
		return nil, res
	}
	out := make([]SurfaceFormatKHR, count)
	res = Result(C.domVkGetPhysicalDeviceSurfaceFormatsKHR(dev.instance.fps[vkGetPhysicalDeviceSurfaceFormatsKHR], dev.hnd, surface.hnd, &count, (*C.VkSurfaceFormatKHR)(unsafe.Pointer(&out[0]))))
	return out, res
}

func (dev *PhysicalDevice) SurfacePresentModesKHR(surface *Surface) ([]PresentModeKHR, Result) {
	var count C.uint32_t
	res := Result(C.domVkGetPhysicalDeviceSurfacePresentModesKHR(dev.instance.fps[vkGetPhysicalDeviceSurfacePresentModesKHR], dev.hnd, surface.hnd, &count, nil))
	if res != Success {
		return nil, res
	}
	out := make([]PresentModeKHR, count)
	res = Result(C.domVkGetPhysicalDeviceSurfacePresentModesKHR(dev.instance.fps[vkGetPhysicalDeviceSurfacePresentModesKHR], dev.hnd, surface.hnd, &count, (*C.VkPresentModeKHR)(unsafe.Pointer(&out[0]))))
	return out, res
}
