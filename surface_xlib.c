// +build xlib

#include <X11/Xlib.h>
#include <vulkan/vulkan_core.h>
#include <vulkan/vulkan_xlib.h>

VkResult domVkCreateXlibSurfaceKHR(PFN_vkCreateXlibSurfaceKHR fp, VkInstance instance, const VkXlibSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return (*fp)(instance, pCreateInfo, pAllocator, pSurface);
}
VkBool32 domVkGetPhysicalDeviceXlibPresentationSupportKHR(PFN_vkGetPhysicalDeviceXlibPresentationSupportKHR fp, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, Display* dpy, VisualID visualID) {
	return (*fp)(physicalDevice, queueFamilyIndex, dpy, visualID);
}
