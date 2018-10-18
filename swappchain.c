// Copyright (c) 2018 Dominik Honnef

#include <vulkan/vulkan_core.h>

VkResult domVkCreateSwapchainKHR(PFN_vkCreateSwapchainKHR fp, VkDevice device, const VkSwapchainCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSwapchainKHR* pSwapchain) {
	return (*fp)(device, pCreateInfo, pAllocator, pSwapchain);
}
VkResult domVkGetSwapchainImagesKHR(PFN_vkGetSwapchainImagesKHR fp, VkDevice device, VkSwapchainKHR swapchain, uint32_t* pSwapchainImageCount, VkImage* pSwapchainImages) {
	return (*fp)(device, swapchain, pSwapchainImageCount, pSwapchainImages);
}
