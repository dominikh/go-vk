// Copyright (c) 2018 Dominik Honnef

#include <vulkan/vulkan_core.h>

VkResult domVkCreateInstance(PFN_vkCreateInstance fp, const VkInstanceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkInstance* pInstance) {
	return (*fp)(pCreateInfo, pAllocator, pInstance);
}
VkResult domVkEnumeratePhysicalDevices(PFN_vkEnumeratePhysicalDevices fp, VkInstance instance, uint32_t* pPhysicalDeviceCount, VkPhysicalDevice* pPhysicalDevices) {
	return (*fp)(instance, pPhysicalDeviceCount, pPhysicalDevices);
}
void domVkGetPhysicalDeviceProperties(PFN_vkGetPhysicalDeviceProperties fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties* pProperties) {
	(*fp)(physicalDevice,pProperties);
}
void domVkGetPhysicalDeviceFeatures(PFN_vkGetPhysicalDeviceFeatures fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures* pFeatures) {
	(*fp)(physicalDevice, pFeatures);
}
void domVkGetPhysicalDeviceQueueFamilyProperties(PFN_vkGetPhysicalDeviceQueueFamilyProperties fp, VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties* pQueueFamilyProperties) {
	(*fp)(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}
VkResult domVkCreateDevice(PFN_vkCreateDevice fp, VkPhysicalDevice physicalDevice, const VkDeviceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDevice* pDevice) {
	return (*fp)(physicalDevice, pCreateInfo, pAllocator, pDevice);
}
void domVkGetDeviceQueue(PFN_vkGetDeviceQueue fp, VkDevice device, uint32_t queueFamilyIndex, uint32_t queueIndex, VkQueue* pQueue) {
	(*fp)(device, queueFamilyIndex, queueIndex, pQueue);
}
PFN_vkVoidFunction domVkGetDeviceProcAddr(PFN_vkGetDeviceProcAddr fp, VkDevice device, const char* pName) {
	return (*fp)(device, pName);
}
VkResult domVkGetPhysicalDeviceSurfaceSupportKHR(PFN_vkGetPhysicalDeviceSurfaceSupportKHR fp, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, VkSurfaceKHR surface, VkBool32* pSupported) {
	return (*fp)(physicalDevice, queueFamilyIndex, surface, pSupported);
}
VkResult domVkGetPhysicalDeviceSurfaceCapabilitiesKHR(PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR fp, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, VkSurfaceCapabilitiesKHR* pSurfaceCapabilities) {
	return (*fp)(physicalDevice, surface, pSurfaceCapabilities);
}
VkResult domVkGetPhysicalDeviceSurfaceFormatsKHR(PFN_vkGetPhysicalDeviceSurfaceFormatsKHR fp, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pSurfaceFormatCount, VkSurfaceFormatKHR* pSurfaceFormats) {
	return (*fp)(physicalDevice, surface, pSurfaceFormatCount, pSurfaceFormats);
}
VkResult domVkGetPhysicalDeviceSurfacePresentModesKHR(PFN_vkGetPhysicalDeviceSurfacePresentModesKHR fp, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pPresentModeCount, VkPresentModeKHR* pPresentModes) {
	return (*fp)(physicalDevice, surface, pPresentModeCount, pPresentModes);
}
VkResult domVkCreateCommandPool(PFN_vkCreateCommandPool fp, VkDevice device, const VkCommandPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkCommandPool* pCommandPool) {
	return (*fp)(device, pCreateInfo, pAllocator, pCommandPool);
}
void domVkTrimCommandPool(PFN_vkTrimCommandPool fp, VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlags flags) {
	(*fp)(device, commandPool, flags);
}
VkResult domVkResetCommandPool(PFN_vkResetCommandPool fp, VkDevice device, VkCommandPool commandPool, VkCommandPoolResetFlags flags) {
	return (*fp)(device, commandPool, flags);
}
VkResult domVkAllocateCommandBuffers(PFN_vkAllocateCommandBuffers fp, VkDevice device, const VkCommandBufferAllocateInfo* pAllocateInfo, VkCommandBuffer* pCommandBuffers) {
	return (*fp)(device, pAllocateInfo, pCommandBuffers);
}
VkResult domVkResetCommandBuffer(PFN_vkResetCommandBuffer fp, VkCommandBuffer commandBuffer, VkCommandBufferResetFlags flags) {
	return (*fp)(commandBuffer, flags);
}
void domVkFreeCommandBuffers(PFN_vkFreeCommandBuffers fp, VkDevice device, VkCommandPool commandPool, uint32_t commandBufferCount, const VkCommandBuffer* pCommandBuffers) {
	(*fp)(device, commandPool, commandBufferCount, pCommandBuffers);
}
VkResult domVkEndCommandBuffer(PFN_vkEndCommandBuffer fp, VkCommandBuffer commandBuffer) {
	return (*fp)(commandBuffer);
}
VkResult domVkBeginCommandBuffer(PFN_vkBeginCommandBuffer fp, VkCommandBuffer commandBuffer, const VkCommandBufferBeginInfo* pBeginInfo) {
	return (*fp)(commandBuffer, pBeginInfo);
}
void domVkCmdSetLineWidth(PFN_vkCmdSetLineWidth fp, VkCommandBuffer commandBuffer, float lineWidth) {
	(*fp)(commandBuffer, lineWidth);
}
void domVkCmdSetDepthBias(PFN_vkCmdSetDepthBias fp, VkCommandBuffer commandBuffer, float depthBiasConstantFactor, float depthBiasClamp, float depthBiasSlopeFactor) {
	(*fp)(commandBuffer, depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor);
}
void domVkCmdSetBlendConstants(PFN_vkCmdSetBlendConstants fp, VkCommandBuffer commandBuffer, const float blendConstants[4]) {
	(*fp)(commandBuffer, blendConstants);
}
void domVkCmdDraw(PFN_vkCmdDraw fp, VkCommandBuffer commandBuffer, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance) {
	(*fp)(commandBuffer, vertexCount, instanceCount, firstVertex, firstInstance);
}
VkResult domVkQueueWaitIdle(PFN_vkQueueWaitIdle fp, VkQueue queue) {
	return (*fp)(queue);
}
VkResult domVkDeviceWaitIdle(PFN_vkDeviceWaitIdle fp, VkDevice device) {
	return (*fp)(device);
}
VkResult domVkCreateImageView(PFN_vkCreateImageView fp, VkDevice device, const VkImageViewCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkImageView* pView) {
	return (*fp)(device, pCreateInfo, pAllocator, pView);
}
VkResult domVkCreateShaderModule(PFN_vkCreateShaderModule fp, VkDevice device, const VkShaderModuleCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkShaderModule* pShaderModule) {
	return (*fp)(device, pCreateInfo, pAllocator, pShaderModule);
}
VkResult domVkCreateGraphicsPipelines(PFN_vkCreateGraphicsPipelines fp, VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkGraphicsPipelineCreateInfo* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return (*fp)(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}
VkResult domVkCreatePipelineLayout(PFN_vkCreatePipelineLayout fp, VkDevice device, const VkPipelineLayoutCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkPipelineLayout* pPipelineLayout) {
	return (*fp)(device, pCreateInfo, pAllocator, pPipelineLayout);
}
VkResult domVkCreateRenderPass(PFN_vkCreateRenderPass fp, VkDevice device, const VkRenderPassCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass) {
	return (*fp)(device, pCreateInfo, pAllocator, pRenderPass);
}
VkResult domVkCreateFramebuffer(PFN_vkCreateFramebuffer fp, VkDevice device, const VkFramebufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFramebuffer* pFramebuffer) {
	return (*fp)(device, pCreateInfo, pAllocator, pFramebuffer);
}
