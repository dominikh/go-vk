#pragma once

#include <xcb/xcb.h>
#include "vk.h"
#include <vulkan/vulkan_xcb.h>

VkResult domVkCreateXcbSurfaceKHR(PFN_vkCreateXcbSurfaceKHR fp, VkInstance instance, const VkXcbSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
VkBool32 domVkGetPhysicalDeviceXcbPresentationSupportKHR(PFN_vkGetPhysicalDeviceXcbPresentationSupportKHR fp, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, xcb_connection_t* connection, xcb_visualid_t visual_id);
