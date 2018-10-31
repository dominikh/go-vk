#include "vk.h"

PFN_vkVoidFunction domVkGetDeviceProcAddr(PFN_vkGetDeviceProcAddr fp, VkDevice device, const char* pName) {
	return (*fp)(device, pName);
}

VkResult domVkAllocateCommandBuffers(PFN_vkAllocateCommandBuffers fp, VkDevice device, const VkCommandBufferAllocateInfo* pAllocateInfo, VkCommandBuffer* pCommandBuffers) {
	return (*fp)(device, pAllocateInfo, pCommandBuffers);
}
VkResult domVkAllocateDescriptorSets(PFN_vkAllocateDescriptorSets fp, VkDevice device, const VkDescriptorSetAllocateInfo* pAllocateInfo, VkDescriptorSet* pDescriptorSets) {
	return (*fp)(device, pAllocateInfo, pDescriptorSets);
}
VkResult domVkAllocateMemory(PFN_vkAllocateMemory fp, VkDevice device, const VkMemoryAllocateInfo* pAllocateInfo, const VkAllocationCallbacks* pAllocator, VkDeviceMemory* pMemory) {
	return (*fp)(device, pAllocateInfo, pAllocator, pMemory);
}
VkResult domVkBeginCommandBuffer(PFN_vkBeginCommandBuffer fp, VkCommandBuffer commandBuffer, const VkCommandBufferBeginInfo* pBeginInfo) {
	return (*fp)(commandBuffer, pBeginInfo);
}
VkResult domVkBindBufferMemory2(PFN_vkBindBufferMemory2 fp, VkDevice device, uint32_t bindInfoCount, const VkBindBufferMemoryInfo* pBindInfos) {
	return (*fp)(device, bindInfoCount, pBindInfos);
}
VkResult domVkBindBufferMemory(PFN_vkBindBufferMemory fp, VkDevice device, VkBuffer buffer, VkDeviceMemory memory, VkDeviceSize memoryOffset) {
	return (*fp)(device, buffer, memory, memoryOffset);
}
VkResult domVkBindImageMemory(PFN_vkBindImageMemory fp, VkDevice device, VkImage image, VkDeviceMemory memory, VkDeviceSize memoryOffset) {
	return (*fp)(device, image, memory, memoryOffset);
}
VkResult domVkCreateBuffer(PFN_vkCreateBuffer fp, VkDevice device, const VkBufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkBuffer* pBuffer) {
	return (*fp)(device, pCreateInfo, pAllocator, pBuffer);
}
VkResult domVkCreateBufferView(PFN_vkCreateBufferView fp, VkDevice device, const VkBufferViewCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkBufferView* pView) {
	return (*fp)(device, pCreateInfo, pAllocator, pView);
}
VkResult domVkCreateCommandPool(PFN_vkCreateCommandPool fp, VkDevice device, const VkCommandPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkCommandPool* pCommandPool) {
	return (*fp)(device, pCreateInfo, pAllocator, pCommandPool);
}
VkResult domVkCreateComputePipelines(PFN_vkCreateComputePipelines fp, VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkComputePipelineCreateInfo* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return (*fp)(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}
VkResult domVkCreateDescriptorPool(PFN_vkCreateDescriptorPool fp, VkDevice device, const VkDescriptorPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDescriptorPool* pDescriptorPool) {
	return (*fp)(device, pCreateInfo, pAllocator, pDescriptorPool);
}
VkResult domVkCreateDescriptorSetLayout(PFN_vkCreateDescriptorSetLayout fp, VkDevice device, const VkDescriptorSetLayoutCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDescriptorSetLayout* pSetLayout) {
	return (*fp)(device, pCreateInfo, pAllocator, pSetLayout);
}
VkResult domVkCreateDevice(PFN_vkCreateDevice fp, VkPhysicalDevice physicalDevice, const VkDeviceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDevice* pDevice) {
	return (*fp)(physicalDevice, pCreateInfo, pAllocator, pDevice);
}
VkResult domVkCreateEvent(PFN_vkCreateEvent fp, VkDevice device, const VkEventCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkEvent* pEvent) {
	return (*fp)(device, pCreateInfo, pAllocator, pEvent);
}
VkResult domVkCreateFence(PFN_vkCreateFence fp, VkDevice device, const VkFenceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFence* pFence) {
	return (*fp)(device, pCreateInfo, pAllocator, pFence);
}
VkResult domVkCreateFramebuffer(PFN_vkCreateFramebuffer fp, VkDevice device, const VkFramebufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFramebuffer* pFramebuffer) {
	return (*fp)(device, pCreateInfo, pAllocator, pFramebuffer);
}
VkResult domVkCreateGraphicsPipelines(PFN_vkCreateGraphicsPipelines fp, VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkGraphicsPipelineCreateInfo* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return (*fp)(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}
VkResult domVkCreateImage(PFN_vkCreateImage fp, VkDevice device, const VkImageCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkImage* pImage) {
	return (*fp)(device, pCreateInfo, pAllocator, pImage);
}
VkResult domVkCreateImageView(PFN_vkCreateImageView fp, VkDevice device, const VkImageViewCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkImageView* pView) {
	return (*fp)(device, pCreateInfo, pAllocator, pView);
}
VkResult domVkCreateInstance(PFN_vkCreateInstance fp, const VkInstanceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkInstance* pInstance) {
	return (*fp)(pCreateInfo, pAllocator, pInstance);
}
VkResult domVkCreatePipelineCache(PFN_vkCreatePipelineCache fp, VkDevice device, const VkPipelineCacheCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkPipelineCache* pPipelineCache) {
	return (*fp)(device, pCreateInfo, pAllocator, pPipelineCache);
}
VkResult domVkCreatePipelineLayout(PFN_vkCreatePipelineLayout fp, VkDevice device, const VkPipelineLayoutCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkPipelineLayout* pPipelineLayout) {
	return (*fp)(device, pCreateInfo, pAllocator, pPipelineLayout);
}
VkResult domVkCreateQueryPool(PFN_vkCreateQueryPool fp, VkDevice device, const VkQueryPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkQueryPool* pQueryPool) {
	return (*fp)(device, pCreateInfo, pAllocator, pQueryPool);
}
VkResult domVkCreateRenderPass(PFN_vkCreateRenderPass fp, VkDevice device, const VkRenderPassCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass) {
	return (*fp)(device, pCreateInfo, pAllocator, pRenderPass);
}
VkResult domVkCreateSampler(PFN_vkCreateSampler fp, VkDevice device, const VkSamplerCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSampler* pSampler) {
	return (*fp)(device, pCreateInfo, pAllocator, pSampler);
}
VkResult domVkCreateSemaphore(PFN_vkCreateSemaphore fp, VkDevice device, const VkSemaphoreCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSemaphore* pSemaphore) {
	return (*fp)(device, pCreateInfo, pAllocator, pSemaphore);
}
VkResult domVkCreateShaderModule(PFN_vkCreateShaderModule fp, VkDevice device, const VkShaderModuleCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkShaderModule* pShaderModule) {
	return (*fp)(device, pCreateInfo, pAllocator, pShaderModule);
}
VkResult domVkDeviceWaitIdle(PFN_vkDeviceWaitIdle fp, VkDevice device) {
	return (*fp)(device);
}
VkResult domVkEndCommandBuffer(PFN_vkEndCommandBuffer fp, VkCommandBuffer commandBuffer) {
	return (*fp)(commandBuffer);
}
VkResult domVkEnumerateDeviceExtensionProperties(PFN_vkEnumerateDeviceExtensionProperties fp, VkPhysicalDevice physicalDevice, const char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties) {
	return (*fp)(physicalDevice, pLayerName, pPropertyCount, pProperties);
}
VkResult domVkEnumerateDeviceLayerProperties(PFN_vkEnumerateDeviceLayerProperties fp, VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkLayerProperties* pProperties) {
	return (*fp)(physicalDevice, pPropertyCount, pProperties);
}
VkResult domVkEnumerateInstanceExtensionProperties(PFN_vkEnumerateInstanceExtensionProperties fp, const char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties) {
	return (*fp)(pLayerName, pPropertyCount, pProperties);
}
VkResult domVkEnumerateInstanceLayerProperties(PFN_vkEnumerateInstanceLayerProperties fp, uint32_t* pPropertyCount, VkLayerProperties* pProperties) {
	return (*fp)(pPropertyCount, pProperties);
}
VkResult domVkEnumeratePhysicalDevices(PFN_vkEnumeratePhysicalDevices fp, VkInstance instance, uint32_t* pPhysicalDeviceCount, VkPhysicalDevice* pPhysicalDevices) {
	return (*fp)(instance, pPhysicalDeviceCount, pPhysicalDevices);
}
VkResult domVkFlushMappedMemoryRanges(PFN_vkFlushMappedMemoryRanges fp, VkDevice device, uint32_t memoryRangeCount, const VkMappedMemoryRange* pMemoryRanges) {
	return (*fp)(device, memoryRangeCount, pMemoryRanges);
}
VkResult domVkFreeDescriptorSets(PFN_vkFreeDescriptorSets fp, VkDevice device, VkDescriptorPool descriptorPool, uint32_t descriptorSetCount, const VkDescriptorSet* pDescriptorSets) {
	return (*fp)(device, descriptorPool, descriptorSetCount, pDescriptorSets);
}
VkResult domVkGetEventStatus(PFN_vkGetEventStatus fp, VkDevice device, VkEvent event) {
	return (*fp)(device, event);
}
VkResult domVkGetFenceStatus(PFN_vkGetFenceStatus fp, VkDevice device, VkFence fence) {
	return (*fp)(device, fence);
}
VkResult domVkGetPhysicalDeviceImageFormatProperties(PFN_vkGetPhysicalDeviceImageFormatProperties fp, VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkImageTiling tiling, VkImageUsageFlags usage, VkImageCreateFlags flags, VkImageFormatProperties* pImageFormatProperties) {
	return (*fp)(physicalDevice, format, type, tiling, usage, flags, pImageFormatProperties);
}
VkResult domVkGetPipelineCacheData(PFN_vkGetPipelineCacheData fp, VkDevice device, VkPipelineCache pipelineCache, size_t* pDataSize, void* pData) {
	return (*fp)(device, pipelineCache,pDataSize, pData);
}
VkResult domVkGetQueryPoolResults(PFN_vkGetQueryPoolResults fp, VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, size_t dataSize, void* pData, VkDeviceSize stride, VkQueryResultFlags flags) {
	return (*fp)(device, queryPool, firstQuery, queryCount, dataSize, pData, stride, flags);
}
VkResult domVkInvalidateMappedMemoryRanges(PFN_vkInvalidateMappedMemoryRanges fp, VkDevice device, uint32_t memoryRangeCount, const VkMappedMemoryRange* pMemoryRanges) {
	return (*fp)(device, memoryRangeCount, pMemoryRanges);
}
VkResult domVkMapMemory(PFN_vkMapMemory fp, VkDevice device, VkDeviceMemory memory, VkDeviceSize offset, VkDeviceSize size, VkMemoryMapFlags flags, void** ppData) {
	return (*fp)(device, memory, offset, size, flags, ppData);
}
VkResult domVkMergePipelineCaches(PFN_vkMergePipelineCaches fp, VkDevice device, VkPipelineCache dstCache, uint32_t srcCacheCount, const VkPipelineCache* pSrcCaches) {
	return (*fp)(device, dstCache, srcCacheCount, pSrcCaches);
}
VkResult domVkQueueBindSparse(PFN_vkQueueBindSparse fp, VkQueue queue, uint32_t bindInfoCount, const VkBindSparseInfo* pBindInfo, VkFence fence) {
	return (*fp)(queue, bindInfoCount, pBindInfo, fence);
}
VkResult domVkQueueSubmit(PFN_vkQueueSubmit fp, VkQueue queue, uint32_t submitCount, const VkSubmitInfo* pSubmits, VkFence fence) {
	return (*fp)(queue, submitCount, pSubmits, fence);
}
VkResult domVkQueueWaitIdle(PFN_vkQueueWaitIdle fp, VkQueue queue) {
	return (*fp)(queue);
}
VkResult domVkResetCommandBuffer(PFN_vkResetCommandBuffer fp, VkCommandBuffer commandBuffer, VkCommandBufferResetFlags flags) {
	return (*fp)(commandBuffer, flags);
}
VkResult domVkResetCommandPool(PFN_vkResetCommandPool fp, VkDevice device, VkCommandPool commandPool, VkCommandPoolResetFlags flags) {
	return (*fp)(device, commandPool, flags);
}
VkResult domVkResetDescriptorPool(PFN_vkResetDescriptorPool fp, VkDevice device, VkDescriptorPool descriptorPool, VkDescriptorPoolResetFlags flags) {
	return (*fp)(device, descriptorPool, flags);
}
VkResult domVkResetEvent(PFN_vkResetEvent fp, VkDevice device, VkEvent event) {
	return (*fp)(device, event);
}
VkResult domVkResetFences(PFN_vkResetFences fp, VkDevice device, uint32_t fenceCount, const VkFence* pFences) {
	return (*fp)(device, fenceCount, pFences);
}
VkResult domVkSetEvent(PFN_vkSetEvent fp, VkDevice device, VkEvent event) {
	return (*fp)(device, event);
}
VkResult domVkWaitForFences(PFN_vkWaitForFences fp, VkDevice device, uint32_t fenceCount, const VkFence* pFences, VkBool32 waitAll, uint64_t timeout) {
	return (*fp)(device, fenceCount, pFences, waitAll, timeout);
}
void domVkCmdBeginQuery(PFN_vkCmdBeginQuery fp, VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query, VkQueryControlFlags flags) {
	(*fp)(commandBuffer, queryPool, query, flags);
}
void domVkCmdBeginRenderPass(PFN_vkCmdBeginRenderPass fp, VkCommandBuffer commandBuffer, const VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassContents contents) {
	(*fp)(commandBuffer, pRenderPassBegin, contents);
}
void domVkCmdBindDescriptorSets(PFN_vkCmdBindDescriptorSets fp, VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t firstSet, uint32_t descriptorSetCount, const VkDescriptorSet* pDescriptorSets, uint32_t dynamicOffsetCount, const uint32_t* pDynamicOffsets) {
	(*fp)(commandBuffer, pipelineBindPoint, layout, firstSet, descriptorSetCount, pDescriptorSets, dynamicOffsetCount, pDynamicOffsets);
}
void domVkCmdBindIndexBuffer(PFN_vkCmdBindIndexBuffer fp, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkIndexType indexType) {
	(*fp)(commandBuffer, buffer, offset, indexType);
}
void domVkCmdBindPipeline(PFN_vkCmdBindPipeline fp, VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipeline pipeline) {
	(*fp)(commandBuffer, pipelineBindPoint, pipeline);
}
void domVkCmdBindVertexBuffers(PFN_vkCmdBindVertexBuffers fp, VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, const VkBuffer* pBuffers, const VkDeviceSize* pOffsets) {
	(*fp)(commandBuffer, firstBinding, bindingCount, pBuffers, pOffsets);
}
void domVkCmdBlitImage(PFN_vkCmdBlitImage fp, VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkImageBlit* pRegions, VkFilter filter) {
	(*fp)(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions, filter);
}
void domVkCmdClearAttachments(PFN_vkCmdClearAttachments fp, VkCommandBuffer commandBuffer, uint32_t attachmentCount, const VkClearAttachment* pAttachments, uint32_t rectCount, const VkClearRect* pRects) {
	(*fp)(commandBuffer, attachmentCount, pAttachments, rectCount, pRects);
}
void domVkCmdClearColorImage(PFN_vkCmdClearColorImage fp, VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, const VkClearColorValue* pColor, uint32_t rangeCount, const VkImageSubresourceRange* pRanges) {
	(*fp)(commandBuffer, image, imageLayout, pColor, rangeCount, pRanges);
}
void domVkCmdClearDepthStencilImage(PFN_vkCmdClearDepthStencilImage fp, VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, const VkClearDepthStencilValue* pDepthStencil, uint32_t rangeCount, const VkImageSubresourceRange* pRanges) {
	(*fp)(commandBuffer, image, imageLayout, pDepthStencil, rangeCount, pRanges);
}
void domVkCmdCopyBuffer(PFN_vkCmdCopyBuffer fp, VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkBuffer dstBuffer, uint32_t regionCount, const VkBufferCopy* pRegions) {
	(*fp)(commandBuffer, srcBuffer, dstBuffer, regionCount, pRegions);
}
void domVkCmdCopyBufferToImage(PFN_vkCmdCopyBufferToImage fp, VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkBufferImageCopy* pRegions) {
	(*fp)(commandBuffer, srcBuffer, dstImage, dstImageLayout, regionCount, pRegions);
}
void domVkCmdCopyImage(PFN_vkCmdCopyImage fp, VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkImageCopy* pRegions) {
	(*fp)(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions);
}
void domVkCmdCopyImageToBuffer(PFN_vkCmdCopyImageToBuffer fp, VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkBuffer dstBuffer, uint32_t regionCount, const VkBufferImageCopy* pRegions) {
	(*fp)(commandBuffer, srcImage, srcImageLayout, dstBuffer, regionCount, pRegions);
}
void domVkCmdCopyQueryPoolResults(PFN_vkCmdCopyQueryPoolResults fp, VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize stride, VkQueryResultFlags flags) {
	(*fp)(commandBuffer, queryPool, firstQuery, queryCount, dstBuffer, dstOffset, stride, flags);
}
void domVkCmdDispatchIndirect(PFN_vkCmdDispatchIndirect fp, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset) {
	(*fp)(commandBuffer, buffer, offset);
}
void domVkCmdDispatch(PFN_vkCmdDispatch fp, VkCommandBuffer commandBuffer, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	(*fp)(commandBuffer, groupCountX, groupCountY, groupCountZ);
}
void domVkCmdDrawIndexedIndirect(PFN_vkCmdDrawIndexedIndirect fp, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
	(*fp)(commandBuffer, buffer, offset, drawCount, stride);
}
void domVkCmdDrawIndexed(PFN_vkCmdDrawIndexed fp, VkCommandBuffer commandBuffer, uint32_t indexCount, uint32_t instanceCount, uint32_t firstIndex, int32_t vertexOffset, uint32_t firstInstance) {
	(*fp)(commandBuffer, indexCount, instanceCount, firstIndex, vertexOffset, firstInstance);
}
void domVkCmdDrawIndirect(PFN_vkCmdDrawIndirect fp, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
	(*fp)(commandBuffer, buffer, offset, drawCount, stride);
}
void domVkCmdDraw(PFN_vkCmdDraw fp, VkCommandBuffer commandBuffer, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance) {
	(*fp)(commandBuffer, vertexCount, instanceCount, firstVertex, firstInstance);
}
void domVkCmdEndQuery(PFN_vkCmdEndQuery fp, VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query) {
	(*fp)(commandBuffer, queryPool, query);
}
void domVkCmdEndRenderPass(PFN_vkCmdEndRenderPass fp, VkCommandBuffer commandBuffer) {
	(*fp)(commandBuffer);
}
void domVkCmdExecuteCommands(PFN_vkCmdExecuteCommands fp, VkCommandBuffer commandBuffer, uint32_t commandBufferCount, const VkCommandBuffer* pCommandBuffers) {
	(*fp)(commandBuffer, commandBufferCount, pCommandBuffers);
}
void domVkCmdFillBuffer(PFN_vkCmdFillBuffer fp, VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize size, uint32_t data) {
	(*fp)(commandBuffer, dstBuffer, dstOffset, size, data);
}
void domVkCmdNextSubpass(PFN_vkCmdNextSubpass fp, VkCommandBuffer commandBuffer, VkSubpassContents contents) {
	(*fp)(commandBuffer, contents);
}
void domVkCmdPipelineBarrier(PFN_vkCmdPipelineBarrier fp, VkCommandBuffer commandBuffer, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, VkDependencyFlags dependencyFlags, uint32_t memoryBarrierCount, const VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, const VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, const VkImageMemoryBarrier* pImageMemoryBarriers) {
	(*fp)(commandBuffer, srcStageMask, dstStageMask, dependencyFlags, memoryBarrierCount, pMemoryBarriers, bufferMemoryBarrierCount, pBufferMemoryBarriers, imageMemoryBarrierCount, pImageMemoryBarriers);
}
void domVkCmdPushConstants(PFN_vkCmdPushConstants fp, VkCommandBuffer commandBuffer, VkPipelineLayout layout, VkShaderStageFlags stageFlags, uint32_t offset, uint32_t size, const void* pValues) {
	(*fp)(commandBuffer, layout, stageFlags, offset, size, pValues);
}
void domVkCmdResetEvent(PFN_vkCmdResetEvent fp, VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask) {
	(*fp)(commandBuffer, event, stageMask);
}
void domVkCmdResetQueryPool(PFN_vkCmdResetQueryPool fp, VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount) {
	(*fp)(commandBuffer, queryPool, firstQuery, queryCount);
}
void domVkCmdResolveImage(PFN_vkCmdResolveImage fp, VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, const VkImageResolve* pRegions) {
	(*fp)(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions);
}
void domVkCmdSetBlendConstants(PFN_vkCmdSetBlendConstants fp, VkCommandBuffer commandBuffer, const float blendConstants[4]) {
	(*fp)(commandBuffer, blendConstants);
}
void domVkCmdSetDepthBias(PFN_vkCmdSetDepthBias fp, VkCommandBuffer commandBuffer, float depthBiasConstantFactor, float depthBiasClamp, float depthBiasSlopeFactor) {
	(*fp)(commandBuffer, depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor);
}
void domVkCmdSetDepthBounds(PFN_vkCmdSetDepthBounds fp, VkCommandBuffer commandBuffer, float minDepthBounds, float maxDepthBounds) {
	(*fp)(commandBuffer, minDepthBounds, maxDepthBounds);
}
void domVkCmdSetDeviceMask(PFN_vkCmdSetDeviceMask fp, VkCommandBuffer commandBuffer, uint32_t deviceMask) {
	(*fp)(commandBuffer, deviceMask);
}
void domVkCmdSetEvent(PFN_vkCmdSetEvent fp, VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask) {
	(*fp)(commandBuffer, event, stageMask);
}
void domVkCmdSetLineWidth(PFN_vkCmdSetLineWidth fp, VkCommandBuffer commandBuffer, float lineWidth) {
	(*fp)(commandBuffer, lineWidth);
}
void domVkCmdSetScissor(PFN_vkCmdSetScissor fp, VkCommandBuffer commandBuffer, uint32_t firstScissor, uint32_t scissorCount, const VkRect2D* pScissors) {
	(*fp)(commandBuffer, firstScissor, scissorCount, pScissors);
}
void domVkCmdSetStencilCompareMask(PFN_vkCmdSetStencilCompareMask fp, VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t compareMask) {
	(*fp)(commandBuffer, faceMask, compareMask);
}
void domVkCmdSetStencilReference(PFN_vkCmdSetStencilReference fp, VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t reference) {
	(*fp)(commandBuffer, faceMask, reference);
}
void domVkCmdSetStencilWriteMask(PFN_vkCmdSetStencilWriteMask fp, VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t writeMask) {
	(*fp)(commandBuffer, faceMask, writeMask);
}
void domVkCmdSetViewport(PFN_vkCmdSetViewport fp, VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, const VkViewport* pViewports) {
	(*fp)(commandBuffer, firstViewport, viewportCount, pViewports);
}
void domVkCmdUpdateBuffer(PFN_vkCmdUpdateBuffer fp, VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize dataSize, const void* pData) {
	(*fp)(commandBuffer, dstBuffer, dstOffset, dataSize, pData);
}
void domVkCmdWaitEvents(PFN_vkCmdWaitEvents fp, VkCommandBuffer commandBuffer, uint32_t eventCount, const VkEvent* pEvents, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, uint32_t memoryBarrierCount, const VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, const VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, const VkImageMemoryBarrier* pImageMemoryBarriers) {
	(*fp)(commandBuffer, eventCount, pEvents, srcStageMask, dstStageMask, memoryBarrierCount, pMemoryBarriers, bufferMemoryBarrierCount, pBufferMemoryBarriers, imageMemoryBarrierCount, pImageMemoryBarriers);
}
void domVkCmdWriteTimestamp(PFN_vkCmdWriteTimestamp fp, VkCommandBuffer commandBuffer, VkPipelineStageFlagBits pipelineStage, VkQueryPool queryPool, uint32_t query) {
	(*fp)(commandBuffer, pipelineStage, queryPool, query);
}
void domVkDestroyBuffer(PFN_vkDestroyBuffer fp, VkDevice device, VkBuffer buffer, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, buffer, pAllocator);
}
void domVkDestroyBufferView(PFN_vkDestroyBufferView fp, VkDevice device, VkBufferView bufferView, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, bufferView, pAllocator);
}
void domVkDestroyCommandPool(PFN_vkDestroyCommandPool fp, VkDevice device, VkCommandPool commandPool, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, commandPool, pAllocator);
}
void domVkDestroyDescriptorPool(PFN_vkDestroyDescriptorPool fp, VkDevice device, VkDescriptorPool descriptorPool, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, descriptorPool, pAllocator);
}
void domVkDestroyDescriptorSetLayout(PFN_vkDestroyDescriptorSetLayout fp, VkDevice device, VkDescriptorSetLayout descriptorSetLayout, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, descriptorSetLayout, pAllocator);
}
void domVkDestroyDevice(PFN_vkDestroyDevice fp, VkDevice device, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, pAllocator);
}
void domVkDestroyEvent(PFN_vkDestroyEvent fp, VkDevice device, VkEvent event, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, event, pAllocator);
}
void domVkDestroyFence(PFN_vkDestroyFence fp, VkDevice device, VkFence fence, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, fence, pAllocator);
}
void domVkDestroyFramebuffer(PFN_vkDestroyFramebuffer fp, VkDevice device, VkFramebuffer framebuffer, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, framebuffer, pAllocator);
}
void domVkDestroyImage(PFN_vkDestroyImage fp, VkDevice device, VkImage image, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, image, pAllocator);
}
void domVkDestroyImageView(PFN_vkDestroyImageView fp, VkDevice device, VkImageView imageView, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, imageView, pAllocator);
}
void domVkDestroyInstance(PFN_vkDestroyInstance fp, VkInstance instance, const VkAllocationCallbacks* pAllocator) {
	(*fp)(instance, pAllocator);
}
void domVkDestroyPipelineCache(PFN_vkDestroyPipelineCache fp, VkDevice device, VkPipelineCache pipelineCache, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, pipelineCache, pAllocator);
}
void domVkDestroyPipelineLayout(PFN_vkDestroyPipelineLayout fp, VkDevice device, VkPipelineLayout pipelineLayout, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, pipelineLayout, pAllocator);
}
void domVkDestroyPipeline(PFN_vkDestroyPipeline fp, VkDevice device, VkPipeline pipeline, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, pipeline, pAllocator);
}
void domVkDestroyQueryPool(PFN_vkDestroyQueryPool fp, VkDevice device, VkQueryPool queryPool, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, queryPool, pAllocator);
}
void domVkDestroyRenderPass(PFN_vkDestroyRenderPass fp, VkDevice device, VkRenderPass renderPass, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, renderPass, pAllocator);
}
void domVkDestroySampler(PFN_vkDestroySampler fp, VkDevice device, VkSampler sampler, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, sampler, pAllocator);
}
void domVkDestroySemaphore(PFN_vkDestroySemaphore fp, VkDevice device, VkSemaphore semaphore, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, semaphore, pAllocator);
}
void domVkDestroyShaderModule(PFN_vkDestroyShaderModule fp, VkDevice device, VkShaderModule shaderModule, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, shaderModule, pAllocator);
}
void domVkFreeCommandBuffers(PFN_vkFreeCommandBuffers fp, VkDevice device, VkCommandPool commandPool, uint32_t commandBufferCount, const VkCommandBuffer* pCommandBuffers) {
	(*fp)(device, commandPool, commandBufferCount, pCommandBuffers);
}
void domVkFreeMemory(PFN_vkFreeMemory fp, VkDevice device, VkDeviceMemory memory, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, memory, pAllocator);
}
void domVkGetBufferMemoryRequirements(PFN_vkGetBufferMemoryRequirements fp, VkDevice device, VkBuffer buffer, VkMemoryRequirements* pMemoryRequirements) {
	(*fp)(device, buffer, pMemoryRequirements);
}
void domVkGetDeviceMemoryCommitment(PFN_vkGetDeviceMemoryCommitment fp, VkDevice device, VkDeviceMemory memory, VkDeviceSize* pCommittedMemoryInBytes) {
	(*fp)(device, memory, pCommittedMemoryInBytes);
}
void domVkGetDeviceQueue(PFN_vkGetDeviceQueue fp, VkDevice device, uint32_t queueFamilyIndex, uint32_t queueIndex, VkQueue* pQueue) {
	(*fp)(device, queueFamilyIndex, queueIndex, pQueue);
}
void domVkGetDeviceQueue2(PFN_vkGetDeviceQueue2 fp, VkDevice device, const VkDeviceQueueInfo2* pQueueInfo, VkQueue* pQueue) {
	(*fp)(device, pQueueInfo, pQueue);
}
void domVkGetImageMemoryRequirements(PFN_vkGetImageMemoryRequirements fp, VkDevice device, VkImage image, VkMemoryRequirements* pMemoryRequirements) {
	(*fp)(device, image, pMemoryRequirements);
}
void domVkGetImageSparseMemoryRequirements(PFN_vkGetImageSparseMemoryRequirements fp, VkDevice device, VkImage image, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements* pSparseMemoryRequirements) {
	(*fp)(device, image, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}
void domVkGetImageSubresourceLayout(PFN_vkGetImageSubresourceLayout fp, VkDevice device, VkImage image, const VkImageSubresource* pSubresource, VkSubresourceLayout* pLayout) {
	(*fp)(device, image, pSubresource, pLayout);
}
void domVkGetPhysicalDeviceFeatures(PFN_vkGetPhysicalDeviceFeatures fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures* pFeatures) {
	(*fp)(physicalDevice, pFeatures);
}
void domVkGetPhysicalDeviceFormatProperties(PFN_vkGetPhysicalDeviceFormatProperties fp, VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties* pFormatProperties) {
	(*fp)(physicalDevice, format, pFormatProperties);
}
void domVkGetPhysicalDeviceMemoryProperties(PFN_vkGetPhysicalDeviceMemoryProperties fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties* pMemoryProperties) {
	(*fp)(physicalDevice, pMemoryProperties);
}
void domVkGetPhysicalDeviceProperties2(PFN_vkGetPhysicalDeviceProperties2 fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties2* pProperties) {
	(*fp)(physicalDevice, pProperties);
}
void domVkGetPhysicalDeviceProperties(PFN_vkGetPhysicalDeviceProperties fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties* pProperties) {
	(*fp)(physicalDevice, pProperties);
}
void domVkGetPhysicalDeviceQueueFamilyProperties(PFN_vkGetPhysicalDeviceQueueFamilyProperties fp, VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties* pQueueFamilyProperties) {
	(*fp)(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}
void domVkGetPhysicalDeviceSparseImageFormatProperties(PFN_vkGetPhysicalDeviceSparseImageFormatProperties fp, VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkSampleCountFlagBits samples, VkImageUsageFlags usage, VkImageTiling tiling, uint32_t* pPropertyCount, VkSparseImageFormatProperties* pProperties) {
	(*fp)(physicalDevice, format, type, samples, usage, tiling, pPropertyCount, pProperties);
}
void domVkGetRenderAreaGranularity(PFN_vkGetRenderAreaGranularity fp, VkDevice device, VkRenderPass renderPass, VkExtent2D* pGranularity) {
	(*fp)(device, renderPass, pGranularity);
}
void domVkTrimCommandPool(PFN_vkTrimCommandPool fp, VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlags flags) {
	(*fp)(device, commandPool, flags);
}
void domVkUnmapMemory(PFN_vkUnmapMemory fp, VkDevice device, VkDeviceMemory memory) {
	(*fp)(device, memory);
}
void domVkUpdateDescriptorSets(PFN_vkUpdateDescriptorSets fp, VkDevice device, uint32_t descriptorWriteCount, const VkWriteDescriptorSet* pDescriptorWrites, uint32_t descriptorCopyCount, const VkCopyDescriptorSet* pDescriptorCopies) {
	(*fp)(device, descriptorWriteCount, pDescriptorWrites, descriptorCopyCount, pDescriptorCopies);
}
VkBool32 domVkDebugReportCallbackEXT(PFN_vkDebugReportCallbackEXT fp, VkDebugReportFlagsEXT flags, VkDebugReportObjectTypeEXT objectType, uint64_t object, size_t location, int32_t messageCode, const char* pLayerPrefix, const char* pMessage, void* pUserData) {
	return (*fp)(flags, objectType, object, location, messageCode, pLayerPrefix, pMessage, pUserData);
}
VkBool32 domVkDebugUtilsMessengerCallbackEXT(PFN_vkDebugUtilsMessengerCallbackEXT fp, VkDebugUtilsMessageSeverityFlagBitsEXT messageSeverity, VkDebugUtilsMessageTypeFlagsEXT messageTypes, const VkDebugUtilsMessengerCallbackDataEXT* pCallbackData, void* pUserData) {
	return (*fp)(messageSeverity, messageTypes, pCallbackData, pUserData);
}
VkResult domVkAcquireNextImage2KHR(PFN_vkAcquireNextImage2KHR fp, VkDevice device, const VkAcquireNextImageInfoKHR* pAcquireInfo, uint32_t* pImageIndex) {
	return (*fp)(device, pAcquireInfo, pImageIndex);
}
VkResult domVkAcquireNextImageKHR(PFN_vkAcquireNextImageKHR fp, VkDevice device, VkSwapchainKHR swapchain, uint64_t timeout, VkSemaphore semaphore, VkFence fence, uint32_t* pImageIndex) {
	return (*fp)(device, swapchain, timeout, semaphore, fence, pImageIndex);
}
VkResult domVkBindAccelerationStructureMemoryNVX(PFN_vkBindAccelerationStructureMemoryNVX fp, VkDevice device, uint32_t bindInfoCount, const VkBindAccelerationStructureMemoryInfoNVX* pBindInfos) {
	return (*fp)(device, bindInfoCount, pBindInfos);
}
VkResult domVkBindBufferMemory2KHR(PFN_vkBindBufferMemory2KHR fp, VkDevice device, uint32_t bindInfoCount, const VkBindBufferMemoryInfo* pBindInfos) {
	return (*fp)(device, bindInfoCount, pBindInfos);
}
VkResult domVkBindImageMemory2KHR(PFN_vkBindImageMemory2KHR fp, VkDevice device, uint32_t bindInfoCount, const VkBindImageMemoryInfo* pBindInfos) {
	return (*fp)(device, bindInfoCount, pBindInfos);
}
VkResult domVkCompileDeferredNVX(PFN_vkCompileDeferredNVX fp, VkDevice device, VkPipeline pipeline, uint32_t shader) {
	return (*fp)(device, pipeline, shader);
}
VkResult domVkCreateAccelerationStructureNVX(PFN_vkCreateAccelerationStructureNVX fp, VkDevice device, const VkAccelerationStructureCreateInfoNVX* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkAccelerationStructureNVX* pAccelerationStructure) {
	return (*fp)(device, pCreateInfo, pAllocator, pAccelerationStructure);
}
VkResult domVkCreateDebugReportCallbackEXT(PFN_vkCreateDebugReportCallbackEXT fp, VkInstance instance, const VkDebugReportCallbackCreateInfoEXT* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDebugReportCallbackEXT* pCallback) {
	return (*fp)(instance, pCreateInfo, pAllocator, pCallback);
}
VkResult domVkCreateDebugUtilsMessengerEXT(PFN_vkCreateDebugUtilsMessengerEXT fp, VkInstance instance, const VkDebugUtilsMessengerCreateInfoEXT* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDebugUtilsMessengerEXT* pMessenger) {
	return (*fp)(instance, pCreateInfo, pAllocator, pMessenger);
}
VkResult domVkCreateDescriptorUpdateTemplateKHR(PFN_vkCreateDescriptorUpdateTemplateKHR fp, VkDevice device, const VkDescriptorUpdateTemplateCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDescriptorUpdateTemplate* pDescriptorUpdateTemplate) {
	return (*fp)(device, pCreateInfo, pAllocator, pDescriptorUpdateTemplate);
}
VkResult domVkCreateDisplayModeKHR(PFN_vkCreateDisplayModeKHR fp, VkPhysicalDevice physicalDevice, VkDisplayKHR display, const VkDisplayModeCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDisplayModeKHR* pMode) {
	return (*fp)(physicalDevice, display, pCreateInfo, pAllocator, pMode);
}
VkResult domVkCreateDisplayPlaneSurfaceKHR(PFN_vkCreateDisplayPlaneSurfaceKHR fp, VkInstance instance, const VkDisplaySurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return (*fp)(instance, pCreateInfo, pAllocator, pSurface);
}
VkResult domVkCreateIndirectCommandsLayoutNVX(PFN_vkCreateIndirectCommandsLayoutNVX fp, VkDevice device, const VkIndirectCommandsLayoutCreateInfoNVX* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkIndirectCommandsLayoutNVX* pIndirectCommandsLayout) {
	return (*fp)(device, pCreateInfo, pAllocator, pIndirectCommandsLayout);
}
VkResult domVkCreateObjectTableNVX(PFN_vkCreateObjectTableNVX fp, VkDevice device, const VkObjectTableCreateInfoNVX* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkObjectTableNVX* pObjectTable) {
	return (*fp)(device, pCreateInfo, pAllocator, pObjectTable);
}
VkResult domVkCreateRaytracingPipelinesNVX(PFN_vkCreateRaytracingPipelinesNVX fp, VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkRaytracingPipelineCreateInfoNVX* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return (*fp)(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}
VkResult domVkCreateRenderPass2KHR(PFN_vkCreateRenderPass2KHR fp, VkDevice device, const VkRenderPassCreateInfo2KHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass) {
	return (*fp)(device, pCreateInfo, pAllocator, pRenderPass);
}
VkResult domVkCreateSamplerYcbcrConversionKHR(PFN_vkCreateSamplerYcbcrConversionKHR fp, VkDevice device, const VkSamplerYcbcrConversionCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSamplerYcbcrConversion* pYcbcrConversion) {
	return (*fp)(device, pCreateInfo, pAllocator, pYcbcrConversion);
}
VkResult domVkCreateSharedSwapchainsKHR(PFN_vkCreateSharedSwapchainsKHR fp, VkDevice device, uint32_t swapchainCount, const VkSwapchainCreateInfoKHR* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkSwapchainKHR* pSwapchains) {
	return (*fp)(device, swapchainCount, pCreateInfos, pAllocator, pSwapchains);
}
VkResult domVkCreateSwapchainKHR(PFN_vkCreateSwapchainKHR fp, VkDevice device, const VkSwapchainCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSwapchainKHR* pSwapchain) {
	return (*fp)(device, pCreateInfo, pAllocator, pSwapchain);
}
VkResult domVkCreateValidationCacheEXT(PFN_vkCreateValidationCacheEXT fp, VkDevice device, const VkValidationCacheCreateInfoEXT* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkValidationCacheEXT* pValidationCache) {
	return (*fp)(device, pCreateInfo, pAllocator, pValidationCache);
}
VkResult domVkDebugMarkerSetObjectNameEXT(PFN_vkDebugMarkerSetObjectNameEXT fp, VkDevice device, const VkDebugMarkerObjectNameInfoEXT* pNameInfo) {
	return (*fp)(device, pNameInfo);
}
VkResult domVkDebugMarkerSetObjectTagEXT(PFN_vkDebugMarkerSetObjectTagEXT fp, VkDevice device, const VkDebugMarkerObjectTagInfoEXT* pTagInfo) {
	return (*fp)(device, pTagInfo);
}
VkResult domVkDisplayPowerControlEXT(PFN_vkDisplayPowerControlEXT fp, VkDevice device, VkDisplayKHR display, const VkDisplayPowerInfoEXT* pDisplayPowerInfo) {
	return (*fp)(device, display, pDisplayPowerInfo);
}
VkResult domVkEnumeratePhysicalDeviceGroupsKHR(PFN_vkEnumeratePhysicalDeviceGroupsKHR fp, VkInstance instance, uint32_t* pPhysicalDeviceGroupCount, VkPhysicalDeviceGroupProperties* pPhysicalDeviceGroupProperties) {
	return (*fp)(instance, pPhysicalDeviceGroupCount, pPhysicalDeviceGroupProperties);
}
VkResult domVkGetAccelerationStructureHandleNVX(PFN_vkGetAccelerationStructureHandleNVX fp, VkDevice device, VkAccelerationStructureNVX accelerationStructure, size_t dataSize, void* pData) {
	return (*fp)(device, accelerationStructure, dataSize, pData);
}
VkResult domVkGetCalibratedTimestampsEXT(PFN_vkGetCalibratedTimestampsEXT fp, VkDevice device, uint32_t timestampCount, const VkCalibratedTimestampInfoEXT* pTimestampInfos, uint64_t* pTimestamps, uint64_t* pMaxDeviation) {
	return (*fp)(device, timestampCount, pTimestampInfos, pTimestamps, pMaxDeviation);
}
VkResult domVkGetDeviceGroupPresentCapabilitiesKHR(PFN_vkGetDeviceGroupPresentCapabilitiesKHR fp, VkDevice device, VkDeviceGroupPresentCapabilitiesKHR* pDeviceGroupPresentCapabilities) {
	return (*fp)(device, pDeviceGroupPresentCapabilities);
}
VkResult domVkGetDeviceGroupSurfacePresentModesKHR(PFN_vkGetDeviceGroupSurfacePresentModesKHR fp, VkDevice device, VkSurfaceKHR surface, VkDeviceGroupPresentModeFlagsKHR* pModes) {
	return (*fp)(device, surface, pModes);
}
VkResult domVkGetDisplayModeProperties2KHR(PFN_vkGetDisplayModeProperties2KHR fp, VkPhysicalDevice physicalDevice, VkDisplayKHR display, uint32_t* pPropertyCount, VkDisplayModeProperties2KHR* pProperties) {
	return (*fp)(physicalDevice, display, pPropertyCount, pProperties);
}
VkResult domVkGetDisplayModePropertiesKHR(PFN_vkGetDisplayModePropertiesKHR fp, VkPhysicalDevice physicalDevice, VkDisplayKHR display, uint32_t* pPropertyCount, VkDisplayModePropertiesKHR* pProperties) {
	return (*fp)(physicalDevice, display, pPropertyCount, pProperties);
}
VkResult domVkGetDisplayPlaneCapabilities2KHR(PFN_vkGetDisplayPlaneCapabilities2KHR fp, VkPhysicalDevice physicalDevice, const VkDisplayPlaneInfo2KHR* pDisplayPlaneInfo, VkDisplayPlaneCapabilities2KHR* pCapabilities) {
	return (*fp)(physicalDevice, pDisplayPlaneInfo, pCapabilities);
}
VkResult domVkGetDisplayPlaneCapabilitiesKHR(PFN_vkGetDisplayPlaneCapabilitiesKHR fp, VkPhysicalDevice physicalDevice, VkDisplayModeKHR mode, uint32_t planeIndex, VkDisplayPlaneCapabilitiesKHR* pCapabilities) {
	return (*fp)(physicalDevice, mode, planeIndex, pCapabilities);
}
VkResult domVkGetDisplayPlaneSupportedDisplaysKHR(PFN_vkGetDisplayPlaneSupportedDisplaysKHR fp, VkPhysicalDevice physicalDevice, uint32_t planeIndex, uint32_t* pDisplayCount, VkDisplayKHR* pDisplays) {
	return (*fp)(physicalDevice, planeIndex, pDisplayCount, pDisplays);
}
VkResult domVkGetFenceFdKHR(PFN_vkGetFenceFdKHR fp, VkDevice device, const VkFenceGetFdInfoKHR* pGetFdInfo, int* pFd) {
	return (*fp)(device, pGetFdInfo, pFd);
}
VkResult domVkGetImageDrmFormatModifierPropertiesEXT(PFN_vkGetImageDrmFormatModifierPropertiesEXT fp, VkDevice device, VkImage image, VkImageDrmFormatModifierPropertiesEXT* pProperties) {
	return (*fp)(device, image, pProperties);
}
VkResult domVkGetMemoryFdKHR(PFN_vkGetMemoryFdKHR fp, VkDevice device, const VkMemoryGetFdInfoKHR* pGetFdInfo, int* pFd) {
	return (*fp)(device, pGetFdInfo, pFd);
}
VkResult domVkGetMemoryFdPropertiesKHR(PFN_vkGetMemoryFdPropertiesKHR fp, VkDevice device, VkExternalMemoryHandleTypeFlagBits handleType, int fd, VkMemoryFdPropertiesKHR* pMemoryFdProperties) {
	return (*fp)(device, handleType, fd, pMemoryFdProperties);
}
VkResult domVkGetMemoryHostPointerPropertiesEXT(PFN_vkGetMemoryHostPointerPropertiesEXT fp, VkDevice device, VkExternalMemoryHandleTypeFlagBits handleType, const void* pHostPointer, VkMemoryHostPointerPropertiesEXT* pMemoryHostPointerProperties) {
	return (*fp)(device, handleType, pHostPointer, pMemoryHostPointerProperties);
}
VkResult domVkGetPastPresentationTimingGOOGLE(PFN_vkGetPastPresentationTimingGOOGLE fp, VkDevice device, VkSwapchainKHR swapchain, uint32_t* pPresentationTimingCount, VkPastPresentationTimingGOOGLE* pPresentationTimings) {
	return (*fp)(device, swapchain, pPresentationTimingCount, pPresentationTimings);
}
VkResult domVkGetPhysicalDeviceCalibrateableTimeDomainsEXT(PFN_vkGetPhysicalDeviceCalibrateableTimeDomainsEXT fp, VkPhysicalDevice physicalDevice, uint32_t* pTimeDomainCount, VkTimeDomainEXT* pTimeDomains) {
	return (*fp)(physicalDevice, pTimeDomainCount, pTimeDomains);
}
VkResult domVkGetPhysicalDeviceDisplayPlaneProperties2KHR(PFN_vkGetPhysicalDeviceDisplayPlaneProperties2KHR fp, VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkDisplayPlaneProperties2KHR* pProperties) {
	return (*fp)(physicalDevice, pPropertyCount, pProperties);
}
VkResult domVkGetPhysicalDeviceDisplayPlanePropertiesKHR(PFN_vkGetPhysicalDeviceDisplayPlanePropertiesKHR fp, VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkDisplayPlanePropertiesKHR* pProperties) {
	return (*fp)(physicalDevice, pPropertyCount, pProperties);
}
VkResult domVkGetPhysicalDeviceDisplayProperties2KHR(PFN_vkGetPhysicalDeviceDisplayProperties2KHR fp, VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkDisplayProperties2KHR* pProperties) {
	return (*fp)(physicalDevice, pPropertyCount, pProperties);
}
VkResult domVkGetPhysicalDeviceDisplayPropertiesKHR(PFN_vkGetPhysicalDeviceDisplayPropertiesKHR fp, VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkDisplayPropertiesKHR* pProperties) {
	return (*fp)(physicalDevice, pPropertyCount, pProperties);
}
VkResult domVkGetPhysicalDeviceExternalImageFormatPropertiesNV(PFN_vkGetPhysicalDeviceExternalImageFormatPropertiesNV fp, VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkImageTiling tiling, VkImageUsageFlags usage, VkImageCreateFlags flags, VkExternalMemoryHandleTypeFlagsNV externalHandleType, VkExternalImageFormatPropertiesNV* pExternalImageFormatProperties) {
	return (*fp)(physicalDevice, format, type, tiling, usage, flags, externalHandleType, pExternalImageFormatProperties);
}
VkResult domVkGetPhysicalDeviceImageFormatProperties2KHR(PFN_vkGetPhysicalDeviceImageFormatProperties2KHR fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceImageFormatInfo2* pImageFormatInfo, VkImageFormatProperties2* pImageFormatProperties) {
	return (*fp)(physicalDevice, pImageFormatInfo, pImageFormatProperties);
}
VkResult domVkGetPhysicalDevicePresentRectanglesKHR(PFN_vkGetPhysicalDevicePresentRectanglesKHR fp, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pRectCount, VkRect2D* pRects) {
	return (*fp)(physicalDevice, surface, pRectCount, pRects);
}
VkResult domVkGetPhysicalDeviceSurfaceCapabilities2EXT(PFN_vkGetPhysicalDeviceSurfaceCapabilities2EXT fp, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, VkSurfaceCapabilities2EXT* pSurfaceCapabilities) {
	return (*fp)(physicalDevice, surface, pSurfaceCapabilities);
}
VkResult domVkGetPhysicalDeviceSurfaceCapabilities2KHR(PFN_vkGetPhysicalDeviceSurfaceCapabilities2KHR fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceSurfaceInfo2KHR* pSurfaceInfo, VkSurfaceCapabilities2KHR* pSurfaceCapabilities) {
	return (*fp)(physicalDevice, pSurfaceInfo, pSurfaceCapabilities);
}
VkResult domVkGetPhysicalDeviceSurfaceCapabilitiesKHR(PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR fp, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, VkSurfaceCapabilitiesKHR* pSurfaceCapabilities) {
	return (*fp)(physicalDevice, surface, pSurfaceCapabilities);
}
VkResult domVkGetPhysicalDeviceSurfaceFormats2KHR(PFN_vkGetPhysicalDeviceSurfaceFormats2KHR fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceSurfaceInfo2KHR* pSurfaceInfo, uint32_t* pSurfaceFormatCount, VkSurfaceFormat2KHR* pSurfaceFormats) {
	return (*fp)(physicalDevice, pSurfaceInfo, pSurfaceFormatCount, pSurfaceFormats);
}
VkResult domVkGetPhysicalDeviceSurfaceFormatsKHR(PFN_vkGetPhysicalDeviceSurfaceFormatsKHR fp, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pSurfaceFormatCount, VkSurfaceFormatKHR* pSurfaceFormats) {
	return (*fp)(physicalDevice, surface, pSurfaceFormatCount, pSurfaceFormats);
}
VkResult domVkGetPhysicalDeviceSurfacePresentModesKHR(PFN_vkGetPhysicalDeviceSurfacePresentModesKHR fp, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pPresentModeCount, VkPresentModeKHR* pPresentModes) {
	return (*fp)(physicalDevice, surface, pPresentModeCount, pPresentModes);
}
VkResult domVkGetPhysicalDeviceSurfaceSupportKHR(PFN_vkGetPhysicalDeviceSurfaceSupportKHR fp, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, VkSurfaceKHR surface, VkBool32* pSupported) {
	return (*fp)(physicalDevice, queueFamilyIndex, surface, pSupported);
}
VkResult domVkGetRaytracingShaderHandlesNVX(PFN_vkGetRaytracingShaderHandlesNVX fp, VkDevice device, VkPipeline pipeline, uint32_t firstGroup, uint32_t groupCount, size_t dataSize, void* pData) {
	return (*fp)(device, pipeline, firstGroup, groupCount, dataSize, pData);
}
VkResult domVkGetRefreshCycleDurationGOOGLE(PFN_vkGetRefreshCycleDurationGOOGLE fp, VkDevice device, VkSwapchainKHR swapchain, VkRefreshCycleDurationGOOGLE* pDisplayTimingProperties) {
	return (*fp)(device, swapchain, pDisplayTimingProperties);
}
VkResult domVkGetSemaphoreFdKHR(PFN_vkGetSemaphoreFdKHR fp, VkDevice device, const VkSemaphoreGetFdInfoKHR* pGetFdInfo, int* pFd) {
	return (*fp)(device, pGetFdInfo, pFd);
}
VkResult domVkGetShaderInfoAMD(PFN_vkGetShaderInfoAMD fp, VkDevice device, VkPipeline pipeline, VkShaderStageFlagBits shaderStage, VkShaderInfoTypeAMD infoType, size_t* pInfoSize, void* pInfo) {
	return (*fp)(device, pipeline, shaderStage, infoType, pInfoSize, pInfo);
}
VkResult domVkGetSwapchainCounterEXT(PFN_vkGetSwapchainCounterEXT fp, VkDevice device, VkSwapchainKHR swapchain, VkSurfaceCounterFlagBitsEXT counter, uint64_t* pCounterValue) {
	return (*fp)(device, swapchain, counter, pCounterValue);
}
VkResult domVkGetSwapchainImagesKHR(PFN_vkGetSwapchainImagesKHR fp, VkDevice device, VkSwapchainKHR swapchain, uint32_t* pSwapchainImageCount, VkImage* pSwapchainImages) {
	return (*fp)(device, swapchain, pSwapchainImageCount, pSwapchainImages);
}
VkResult domVkGetSwapchainStatusKHR(PFN_vkGetSwapchainStatusKHR fp, VkDevice device, VkSwapchainKHR swapchain) {
	return (*fp)(device, swapchain);
}
VkResult domVkGetValidationCacheDataEXT(PFN_vkGetValidationCacheDataEXT fp, VkDevice device, VkValidationCacheEXT validationCache, size_t* pDataSize, void* pData) {
	return (*fp)(device, validationCache, pDataSize, pData);
}
VkResult domVkImportFenceFdKHR(PFN_vkImportFenceFdKHR fp, VkDevice device, const VkImportFenceFdInfoKHR* pImportFenceFdInfo) {
	return (*fp)(device, pImportFenceFdInfo);
}
VkResult domVkImportSemaphoreFdKHR(PFN_vkImportSemaphoreFdKHR fp, VkDevice device, const VkImportSemaphoreFdInfoKHR* pImportSemaphoreFdInfo) {
	return (*fp)(device, pImportSemaphoreFdInfo);
}
VkResult domVkMergeValidationCachesEXT(PFN_vkMergeValidationCachesEXT fp, VkDevice device, VkValidationCacheEXT dstCache, uint32_t srcCacheCount, const VkValidationCacheEXT* pSrcCaches) {
	return (*fp)(device, dstCache, srcCacheCount, pSrcCaches);
}
VkResult domVkQueuePresentKHR(PFN_vkQueuePresentKHR fp, VkQueue queue, const VkPresentInfoKHR* pPresentInfo) {
	return (*fp)(queue, pPresentInfo);
}
VkResult domVkRegisterDeviceEventEXT(PFN_vkRegisterDeviceEventEXT fp, VkDevice device, const VkDeviceEventInfoEXT* pDeviceEventInfo, const VkAllocationCallbacks* pAllocator, VkFence* pFence) {
	return (*fp)(device, pDeviceEventInfo, pAllocator, pFence);
}
VkResult domVkRegisterDisplayEventEXT(PFN_vkRegisterDisplayEventEXT fp, VkDevice device, VkDisplayKHR display, const VkDisplayEventInfoEXT* pDisplayEventInfo, const VkAllocationCallbacks* pAllocator, VkFence* pFence) {
	return (*fp)(device, display, pDisplayEventInfo, pAllocator, pFence);
}
VkResult domVkRegisterObjectsNVX(PFN_vkRegisterObjectsNVX fp, VkDevice device, VkObjectTableNVX objectTable, uint32_t objectCount, const VkObjectTableEntryNVX* const*    ppObjectTableEntries, const uint32_t* pObjectIndices) {
	return (*fp)(device, objectTable, objectCount, ppObjectTableEntries, pObjectIndices);
}
VkResult domVkReleaseDisplayEXT(PFN_vkReleaseDisplayEXT fp, VkPhysicalDevice physicalDevice, VkDisplayKHR display) {
	return (*fp)(physicalDevice, display);
}
VkResult domVkSetDebugUtilsObjectNameEXT(PFN_vkSetDebugUtilsObjectNameEXT fp, VkDevice device, const VkDebugUtilsObjectNameInfoEXT* pNameInfo) {
	return (*fp)(device, pNameInfo);
}
VkResult domVkSetDebugUtilsObjectTagEXT(PFN_vkSetDebugUtilsObjectTagEXT fp, VkDevice device, const VkDebugUtilsObjectTagInfoEXT* pTagInfo) {
	return (*fp)(device, pTagInfo);
}
VkResult domVkUnregisterObjectsNVX(PFN_vkUnregisterObjectsNVX fp, VkDevice device, VkObjectTableNVX objectTable, uint32_t objectCount, const VkObjectEntryTypeNVX* pObjectEntryTypes, const uint32_t* pObjectIndices) {
	return (*fp)(device, objectTable, objectCount, pObjectEntryTypes, pObjectIndices);
}
void domVkCmdBeginConditionalRenderingEXT(PFN_vkCmdBeginConditionalRenderingEXT fp, VkCommandBuffer commandBuffer, const VkConditionalRenderingBeginInfoEXT* pConditionalRenderingBegin) {
	(*fp)(commandBuffer, pConditionalRenderingBegin);
}
void domVkCmdBeginDebugUtilsLabelEXT(PFN_vkCmdBeginDebugUtilsLabelEXT fp, VkCommandBuffer commandBuffer, const VkDebugUtilsLabelEXT* pLabelInfo) {
	(*fp)(commandBuffer, pLabelInfo);
}
void domVkCmdBeginQueryIndexedEXT(PFN_vkCmdBeginQueryIndexedEXT fp, VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query, VkQueryControlFlags flags, uint32_t index) {
	(*fp)(commandBuffer, queryPool, query, flags, index);
}
void domVkCmdBeginRenderPass2KHR(PFN_vkCmdBeginRenderPass2KHR fp, VkCommandBuffer commandBuffer, const VkRenderPassBeginInfo*      pRenderPassBegin, const VkSubpassBeginInfoKHR*      pSubpassBeginInfo) {
	(*fp)(commandBuffer, pRenderPassBegin, pSubpassBeginInfo);
}
void domVkCmdBeginTransformFeedbackEXT(PFN_vkCmdBeginTransformFeedbackEXT fp, VkCommandBuffer commandBuffer, uint32_t firstCounterBuffer, uint32_t counterBufferCount, const VkBuffer* pCounterBuffers, const VkDeviceSize* pCounterBufferOffsets) {
	(*fp)(commandBuffer, firstCounterBuffer, counterBufferCount, pCounterBuffers, pCounterBufferOffsets);
}
void domVkCmdBindShadingRateImageNV(PFN_vkCmdBindShadingRateImageNV fp, VkCommandBuffer commandBuffer, VkImageView imageView, VkImageLayout imageLayout) {
	(*fp)(commandBuffer, imageView, imageLayout);
}
void domVkCmdBindTransformFeedbackBuffersEXT(PFN_vkCmdBindTransformFeedbackBuffersEXT fp, VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, const VkBuffer* pBuffers, const VkDeviceSize* pOffsets, const VkDeviceSize* pSizes) {
	(*fp)(commandBuffer, firstBinding, bindingCount, pBuffers, pOffsets, pSizes);
}
void domVkCmdBuildAccelerationStructureNVX(PFN_vkCmdBuildAccelerationStructureNVX fp, VkCommandBuffer commandBuffer, VkAccelerationStructureTypeNVX type, uint32_t instanceCount, VkBuffer instanceData, VkDeviceSize instanceOffset, uint32_t geometryCount, const VkGeometryNVX* pGeometries, VkBuildAccelerationStructureFlagsNVX flags, VkBool32 update, VkAccelerationStructureNVX dst, VkAccelerationStructureNVX src, VkBuffer scratch, VkDeviceSize scratchOffset) {
	(*fp)(commandBuffer, type, instanceCount, instanceData, instanceOffset, geometryCount, pGeometries, flags, update, dst, src, scratch, scratchOffset);
}
void domVkCmdCopyAccelerationStructureNVX(PFN_vkCmdCopyAccelerationStructureNVX fp, VkCommandBuffer commandBuffer, VkAccelerationStructureNVX dst, VkAccelerationStructureNVX src, VkCopyAccelerationStructureModeNVX mode) {
	(*fp)(commandBuffer, dst, src, mode);
}
void domVkCmdDebugMarkerBeginEXT(PFN_vkCmdDebugMarkerBeginEXT fp, VkCommandBuffer commandBuffer, const VkDebugMarkerMarkerInfoEXT* pMarkerInfo) {
	(*fp)(commandBuffer, pMarkerInfo);
}
void domVkCmdDebugMarkerEndEXT(PFN_vkCmdDebugMarkerEndEXT fp, VkCommandBuffer commandBuffer) {
	(*fp)(commandBuffer);
}
void domVkCmdDebugMarkerInsertEXT(PFN_vkCmdDebugMarkerInsertEXT fp, VkCommandBuffer commandBuffer, const VkDebugMarkerMarkerInfoEXT* pMarkerInfo) {
	(*fp)(commandBuffer, pMarkerInfo);
}
void domVkCmdDispatchBaseKHR(PFN_vkCmdDispatchBaseKHR fp, VkCommandBuffer commandBuffer, uint32_t baseGroupX, uint32_t baseGroupY, uint32_t baseGroupZ, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	(*fp)(commandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ);
}
void domVkCmdDrawIndexedIndirectCountAMD(PFN_vkCmdDrawIndexedIndirectCountAMD fp, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	(*fp)(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
void domVkCmdDrawIndexedIndirectCountKHR(PFN_vkCmdDrawIndexedIndirectCountKHR fp, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	(*fp)(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
void domVkCmdDrawIndirectByteCountEXT(PFN_vkCmdDrawIndirectByteCountEXT fp, VkCommandBuffer commandBuffer, uint32_t instanceCount, uint32_t firstInstance, VkBuffer counterBuffer, VkDeviceSize counterBufferOffset, uint32_t counterOffset, uint32_t vertexStride) {
	(*fp)(commandBuffer, instanceCount, firstInstance, counterBuffer, counterBufferOffset, counterOffset, vertexStride);
}
void domVkCmdDrawIndirectCountAMD(PFN_vkCmdDrawIndirectCountAMD fp, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	(*fp)(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
void domVkCmdDrawIndirectCountKHR(PFN_vkCmdDrawIndirectCountKHR fp, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	(*fp)(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
void domVkCmdDrawMeshTasksIndirectCountNV(PFN_vkCmdDrawMeshTasksIndirectCountNV fp, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	(*fp)(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
void domVkCmdDrawMeshTasksIndirectNV(PFN_vkCmdDrawMeshTasksIndirectNV fp, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
	(*fp)(commandBuffer, buffer, offset, drawCount, stride);
}
void domVkCmdDrawMeshTasksNV(PFN_vkCmdDrawMeshTasksNV fp, VkCommandBuffer commandBuffer, uint32_t taskCount, uint32_t firstTask) {
	(*fp)(commandBuffer, taskCount, firstTask);
}
void domVkCmdEndConditionalRenderingEXT(PFN_vkCmdEndConditionalRenderingEXT fp, VkCommandBuffer commandBuffer) {
	(*fp)(commandBuffer);
}
void domVkCmdEndDebugUtilsLabelEXT(PFN_vkCmdEndDebugUtilsLabelEXT fp, VkCommandBuffer commandBuffer) {
	(*fp)(commandBuffer);
}
void domVkCmdEndQueryIndexedEXT(PFN_vkCmdEndQueryIndexedEXT fp, VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query, uint32_t index) {
	(*fp)(commandBuffer, queryPool, query, index);
}
void domVkCmdEndRenderPass2KHR(PFN_vkCmdEndRenderPass2KHR fp, VkCommandBuffer commandBuffer, const VkSubpassEndInfoKHR*        pSubpassEndInfo) {
	(*fp)(commandBuffer, pSubpassEndInfo);
}
void domVkCmdEndTransformFeedbackEXT(PFN_vkCmdEndTransformFeedbackEXT fp, VkCommandBuffer commandBuffer, uint32_t firstCounterBuffer, uint32_t counterBufferCount, const VkBuffer* pCounterBuffers, const VkDeviceSize* pCounterBufferOffsets) {
	(*fp)(commandBuffer, firstCounterBuffer, counterBufferCount, pCounterBuffers, pCounterBufferOffsets);
}
void domVkCmdInsertDebugUtilsLabelEXT(PFN_vkCmdInsertDebugUtilsLabelEXT fp, VkCommandBuffer commandBuffer, const VkDebugUtilsLabelEXT* pLabelInfo) {
	(*fp)(commandBuffer, pLabelInfo);
}
void domVkCmdNextSubpass2KHR(PFN_vkCmdNextSubpass2KHR fp, VkCommandBuffer commandBuffer, const VkSubpassBeginInfoKHR*      pSubpassBeginInfo, const VkSubpassEndInfoKHR*        pSubpassEndInfo) {
	(*fp)(commandBuffer, pSubpassBeginInfo, pSubpassEndInfo);
}
void domVkCmdProcessCommandsNVX(PFN_vkCmdProcessCommandsNVX fp, VkCommandBuffer commandBuffer, const VkCmdProcessCommandsInfoNVX* pProcessCommandsInfo) {
	(*fp)(commandBuffer, pProcessCommandsInfo);
}
void domVkCmdPushDescriptorSetKHR(PFN_vkCmdPushDescriptorSetKHR fp, VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t set, uint32_t descriptorWriteCount, const VkWriteDescriptorSet* pDescriptorWrites) {
	(*fp)(commandBuffer, pipelineBindPoint, layout, set, descriptorWriteCount, pDescriptorWrites);
}
void domVkCmdPushDescriptorSetWithTemplateKHR(PFN_vkCmdPushDescriptorSetWithTemplateKHR fp, VkCommandBuffer commandBuffer, VkDescriptorUpdateTemplate descriptorUpdateTemplate, VkPipelineLayout layout, uint32_t set, const void* pData) {
	(*fp)(commandBuffer, descriptorUpdateTemplate, layout, set, pData);
}
void domVkCmdReserveSpaceForCommandsNVX(PFN_vkCmdReserveSpaceForCommandsNVX fp, VkCommandBuffer commandBuffer, const VkCmdReserveSpaceForCommandsInfoNVX* pReserveSpaceInfo) {
	(*fp)(commandBuffer, pReserveSpaceInfo);
}
void domVkCmdSetCheckpointNV(PFN_vkCmdSetCheckpointNV fp, VkCommandBuffer commandBuffer, const void* pCheckpointMarker) {
	(*fp)(commandBuffer, pCheckpointMarker);
}
void domVkCmdSetCoarseSampleOrderNV(PFN_vkCmdSetCoarseSampleOrderNV fp, VkCommandBuffer commandBuffer, VkCoarseSampleOrderTypeNV sampleOrderType, uint32_t customSampleOrderCount, const VkCoarseSampleOrderCustomNV* pCustomSampleOrders) {
	(*fp)(commandBuffer, sampleOrderType, customSampleOrderCount, pCustomSampleOrders);
}
void domVkCmdSetDeviceMaskKHR(PFN_vkCmdSetDeviceMaskKHR fp, VkCommandBuffer commandBuffer, uint32_t deviceMask) {
	(*fp)(commandBuffer, deviceMask);
}
void domVkCmdSetDiscardRectangleEXT(PFN_vkCmdSetDiscardRectangleEXT fp, VkCommandBuffer commandBuffer, uint32_t firstDiscardRectangle, uint32_t discardRectangleCount, const VkRect2D* pDiscardRectangles) {
	(*fp)(commandBuffer, firstDiscardRectangle, discardRectangleCount, pDiscardRectangles);
}
void domVkCmdSetExclusiveScissorNV(PFN_vkCmdSetExclusiveScissorNV fp, VkCommandBuffer commandBuffer, uint32_t firstExclusiveScissor, uint32_t exclusiveScissorCount, const VkRect2D* pExclusiveScissors) {
	(*fp)(commandBuffer, firstExclusiveScissor, exclusiveScissorCount, pExclusiveScissors);
}
void domVkCmdSetSampleLocationsEXT(PFN_vkCmdSetSampleLocationsEXT fp, VkCommandBuffer commandBuffer, const VkSampleLocationsInfoEXT* pSampleLocationsInfo) {
	(*fp)(commandBuffer, pSampleLocationsInfo);
}
void domVkCmdSetViewportShadingRatePaletteNV(PFN_vkCmdSetViewportShadingRatePaletteNV fp, VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, const VkShadingRatePaletteNV* pShadingRatePalettes) {
	(*fp)(commandBuffer, firstViewport, viewportCount, pShadingRatePalettes);
}
void domVkCmdSetViewportWScalingNV(PFN_vkCmdSetViewportWScalingNV fp, VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, const VkViewportWScalingNV* pViewportWScalings) {
	(*fp)(commandBuffer, firstViewport, viewportCount, pViewportWScalings);
}
void domVkCmdTraceRaysNVX(PFN_vkCmdTraceRaysNVX fp, VkCommandBuffer commandBuffer, VkBuffer raygenShaderBindingTableBuffer, VkDeviceSize raygenShaderBindingOffset, VkBuffer missShaderBindingTableBuffer, VkDeviceSize missShaderBindingOffset, VkDeviceSize missShaderBindingStride, VkBuffer hitShaderBindingTableBuffer, VkDeviceSize hitShaderBindingOffset, VkDeviceSize hitShaderBindingStride, uint32_t width, uint32_t height) {
	(*fp)(commandBuffer, raygenShaderBindingTableBuffer, raygenShaderBindingOffset, missShaderBindingTableBuffer, missShaderBindingOffset, missShaderBindingStride, hitShaderBindingTableBuffer, hitShaderBindingOffset, hitShaderBindingStride, width, height);
}
void domVkCmdWriteAccelerationStructurePropertiesNVX(PFN_vkCmdWriteAccelerationStructurePropertiesNVX fp, VkCommandBuffer commandBuffer, VkAccelerationStructureNVX accelerationStructure, VkQueryType queryType, VkQueryPool queryPool, uint32_t query) {
	(*fp)(commandBuffer, accelerationStructure, queryType, queryPool, query);
}
void domVkCmdWriteBufferMarkerAMD(PFN_vkCmdWriteBufferMarkerAMD fp, VkCommandBuffer commandBuffer, VkPipelineStageFlagBits pipelineStage, VkBuffer dstBuffer, VkDeviceSize dstOffset, uint32_t marker) {
	(*fp)(commandBuffer, pipelineStage, dstBuffer, dstOffset, marker);
}
void domVkDebugReportMessageEXT(PFN_vkDebugReportMessageEXT fp, VkInstance instance, VkDebugReportFlagsEXT flags, VkDebugReportObjectTypeEXT objectType, uint64_t object, size_t location, int32_t messageCode, const char* pLayerPrefix, const char* pMessage) {
	(*fp)(instance, flags, objectType, object, location, messageCode, pLayerPrefix, pMessage);
}
void domVkDestroyAccelerationStructureNVX(PFN_vkDestroyAccelerationStructureNVX fp, VkDevice device, VkAccelerationStructureNVX accelerationStructure, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, accelerationStructure, pAllocator);
}
void domVkDestroyDebugReportCallbackEXT(PFN_vkDestroyDebugReportCallbackEXT fp, VkInstance instance, VkDebugReportCallbackEXT callback, const VkAllocationCallbacks* pAllocator) {
	(*fp)(instance, callback, pAllocator);
}
void domVkDestroyDebugUtilsMessengerEXT(PFN_vkDestroyDebugUtilsMessengerEXT fp, VkInstance instance, VkDebugUtilsMessengerEXT messenger, const VkAllocationCallbacks* pAllocator) {
	(*fp)(instance, messenger, pAllocator);
}
void domVkDestroyDescriptorUpdateTemplateKHR(PFN_vkDestroyDescriptorUpdateTemplateKHR fp, VkDevice device, VkDescriptorUpdateTemplate descriptorUpdateTemplate, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, descriptorUpdateTemplate, pAllocator);
}
void domVkDestroyIndirectCommandsLayoutNVX(PFN_vkDestroyIndirectCommandsLayoutNVX fp, VkDevice device, VkIndirectCommandsLayoutNVX indirectCommandsLayout, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, indirectCommandsLayout, pAllocator);
}
void domVkDestroyObjectTableNVX(PFN_vkDestroyObjectTableNVX fp, VkDevice device, VkObjectTableNVX objectTable, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, objectTable, pAllocator);
}
void domVkDestroySamplerYcbcrConversionKHR(PFN_vkDestroySamplerYcbcrConversionKHR fp, VkDevice device, VkSamplerYcbcrConversion ycbcrConversion, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, ycbcrConversion, pAllocator);
}
void domVkDestroySurfaceKHR(PFN_vkDestroySurfaceKHR fp, VkInstance instance, VkSurfaceKHR surface, const VkAllocationCallbacks* pAllocator) {
	(*fp)(instance, surface, pAllocator);
}
void domVkDestroySwapchainKHR(PFN_vkDestroySwapchainKHR fp, VkDevice device, VkSwapchainKHR swapchain, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, swapchain, pAllocator);
}
void domVkDestroyValidationCacheEXT(PFN_vkDestroyValidationCacheEXT fp, VkDevice device, VkValidationCacheEXT validationCache, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, validationCache, pAllocator);
}
void domVkGetAccelerationStructureMemoryRequirementsNVX(PFN_vkGetAccelerationStructureMemoryRequirementsNVX fp, VkDevice device, const VkAccelerationStructureMemoryRequirementsInfoNVX* pInfo, VkMemoryRequirements2KHR* pMemoryRequirements) {
	(*fp)(device, pInfo, pMemoryRequirements);
}
void domVkGetAccelerationStructureScratchMemoryRequirementsNVX(PFN_vkGetAccelerationStructureScratchMemoryRequirementsNVX fp, VkDevice device, const VkAccelerationStructureMemoryRequirementsInfoNVX* pInfo, VkMemoryRequirements2KHR* pMemoryRequirements) {
	(*fp)(device, pInfo, pMemoryRequirements);
}
void domVkGetBufferMemoryRequirements2KHR(PFN_vkGetBufferMemoryRequirements2KHR fp, VkDevice device, const VkBufferMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	(*fp)(device, pInfo, pMemoryRequirements);
}
void domVkGetDescriptorSetLayoutSupportKHR(PFN_vkGetDescriptorSetLayoutSupportKHR fp, VkDevice device, const VkDescriptorSetLayoutCreateInfo* pCreateInfo, VkDescriptorSetLayoutSupport* pSupport) {
	(*fp)(device, pCreateInfo, pSupport);
}
void domVkGetDeviceGroupPeerMemoryFeaturesKHR(PFN_vkGetDeviceGroupPeerMemoryFeaturesKHR fp, VkDevice device, uint32_t heapIndex, uint32_t localDeviceIndex, uint32_t remoteDeviceIndex, VkPeerMemoryFeatureFlags* pPeerMemoryFeatures) {
	(*fp)(device, heapIndex, localDeviceIndex, remoteDeviceIndex, pPeerMemoryFeatures);
}
void domVkGetImageMemoryRequirements2KHR(PFN_vkGetImageMemoryRequirements2KHR fp, VkDevice device, const VkImageMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	(*fp)(device, pInfo, pMemoryRequirements);
}
void domVkGetImageSparseMemoryRequirements2KHR(PFN_vkGetImageSparseMemoryRequirements2KHR fp, VkDevice device, const VkImageSparseMemoryRequirementsInfo2* pInfo, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2* pSparseMemoryRequirements) {
	(*fp)(device, pInfo, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}
void domVkGetPhysicalDeviceExternalBufferPropertiesKHR(PFN_vkGetPhysicalDeviceExternalBufferPropertiesKHR fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalBufferInfo* pExternalBufferInfo, VkExternalBufferProperties* pExternalBufferProperties) {
	(*fp)(physicalDevice, pExternalBufferInfo, pExternalBufferProperties);
}
void domVkGetPhysicalDeviceExternalFencePropertiesKHR(PFN_vkGetPhysicalDeviceExternalFencePropertiesKHR fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalFenceInfo* pExternalFenceInfo, VkExternalFenceProperties* pExternalFenceProperties) {
	(*fp)(physicalDevice, pExternalFenceInfo, pExternalFenceProperties);
}
void domVkGetPhysicalDeviceExternalSemaphorePropertiesKHR(PFN_vkGetPhysicalDeviceExternalSemaphorePropertiesKHR fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalSemaphoreInfo* pExternalSemaphoreInfo, VkExternalSemaphoreProperties* pExternalSemaphoreProperties) {
	(*fp)(physicalDevice, pExternalSemaphoreInfo, pExternalSemaphoreProperties);
}
void domVkGetPhysicalDeviceFeatures2KHR(PFN_vkGetPhysicalDeviceFeatures2KHR fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures2* pFeatures) {
	(*fp)(physicalDevice, pFeatures);
}
void domVkGetPhysicalDeviceFormatProperties2KHR(PFN_vkGetPhysicalDeviceFormatProperties2KHR fp, VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties2* pFormatProperties) {
	(*fp)(physicalDevice, format, pFormatProperties);
}
void domVkGetPhysicalDeviceGeneratedCommandsPropertiesNVX(PFN_vkGetPhysicalDeviceGeneratedCommandsPropertiesNVX fp, VkPhysicalDevice physicalDevice, VkDeviceGeneratedCommandsFeaturesNVX* pFeatures, VkDeviceGeneratedCommandsLimitsNVX* pLimits) {
	(*fp)(physicalDevice, pFeatures, pLimits);
}
void domVkGetPhysicalDeviceMemoryProperties2KHR(PFN_vkGetPhysicalDeviceMemoryProperties2KHR fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties2* pMemoryProperties) {
	(*fp)(physicalDevice, pMemoryProperties);
}
void domVkGetPhysicalDeviceMultisamplePropertiesEXT(PFN_vkGetPhysicalDeviceMultisamplePropertiesEXT fp, VkPhysicalDevice physicalDevice, VkSampleCountFlagBits samples, VkMultisamplePropertiesEXT* pMultisampleProperties) {
	(*fp)(physicalDevice, samples, pMultisampleProperties);
}
void domVkGetPhysicalDeviceProperties2KHR(PFN_vkGetPhysicalDeviceProperties2KHR fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties2* pProperties) {
	(*fp)(physicalDevice, pProperties);
}
void domVkGetPhysicalDeviceQueueFamilyProperties2KHR(PFN_vkGetPhysicalDeviceQueueFamilyProperties2KHR fp, VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties2* pQueueFamilyProperties) {
	(*fp)(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}
void domVkGetPhysicalDeviceSparseImageFormatProperties2KHR(PFN_vkGetPhysicalDeviceSparseImageFormatProperties2KHR fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceSparseImageFormatInfo2* pFormatInfo, uint32_t* pPropertyCount, VkSparseImageFormatProperties2* pProperties) {
	(*fp)(physicalDevice, pFormatInfo, pPropertyCount, pProperties);
}
void domVkGetQueueCheckpointDataNV(PFN_vkGetQueueCheckpointDataNV fp, VkQueue queue, uint32_t* pCheckpointDataCount, VkCheckpointDataNV* pCheckpointData) {
	(*fp)(queue, pCheckpointDataCount, pCheckpointData);
}
void domVkQueueBeginDebugUtilsLabelEXT(PFN_vkQueueBeginDebugUtilsLabelEXT fp, VkQueue queue, const VkDebugUtilsLabelEXT* pLabelInfo) {
	(*fp)(queue, pLabelInfo);
}
void domVkQueueEndDebugUtilsLabelEXT(PFN_vkQueueEndDebugUtilsLabelEXT fp, VkQueue queue) {
	(*fp)(queue);
}
void domVkQueueInsertDebugUtilsLabelEXT(PFN_vkQueueInsertDebugUtilsLabelEXT fp, VkQueue queue, const VkDebugUtilsLabelEXT* pLabelInfo) {
	(*fp)(queue, pLabelInfo);
}
void domVkSetHdrMetadataEXT(PFN_vkSetHdrMetadataEXT fp, VkDevice device, uint32_t swapchainCount, const VkSwapchainKHR* pSwapchains, const VkHdrMetadataEXT* pMetadata) {
	(*fp)(device, swapchainCount, pSwapchains, pMetadata);
}
void domVkSubmitDebugUtilsMessageEXT(PFN_vkSubmitDebugUtilsMessageEXT fp, VkInstance instance, VkDebugUtilsMessageSeverityFlagBitsEXT messageSeverity, VkDebugUtilsMessageTypeFlagsEXT messageTypes, const VkDebugUtilsMessengerCallbackDataEXT* pCallbackData) {
	(*fp)(instance, messageSeverity, messageTypes, pCallbackData);
}
void domVkTrimCommandPoolKHR(PFN_vkTrimCommandPoolKHR fp, VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlags flags) {
	(*fp)(device, commandPool, flags);
}
void domVkUpdateDescriptorSetWithTemplateKHR(PFN_vkUpdateDescriptorSetWithTemplateKHR fp, VkDevice device, VkDescriptorSet descriptorSet, VkDescriptorUpdateTemplate descriptorUpdateTemplate, const void* pData) {
	(*fp)(device, descriptorSet, descriptorUpdateTemplate, pData);
}
VkResult domVkEnumerateInstanceVersion(PFN_vkEnumerateInstanceVersion fp, uint32_t* pApiVersion) {
	return (*fp)(pApiVersion);
}
VkResult domVkBindImageMemory2(PFN_vkBindImageMemory2 fp, VkDevice device, uint32_t bindInfoCount, const VkBindImageMemoryInfo* pBindInfos) {
	return (*fp)(device, bindInfoCount, pBindInfos);
}
void domVkGetDeviceGroupPeerMemoryFeatures(PFN_vkGetDeviceGroupPeerMemoryFeatures fp, VkDevice device, uint32_t heapIndex, uint32_t localDeviceIndex, uint32_t remoteDeviceIndex, VkPeerMemoryFeatureFlags* pPeerMemoryFeatures) {
	(*fp)(device, heapIndex, localDeviceIndex, remoteDeviceIndex, pPeerMemoryFeatures);
}
void domVkCmdDispatchBase(PFN_vkCmdDispatchBase fp, VkCommandBuffer commandBuffer, uint32_t baseGroupX, uint32_t baseGroupY, uint32_t baseGroupZ, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	(*fp)(commandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ);
}
VkResult domVkEnumeratePhysicalDeviceGroups(PFN_vkEnumeratePhysicalDeviceGroups fp, VkInstance instance, uint32_t* pPhysicalDeviceGroupCount, VkPhysicalDeviceGroupProperties* pPhysicalDeviceGroupProperties) {
	return (*fp)(instance, pPhysicalDeviceGroupCount, pPhysicalDeviceGroupProperties);
}
void domVkGetImageMemoryRequirements2(PFN_vkGetImageMemoryRequirements2 fp, VkDevice device, const VkImageMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	(*fp)(device, pInfo, pMemoryRequirements);
}
void domVkGetBufferMemoryRequirements2(PFN_vkGetBufferMemoryRequirements2 fp, VkDevice device, const VkBufferMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	(*fp)(device, pInfo, pMemoryRequirements);
}
void domVkGetImageSparseMemoryRequirements2(PFN_vkGetImageSparseMemoryRequirements2 fp, VkDevice device, const VkImageSparseMemoryRequirementsInfo2* pInfo, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2* pSparseMemoryRequirements) {
	(*fp)(device, pInfo, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}
void domVkGetPhysicalDeviceFeatures2(PFN_vkGetPhysicalDeviceFeatures2 fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures2* pFeatures) {
	(*fp)(physicalDevice, pFeatures);
}
void domVkGetPhysicalDeviceFormatProperties2(PFN_vkGetPhysicalDeviceFormatProperties2 fp, VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties2* pFormatProperties) {
	(*fp)(physicalDevice, format, pFormatProperties);
}
VkResult domVkGetPhysicalDeviceImageFormatProperties2(PFN_vkGetPhysicalDeviceImageFormatProperties2 fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceImageFormatInfo2* pImageFormatInfo, VkImageFormatProperties2* pImageFormatProperties) {
	return (*fp)(physicalDevice, pImageFormatInfo, pImageFormatProperties);
}
void domVkGetPhysicalDeviceQueueFamilyProperties2(PFN_vkGetPhysicalDeviceQueueFamilyProperties2 fp, VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties2* pQueueFamilyProperties) {
	(*fp)(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}
void domVkGetPhysicalDeviceMemoryProperties2(PFN_vkGetPhysicalDeviceMemoryProperties2 fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties2* pMemoryProperties) {
	(*fp)(physicalDevice, pMemoryProperties);
}
void domVkGetPhysicalDeviceSparseImageFormatProperties2(PFN_vkGetPhysicalDeviceSparseImageFormatProperties2 fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceSparseImageFormatInfo2* pFormatInfo, uint32_t* pPropertyCount, VkSparseImageFormatProperties2* pProperties) {
	(*fp)(physicalDevice, pFormatInfo, pPropertyCount, pProperties);
}
VkResult domVkCreateSamplerYcbcrConversion(PFN_vkCreateSamplerYcbcrConversion fp, VkDevice device, const VkSamplerYcbcrConversionCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSamplerYcbcrConversion* pYcbcrConversion) {
	return (*fp)(device, pCreateInfo, pAllocator, pYcbcrConversion);
}
void domVkDestroySamplerYcbcrConversion(PFN_vkDestroySamplerYcbcrConversion fp, VkDevice device, VkSamplerYcbcrConversion ycbcrConversion, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, ycbcrConversion, pAllocator);
}
VkResult domVkCreateDescriptorUpdateTemplate(PFN_vkCreateDescriptorUpdateTemplate fp, VkDevice device, const VkDescriptorUpdateTemplateCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDescriptorUpdateTemplate* pDescriptorUpdateTemplate) {
	return (*fp)(device, pCreateInfo, pAllocator, pDescriptorUpdateTemplate);
}
void domVkDestroyDescriptorUpdateTemplate(PFN_vkDestroyDescriptorUpdateTemplate fp, VkDevice device, VkDescriptorUpdateTemplate descriptorUpdateTemplate, const VkAllocationCallbacks* pAllocator) {
	(*fp)(device, descriptorUpdateTemplate, pAllocator);
}
void domVkUpdateDescriptorSetWithTemplate(PFN_vkUpdateDescriptorSetWithTemplate fp, VkDevice device, VkDescriptorSet descriptorSet, VkDescriptorUpdateTemplate descriptorUpdateTemplate, const void* pData) {
	(*fp)(device, descriptorSet, descriptorUpdateTemplate, pData);
}
void domVkGetPhysicalDeviceExternalBufferProperties(PFN_vkGetPhysicalDeviceExternalBufferProperties fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalBufferInfo* pExternalBufferInfo, VkExternalBufferProperties* pExternalBufferProperties) {
	(*fp)(physicalDevice, pExternalBufferInfo, pExternalBufferProperties);
}
void domVkGetPhysicalDeviceExternalFenceProperties(PFN_vkGetPhysicalDeviceExternalFenceProperties fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalFenceInfo* pExternalFenceInfo, VkExternalFenceProperties* pExternalFenceProperties) {
	(*fp)(physicalDevice, pExternalFenceInfo, pExternalFenceProperties);
}
void domVkGetPhysicalDeviceExternalSemaphoreProperties(PFN_vkGetPhysicalDeviceExternalSemaphoreProperties fp, VkPhysicalDevice physicalDevice, const VkPhysicalDeviceExternalSemaphoreInfo* pExternalSemaphoreInfo, VkExternalSemaphoreProperties* pExternalSemaphoreProperties) {
	(*fp)(physicalDevice, pExternalSemaphoreInfo, pExternalSemaphoreProperties);
}
void domVkGetDescriptorSetLayoutSupport(PFN_vkGetDescriptorSetLayoutSupport fp, VkDevice device, const VkDescriptorSetLayoutCreateInfo* pCreateInfo, VkDescriptorSetLayoutSupport* pSupport) {
	(*fp)(device, pCreateInfo, pSupport);
}

