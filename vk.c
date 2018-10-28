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
void domVkCmdBeginRenderPass(PFN_vkCmdBeginRenderPass fp, VkCommandBuffer commandBuffer, const VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassContents contents) {
	(*fp)(commandBuffer, pRenderPassBegin, contents);
}
void domVkCmdBindPipeline(PFN_vkCmdBindPipeline fp, VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipeline pipeline) {
	(*fp)(commandBuffer, pipelineBindPoint, pipeline);
}
void domVkCmdEndRenderPass(PFN_vkCmdEndRenderPass fp, VkCommandBuffer commandBuffer) {
	(*fp)(commandBuffer);
}
VkResult domVkCreateSemaphore(PFN_vkCreateSemaphore fp, VkDevice device, const VkSemaphoreCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSemaphore* pSemaphore) {
	return (*fp)(device, pCreateInfo, pAllocator, pSemaphore);
}
VkResult domVkQueueSubmit(PFN_vkQueueSubmit fp, VkQueue queue, uint32_t submitCount, const VkSubmitInfo* pSubmits, VkFence fence) {
	return (*fp)(queue, submitCount, pSubmits, fence);
}
VkResult domVkQueuePresentKHR(PFN_vkQueuePresentKHR fp, VkQueue queue, const VkPresentInfoKHR* pPresentInfo) {
	return (*fp)(queue, pPresentInfo);
}
VkResult domVkEnumerateDeviceExtensionProperties(PFN_vkEnumerateDeviceExtensionProperties fp, VkPhysicalDevice physicalDevice, const char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties) {
	return (*fp)(physicalDevice, pLayerName, pPropertyCount, pProperties);
}
VkResult domVkCreateFence(PFN_vkCreateFence fp, VkDevice device, const VkFenceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFence* pFence) {
	return (*fp)(device, pCreateInfo, pAllocator, pFence);
}
VkResult domVkWaitForFences(PFN_vkWaitForFences fp, VkDevice device, uint32_t fenceCount, const VkFence* pFences, VkBool32 waitAll, uint64_t timeout) {
	return (*fp)(device, fenceCount, pFences, waitAll, timeout);
}
VkResult domVkResetFences(PFN_vkResetFences fp, VkDevice device, uint32_t fenceCount, const VkFence* pFences) {
	return (*fp)(device, fenceCount, pFences);
}
VkResult domVkCreateBuffer(PFN_vkCreateBuffer fp, VkDevice device, const VkBufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkBuffer* pBuffer) {
	return (*fp)(device, pCreateInfo, pAllocator, pBuffer);
}
void domVkGetBufferMemoryRequirements(PFN_vkGetBufferMemoryRequirements fp, VkDevice device, VkBuffer buffer, VkMemoryRequirements* pMemoryRequirements) {
	(*fp)(device, buffer, pMemoryRequirements);
}
void domVkGetPhysicalDeviceMemoryProperties(PFN_vkGetPhysicalDeviceMemoryProperties fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties* pMemoryProperties) {
	(*fp)(physicalDevice, pMemoryProperties);
}
void domVkGetPhysicalDeviceProperties2(PFN_vkGetPhysicalDeviceProperties2 fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties2* pProperties) {
	(*fp)(physicalDevice, pProperties);
}
VkResult domVkAllocateMemory(PFN_vkAllocateMemory fp, VkDevice device, const VkMemoryAllocateInfo* pAllocateInfo, const VkAllocationCallbacks* pAllocator, VkDeviceMemory* pMemory) {
	return (*fp)(device, pAllocateInfo, pAllocator, pMemory);
}
VkResult domVkBindBufferMemory(PFN_vkBindBufferMemory fp, VkDevice device, VkBuffer buffer, VkDeviceMemory memory, VkDeviceSize memoryOffset) {
	return (*fp)(device, buffer, memory, memoryOffset);
}
VkResult domVkBindBufferMemory2(PFN_vkBindBufferMemory2 fp, VkDevice device, uint32_t bindInfoCount, const VkBindBufferMemoryInfo* pBindInfos) {
	return (*fp)(device, bindInfoCount, pBindInfos);
}
VkResult domVkMapMemory(PFN_vkMapMemory fp, VkDevice device, VkDeviceMemory memory, VkDeviceSize offset, VkDeviceSize size, VkMemoryMapFlags flags, void** ppData) {
	return (*fp)(device, memory, offset, size, flags, ppData);
}
void domVkUnmapMemory(PFN_vkUnmapMemory fp, VkDevice device, VkDeviceMemory memory) {
	(*fp)(device, memory);
}
void domVkFreeMemory(PFN_vkFreeMemory fp, VkDevice device, VkDeviceMemory memory, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, memory, pAllocator);
}
