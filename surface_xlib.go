// +build xlib

package vk

// #cgo pkg-config: vulkan x11
// #cgo CFLAGS: -DVK_NO_PROTOTYPES
// #include <X11/Xlib.h>
// #include <vulkan/vulkan_core.h>
// #include <vulkan/vulkan_xlib.h>
// #include <stdlib.h>
// VkResult domVkCreateXlibSurfaceKHR(PFN_vkCreateXlibSurfaceKHR fp, VkInstance instance, const VkXlibSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
// VkBool32 domVkGetPhysicalDeviceXlibPresentationSupportKHR(PFN_vkGetPhysicalDeviceXlibPresentationSupportKHR fp, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, Display* dpy, VisualID visualID);
import "C"
import "unsafe"

type XlibDisplay = C.Display
type XlibWindow = C.Window
type XlibVisualID = C.VisualID

type XlibSurfaceCreateInfoKHR struct {
	Next   unsafe.Pointer
	Dpy    *XlibDisplay
	Window XlibWindow
}

func (ins *Instance) CreateXlibSurfaceKHR(info *XlibSurfaceCreateInfoKHR) (*Surface, Result) {
	// TODO(dh): support custom allocator
	cInfo := (*C.VkXlibSurfaceCreateInfoKHR)(C.calloc(1, C.size_t(unsafe.Sizeof(C.VkXlibSurfaceCreateInfoKHR{}))))
	defer C.free(unsafe.Pointer(cInfo))
	cInfo.sType = STRUCTURE_TYPE_XLIB_SURFACE_CREATE_INFO_KHR
	cInfo.pNext = info.Next
	cInfo.dpy = info.Dpy
	cInfo.window = info.Window
	var hnd C.VkSurfaceKHR
	res := Result(C.domVkCreateXlibSurfaceKHR(ins.fps[vkCreateXlibSurfaceKHR], ins.hnd, cInfo, nil, &hnd))
	return &Surface{hnd: hnd}, res
}

func (dev *PhysicalDevice) XlibPresentationSupportKHS(queueFamilyIndex uint32, dpy *XlibDisplay, visualID XlibVisualID) bool {
	return C.domVkGetPhysicalDeviceXlibPresentationSupportKHR(dev.instance.fps[vkGetPhysicalDeviceXlibPresentationSupportKHR], dev.hnd, C.uint32_t(queueFamilyIndex), dpy, visualID) == C.VK_TRUE
}
