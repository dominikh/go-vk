#pragma once

#include <X11/Xlib.h>
#include "vk.h"
#include <vulkan/vulkan_xlib.h>
#include <stdlib.h>
VkResult domVkCreateXlibSurfaceKHR(PFN_vkCreateXlibSurfaceKHR fp, VkInstance instance, const VkXlibSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
VkBool32 domVkGetPhysicalDeviceXlibPresentationSupportKHR(PFN_vkGetPhysicalDeviceXlibPresentationSupportKHR fp, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, Display* dpy, VisualID visualID);
