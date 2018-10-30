// +build xlib_xrandr

#include "xlib_xrandr.h"

VkResult domVkAcquireXlibDisplayEXT(PFN_vkAcquireXlibDisplayEXT fp, VkPhysicalDevice physicalDevice, Display* dpy, VkDisplayKHR display) {
	return (*fp)(physicalDevice, dpy, display);
}
VkResult domVkGetRandROutputDisplayEXT(PFN_vkGetRandROutputDisplayEXT fp, VkPhysicalDevice physicalDevice, Display* dpy, RROutput rrOutput, VkDisplayKHR* pDisplay) {
	return (*fp)(physicalDevice, dpy, rrOutput, pDisplay);
}
