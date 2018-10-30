#pragma once

#define VK_NO_PROTOTYPES 1;
#define VK_DEFINE_NON_DISPATCHABLE_HANDLE(object) typedef uintptr_t object;

#include <vulkan/vulkan_core.h>

VKAPI_ATTR PFN_vkVoidFunction VKAPI_CALL vkGetInstanceProcAddr(VkInstance instance, const char *pName);

VkResult domVkCreateInstance(PFN_vkCreateInstance fp, const VkInstanceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkInstance* pInstance);
VkResult domVkEnumeratePhysicalDevices(PFN_vkEnumeratePhysicalDevices fp, VkInstance instance, uint32_t* pPhysicalDeviceCount, VkPhysicalDevice* pPhysicalDevices);
void     domVkGetPhysicalDeviceProperties(PFN_vkGetPhysicalDeviceProperties fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties* pProperties);
void     domVkGetPhysicalDeviceFeatures(PFN_vkGetPhysicalDeviceFeatures fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures* pFeatures);
void     domVkGetPhysicalDeviceQueueFamilyProperties(PFN_vkGetPhysicalDeviceQueueFamilyProperties fp, VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties* pQueueFamilyProperties);
VkResult domVkCreateDevice(PFN_vkCreateDevice fp, VkPhysicalDevice physicalDevice, const VkDeviceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDevice* pDevice);
void     domVkGetDeviceQueue(PFN_vkGetDeviceQueue fp, VkDevice device, uint32_t queueFamilyIndex, uint32_t queueIndex, VkQueue* pQueue);
PFN_vkVoidFunction domVkGetDeviceProcAddr(PFN_vkGetDeviceProcAddr fp, VkDevice device, const char* pName);
VkResult domVkCreateCommandPool(PFN_vkCreateCommandPool fp, VkDevice device, const VkCommandPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkCommandPool* pCommandPool);
void     domVkTrimCommandPool(PFN_vkTrimCommandPool fp, VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlags flags);
VkResult domVkResetCommandPool(PFN_vkResetCommandPool fp, VkDevice device, VkCommandPool commandPool, VkCommandPoolResetFlags flags);
VkResult domVkAllocateCommandBuffers(PFN_vkAllocateCommandBuffers fp, VkDevice device, const VkCommandBufferAllocateInfo* pAllocateInfo, VkCommandBuffer* pCommandBuffers);
VkResult domVkResetCommandBuffer(PFN_vkResetCommandBuffer fp, VkCommandBuffer commandBuffer, VkCommandBufferResetFlags flags);
void     domVkFreeCommandBuffers(PFN_vkFreeCommandBuffers fp, VkDevice device, VkCommandPool commandPool, uint32_t commandBufferCount, const VkCommandBuffer* pCommandBuffers);
VkResult domVkEndCommandBuffer(PFN_vkEndCommandBuffer fp, VkCommandBuffer commandBuffer);
VkResult domVkBeginCommandBuffer(PFN_vkBeginCommandBuffer fp, VkCommandBuffer commandBuffer, const VkCommandBufferBeginInfo* pBeginInfo);
void     domVkCmdSetLineWidth(PFN_vkCmdSetLineWidth fp, VkCommandBuffer commandBuffer, float lineWidth);
void     domVkCmdSetDepthBias(PFN_vkCmdSetDepthBias fp, VkCommandBuffer commandBuffer, float depthBiasConstantFactor, float depthBiasClamp, float depthBiasSlopeFactor);
void     domVkCmdSetBlendConstants(PFN_vkCmdSetBlendConstants fp, VkCommandBuffer commandBuffer, const float blendConstants[4]);
void     domVkCmdDraw(PFN_vkCmdDraw fp, VkCommandBuffer commandBuffer, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance);
VkResult domVkQueueWaitIdle(PFN_vkQueueWaitIdle fp, VkQueue queue);
VkResult domVkDeviceWaitIdle(PFN_vkDeviceWaitIdle fp, VkDevice device);
VkResult domVkCreateImageView(PFN_vkCreateImageView fp, VkDevice device, const VkImageViewCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkImageView* pView);
VkResult domVkCreateShaderModule(PFN_vkCreateShaderModule fp, VkDevice device, const VkShaderModuleCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkShaderModule* pShaderModule);
VkResult domVkCreateGraphicsPipelines(PFN_vkCreateGraphicsPipelines fp, VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkGraphicsPipelineCreateInfo* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines);
VkResult domVkCreatePipelineLayout(PFN_vkCreatePipelineLayout fp, VkDevice device, const VkPipelineLayoutCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkPipelineLayout* pPipelineLayout);
VkResult domVkCreateRenderPass(PFN_vkCreateRenderPass fp, VkDevice device, const VkRenderPassCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass);
VkResult domVkCreateFramebuffer(PFN_vkCreateFramebuffer fp, VkDevice device, const VkFramebufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFramebuffer* pFramebuffer);
void     domVkCmdBeginRenderPass(PFN_vkCmdBeginRenderPass fp, VkCommandBuffer commandBuffer, const VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassContents contents);
void     domVkCmdBindPipeline(PFN_vkCmdBindPipeline fp, VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipeline pipeline);
void     domVkCmdEndRenderPass(PFN_vkCmdEndRenderPass fp, VkCommandBuffer commandBuffer);
VkResult domVkCreateSemaphore(PFN_vkCreateSemaphore fp, VkDevice device, const VkSemaphoreCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSemaphore* pSemaphore);
VkResult domVkQueueSubmit(PFN_vkQueueSubmit fp, VkQueue queue, uint32_t submitCount, const VkSubmitInfo* pSubmits, VkFence fence);
VkResult domVkEnumerateDeviceExtensionProperties(PFN_vkEnumerateDeviceExtensionProperties fp, VkPhysicalDevice physicalDevice, const char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties);
VkResult domVkCreateFence(PFN_vkCreateFence fp, VkDevice device, const VkFenceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFence* pFence);
VkResult domVkWaitForFences(PFN_vkWaitForFences fp, VkDevice device, uint32_t fenceCount, const VkFence* pFences, VkBool32 waitAll, uint64_t timeout);
VkResult domVkResetFences(PFN_vkResetFences fp, VkDevice device, uint32_t fenceCount, const VkFence* pFences);
VkResult domVkCreateBuffer(PFN_vkCreateBuffer fp, VkDevice device, const VkBufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkBuffer* pBuffer);
void     domVkGetBufferMemoryRequirements(PFN_vkGetBufferMemoryRequirements fp, VkDevice device, VkBuffer buffer, VkMemoryRequirements* pMemoryRequirements);
void     domVkGetPhysicalDeviceMemoryProperties(PFN_vkGetPhysicalDeviceMemoryProperties fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties* pMemoryProperties);
void     domVkGetPhysicalDeviceProperties2(PFN_vkGetPhysicalDeviceProperties2 fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties2* pProperties);
VkResult domVkAllocateMemory(PFN_vkAllocateMemory fp, VkDevice device, const VkMemoryAllocateInfo* pAllocateInfo, const VkAllocationCallbacks* pAllocator, VkDeviceMemory* pMemory);
VkResult domVkBindBufferMemory(PFN_vkBindBufferMemory fp, VkDevice device, VkBuffer buffer, VkDeviceMemory memory, VkDeviceSize memoryOffset);
VkResult domVkBindBufferMemory2(PFN_vkBindBufferMemory2 fp, VkDevice device, uint32_t bindInfoCount, const VkBindBufferMemoryInfo* pBindInfos);
VkResult domVkMapMemory(PFN_vkMapMemory fp, VkDevice device, VkDeviceMemory memory, VkDeviceSize offset, VkDeviceSize size, VkMemoryMapFlags flags, void** ppData);
void     domVkUnmapMemory(PFN_vkUnmapMemory fp, VkDevice device, VkDeviceMemory memory);
void     domVkFreeMemory(PFN_vkFreeMemory fp, VkDevice device, VkDeviceMemory memory, const VkAllocationCallbacks* pAllocator);
VkResult domVkCreateImage(PFN_vkCreateImage fp, VkDevice device, const VkImageCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkImage* pImage);
void     domVkCmdSetViewport(PFN_vkCmdSetViewport fp, VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, const VkViewport* pViewports);
void     domVkCmdSetScissor(PFN_vkCmdSetScissor fp, VkCommandBuffer commandBuffer, uint32_t firstScissor, uint32_t scissorCount, const VkRect2D* pScissors);
void     domVkCmdSetDeviceMask(PFN_vkCmdSetDeviceMask fp, VkCommandBuffer commandBuffer, uint32_t deviceMask);
void     domVkCmdSetDepthBounds(PFN_vkCmdSetDepthBounds fp, VkCommandBuffer commandBuffer, float minDepthBounds, float maxDepthBounds);
void     domVkCmdPushConstants(PFN_vkCmdPushConstants fp, VkCommandBuffer commandBuffer, VkPipelineLayout layout, VkShaderStageFlags stageFlags, uint32_t offset, uint32_t size, const void* pValues);
void     domVkCmdFillBuffer(PFN_vkCmdFillBuffer fp, VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize size, uint32_t data);
void     domVkCmdDispatch(PFN_vkCmdDispatch fp, VkCommandBuffer commandBuffer, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ);
