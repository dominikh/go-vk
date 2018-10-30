#pragma once

#include "vk.h"
#include <vulkan/vulkan_win32.h>

VkResult domVkCreateWin32SurfaceKHR(PFN_vkCreateWin32SurfaceKHR fp, VkInstance instance, const VkWin32SurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
VkBool32 domVkGetPhysicalDeviceWin32PresentationSupportKHR(PFN_vkGetPhysicalDeviceWin32PresentationSupportKHR fp, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex);
VkResult domVkGetMemoryWin32HandleKHR(PFN_vkGetMemoryWin32HandleKHR fp, VkDevice device, const VkMemoryGetWin32HandleInfoKHR* pGetWin32HandleInfo, HANDLE* pHandle);
VkResult domVkGetMemoryWin32HandlePropertiesKHR(PFN_vkGetMemoryWin32HandlePropertiesKHR fp, VkDevice device, VkExternalMemoryHandleTypeFlagBits handleType, HANDLE handle, VkMemoryWin32HandlePropertiesKHR* pMemoryWin32HandleProperties);
VkResult domVkImportSemaphoreWin32HandleKHR(PFN_vkImportSemaphoreWin32HandleKHR fp, VkDevice device, const VkImportSemaphoreWin32HandleInfoKHR* pImportSemaphoreWin32HandleInfo);
VkResult domVkGetSemaphoreWin32HandleKHR(PFN_vkGetSemaphoreWin32HandleKHR fp, VkDevice device, const VkSemaphoreGetWin32HandleInfoKHR* pGetWin32HandleInfo, HANDLE* pHandle);
VkResult domVkImportFenceWin32HandleKHR(PFN_vkImportFenceWin32HandleKHR fp, VkDevice device, const VkImportFenceWin32HandleInfoKHR* pImportFenceWin32HandleInfo);
VkResult domVkGetFenceWin32HandleKHR(PFN_vkGetFenceWin32HandleKHR fp, VkDevice device, const VkFenceGetWin32HandleInfoKHR* pGetWin32HandleInfo, HANDLE* pHandle);
VkResult domVkGetMemoryWin32HandleNV(PFN_vkGetMemoryWin32HandleNV fp, VkDevice device, VkDeviceMemory memory, VkExternalMemoryHandleTypeFlagsNV handleType, HANDLE* pHandle);
