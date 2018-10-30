#pragma once

#include "vk.h"
#include <X11/extensions/Xrandr.h>
#include <vulkan/vulkan_xlib_xrandr.h>

VkResult domVkAcquireXlibDisplayEXT(PFN_vkAcquireXlibDisplayEXT fp, VkPhysicalDevice physicalDevice, Display* dpy, VkDisplayKHR display);
VkResult domVkGetRandROutputDisplayEXT(PFN_vkGetRandROutputDisplayEXT fp, VkPhysicalDevice physicalDevice, Display* dpy, RROutput rrOutput, VkDisplayKHR* pDisplay);
