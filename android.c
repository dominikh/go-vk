// +build android

#include "android.h"

VkResult domVkCreateAndroidSurfaceKHR(PFN_vkCreateAndroidSurfaceKHR fp, VkInstance instance, const VkAndroidSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return (*fp)(instance, pCreateInfo, pAllocator, pSurface);
}
VkResult domVkGetAndroidHardwareBufferPropertiesANDROID(PFN_vkGetAndroidHardwareBufferPropertiesANDROID fp, VkDevice device, const struct AHardwareBuffer* buffer, VkAndroidHardwareBufferPropertiesANDROID* pProperties) {
	return (*fp)(device, buffer, pProperties);
}
VkResult domVkGetMemoryAndroidHardwareBufferANDROID(PFN_vkGetMemoryAndroidHardwareBufferANDROID fp, VkDevice device, const VkMemoryGetAndroidHardwareBufferInfoANDROID* pInfo, struct AHardwareBuffer** pBuffer) {
	return (*fp)(device, pInfo, pBuffer);
}
