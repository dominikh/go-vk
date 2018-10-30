// +build wayland

#include "vk.h"
#include <vulkan/vulkan_wayland.h>

VkResult domVkCreateWaylandSurfaceKHR(PFN_vkCreateWaylandSurfaceKHR fp, VkInstance instance, const VkWaylandSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return (*fp)(instance, pCreateInfo, pAllocator, pSurface);
}
VkBool32 domVkGetPhysicalDeviceWaylandPresentationSupportKHR(PFN_vkGetPhysicalDeviceWaylandPresentationSupportKHR fp, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, struct wl_display* display) {
	return (*fp)(physicalDevice, queueFamilyIndex, display);
}
