#pragma once

#include "vk.h"
#include <vulkan/vulkan_android.h>

VkResult domVkCreateAndroidSurfaceKHR(PFN_vkCreateAndroidSurfaceKHR fp, VkInstance instance, const VkAndroidSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
VkResult domVkGetAndroidHardwareBufferPropertiesANDROID(PFN_vkGetAndroidHardwareBufferPropertiesANDROID fp, VkDevice device, const struct AHardwareBuffer* buffer, VkAndroidHardwareBufferPropertiesANDROID* pProperties);
VkResult domVkGetMemoryAndroidHardwareBufferANDROID(PFN_vkGetMemoryAndroidHardwareBufferANDROID fp, VkDevice device, const VkMemoryGetAndroidHardwareBufferInfoANDROID* pInfo, struct AHardwareBuffer** pBuffer);
