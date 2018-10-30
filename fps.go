// Copyright (c) 2018 Dominik Honnef

package vk

const (
	vkCreateAndroidSurfaceKHR = iota
	vkCreateDebugReportCallbackEXT
	vkCreateDebugUtilsMessengerEXT
	vkCreateDevice
	vkCreateDisplayModeKHR
	vkCreateDisplayPlaneSurfaceKHR
	vkCreateIOSSurfaceMVK
	vkCreateImagePipeSurfaceFUCHSIA
	vkCreateMacOSSurfaceMVK
	vkCreateMirSurfaceKHR
	vkCreateViSurfaceNN
	vkCreateWaylandSurfaceKHR
	vkCreateWin32SurfaceKHR
	vkCreateXcbSurfaceKHR
	vkCreateXlibSurfaceKHR
	vkDestroyDebugReportCallbackEXT
	vkDestroyDebugUtilsMessengerEXT
	vkDestroyInstance
	vkDestroySurfaceKHR
	vkEnumerateDeviceExtensionProperties
	vkEnumerateDeviceLayerProperties
	vkEnumeratePhysicalDeviceGroups
	vkEnumeratePhysicalDevices
	vkGetDeviceProcAddr
	vkGetPhysicalDeviceCalibrateableTimeDomainsEXT
	vkGetPhysicalDeviceDisplayPlaneProperties2KHR
	vkGetPhysicalDeviceDisplayPlanePropertiesKHR
	vkGetPhysicalDeviceDisplayProperties2KHR
	vkGetPhysicalDeviceDisplayPropertiesKHR
	vkGetPhysicalDeviceExternalBufferProperties
	vkGetPhysicalDeviceExternalFenceProperties
	vkGetPhysicalDeviceExternalImageFormatPropertiesNV
	vkGetPhysicalDeviceExternalSemaphoreProperties
	vkGetPhysicalDeviceFeatures
	vkGetPhysicalDeviceFeatures2
	vkGetPhysicalDeviceFormatProperties
	vkGetPhysicalDeviceFormatProperties2
	vkGetPhysicalDeviceGeneratedCommandsPropertiesNVX
	vkGetPhysicalDeviceImageFormatProperties
	vkGetPhysicalDeviceImageFormatProperties2
	vkGetPhysicalDeviceMemoryProperties
	vkGetPhysicalDeviceMemoryProperties2
	vkGetPhysicalDeviceMirPresentationSupportKHR
	vkGetPhysicalDeviceMultisamplePropertiesEXT
	vkGetPhysicalDevicePresentRectanglesKHR
	vkGetPhysicalDeviceProperties
	vkGetPhysicalDeviceProperties2
	vkGetPhysicalDeviceQueueFamilyProperties
	vkGetPhysicalDeviceQueueFamilyProperties2
	vkGetPhysicalDeviceSparseImageFormatProperties
	vkGetPhysicalDeviceSparseImageFormatProperties2
	vkGetPhysicalDeviceSurfaceCapabilities2EXT
	vkGetPhysicalDeviceSurfaceCapabilities2KHR
	vkGetPhysicalDeviceSurfaceCapabilitiesKHR
	vkGetPhysicalDeviceSurfaceFormats2KHR
	vkGetPhysicalDeviceSurfaceFormatsKHR
	vkGetPhysicalDeviceSurfacePresentModesKHR
	vkGetPhysicalDeviceSurfaceSupportKHR
	vkGetPhysicalDeviceWaylandPresentationSupportKHR
	vkGetPhysicalDeviceWin32PresentationSupportKHR
	vkGetPhysicalDeviceXcbPresentationSupportKHR
	vkGetPhysicalDeviceXlibPresentationSupportKHR

	instanceMaxPFN
)

const (
	vkAcquireImageANDROID = iota
	vkAcquireNextImage2KHR
	vkAcquireNextImageKHR
	vkAcquireXlibDisplayEXT
	vkAllocateCommandBuffers
	vkAllocateDescriptorSets
	vkAllocateMemory
	vkBeginCommandBuffer
	vkBindAccelerationStructureMemoryNVX
	vkBindBufferMemory
	vkBindBufferMemory2
	vkBindImageMemory
	vkBindImageMemory2
	vkCmdBeginConditionalRenderingEXT
	vkCmdBeginDebugUtilsLabelEXT
	vkCmdBeginQuery
	vkCmdBeginQueryIndexedEXT
	vkCmdBeginRenderPass
	vkCmdBeginRenderPass2KHR
	vkCmdBeginTransformFeedbackEXT
	vkCmdBindDescriptorSets
	vkCmdBindIndexBuffer
	vkCmdBindPipeline
	vkCmdBindShadingRateImageNV
	vkCmdBindTransformFeedbackBuffersEXT
	vkCmdBindVertexBuffers
	vkCmdBlitImage
	vkCmdBuildAccelerationStructureNVX
	vkCmdClearAttachments
	vkCmdClearColorImage
	vkCmdClearDepthStencilImage
	vkCmdCopyAccelerationStructureNVX
	vkCmdCopyBuffer
	vkCmdCopyBufferToImage
	vkCmdCopyImage
	vkCmdCopyImageToBuffer
	vkCmdCopyQueryPoolResults
	vkCmdDebugMarkerBeginEXT
	vkCmdDebugMarkerEndEXT
	vkCmdDebugMarkerInsertEXT
	vkCmdDispatch
	vkCmdDispatchBase
	vkCmdDispatchIndirect
	vkCmdDraw
	vkCmdDrawIndexed
	vkCmdDrawIndexedIndirect
	vkCmdDrawIndexedIndirectCountAMD
	vkCmdDrawIndexedIndirectCountKHR
	vkCmdDrawIndirect
	vkCmdDrawIndirectByteCountEXT
	vkCmdDrawIndirectCountAMD
	vkCmdDrawIndirectCountKHR
	vkCmdDrawMeshTasksIndirectCountNV
	vkCmdDrawMeshTasksIndirectNV
	vkCmdDrawMeshTasksNV
	vkCmdEndConditionalRenderingEXT
	vkCmdEndDebugUtilsLabelEXT
	vkCmdEndQuery
	vkCmdEndQueryIndexedEXT
	vkCmdEndRenderPass
	vkCmdEndRenderPass2KHR
	vkCmdEndTransformFeedbackEXT
	vkCmdExecuteCommands
	vkCmdFillBuffer
	vkCmdInsertDebugUtilsLabelEXT
	vkCmdNextSubpass
	vkCmdNextSubpass2KHR
	vkCmdPipelineBarrier
	vkCmdProcessCommandsNVX
	vkCmdPushConstants
	vkCmdPushDescriptorSetKHR
	vkCmdPushDescriptorSetWithTemplateKHR
	vkCmdReserveSpaceForCommandsNVX
	vkCmdResetEvent
	vkCmdResetQueryPool
	vkCmdResolveImage
	vkCmdSetBlendConstants
	vkCmdSetCheckpointNV
	vkCmdSetCoarseSampleOrderNV
	vkCmdSetDepthBias
	vkCmdSetDepthBounds
	vkCmdSetDeviceMask
	vkCmdSetDiscardRectangleEXT
	vkCmdSetEvent
	vkCmdSetExclusiveScissorNV
	vkCmdSetLineWidth
	vkCmdSetSampleLocationsEXT
	vkCmdSetScissor
	vkCmdSetStencilCompareMask
	vkCmdSetStencilReference
	vkCmdSetStencilWriteMask
	vkCmdSetViewport
	vkCmdSetViewportShadingRatePaletteNV
	vkCmdSetViewportWScalingNV
	vkCmdTraceRaysNVX
	vkCmdUpdateBuffer
	vkCmdWaitEvents
	vkCmdWriteAccelerationStructurePropertiesNVX
	vkCmdWriteBufferMarkerAMD
	vkCmdWriteTimestamp
	vkCompileDeferredNVX
	vkCreateAccelerationStructureNVX
	vkCreateBuffer
	vkCreateBufferView
	vkCreateCommandPool
	vkCreateComputePipelines
	vkCreateDescriptorPool
	vkCreateDescriptorSetLayout
	vkCreateDescriptorUpdateTemplate
	vkCreateEvent
	vkCreateFence
	vkCreateFramebuffer
	vkCreateGraphicsPipelines
	vkCreateImage
	vkCreateImageView
	vkCreateIndirectCommandsLayoutNVX
	vkCreateObjectTableNVX
	vkCreatePipelineCache
	vkCreatePipelineLayout
	vkCreateQueryPool
	vkCreateRaytracingPipelinesNVX
	vkCreateRenderPass
	vkCreateRenderPass2KHR
	vkCreateSampler
	vkCreateSamplerYcbcrConversion
	vkCreateSemaphore
	vkCreateShaderModule
	vkCreateSharedSwapchainsKHR
	vkCreateSwapchainKHR
	vkCreateValidationCacheEXT
	vkDebugMarkerSetObjectNameEXT
	vkDebugMarkerSetObjectTagEXT
	vkDebugReportMessageEXT
	vkDestroyAccelerationStructureNVX
	vkDestroyBuffer
	vkDestroyBufferView
	vkDestroyCommandPool
	vkDestroyDescriptorPool
	vkDestroyDescriptorSetLayout
	vkDestroyDescriptorUpdateTemplate
	vkDestroyDevice
	vkDestroyEvent
	vkDestroyFence
	vkDestroyFramebuffer
	vkDestroyImage
	vkDestroyImageView
	vkDestroyIndirectCommandsLayoutNVX
	vkDestroyObjectTableNVX
	vkDestroyPipeline
	vkDestroyPipelineCache
	vkDestroyPipelineLayout
	vkDestroyQueryPool
	vkDestroyRenderPass
	vkDestroySampler
	vkDestroySamplerYcbcrConversion
	vkDestroySemaphore
	vkDestroyShaderModule
	vkDestroySwapchainKHR
	vkDestroyValidationCacheEXT
	vkDeviceWaitIdle
	vkDisplayPowerControlEXT
	vkEndCommandBuffer
	vkFlushMappedMemoryRanges
	vkFreeCommandBuffers
	vkFreeDescriptorSets
	vkFreeMemory
	vkGetAccelerationStructureHandleNVX
	vkGetAccelerationStructureMemoryRequirementsNVX
	vkGetAccelerationStructureScratchMemoryRequirementsNVX
	vkGetAndroidHardwareBufferPropertiesANDROID
	vkGetBufferMemoryRequirements
	vkGetBufferMemoryRequirements2
	vkGetCalibratedTimestampsEXT
	vkGetDescriptorSetLayoutSupport
	vkGetDeviceGroupPeerMemoryFeatures
	vkGetDeviceGroupPresentCapabilitiesKHR
	vkGetDeviceGroupSurfacePresentModesKHR
	vkGetDeviceMemoryCommitment
	vkGetDeviceQueue
	vkGetDeviceQueue2
	vkGetDisplayModeProperties2KHR
	vkGetDisplayModePropertiesKHR
	vkGetDisplayPlaneCapabilities2KHR
	vkGetDisplayPlaneCapabilitiesKHR
	vkGetDisplayPlaneSupportedDisplaysKHR
	vkGetEventStatus
	vkGetFenceFdKHR
	vkGetFenceStatus
	vkGetFenceWin32HandleKHR
	vkGetImageDrmFormatModifierPropertiesEXT
	vkGetImageMemoryRequirements
	vkGetImageMemoryRequirements2
	vkGetImageSparseMemoryRequirements
	vkGetImageSparseMemoryRequirements2
	vkGetImageSubresourceLayout
	vkGetMemoryAndroidHardwareBufferANDROID
	vkGetMemoryFdKHR
	vkGetMemoryFdPropertiesKHR
	vkGetMemoryHostPointerPropertiesEXT
	vkGetMemoryWin32HandleKHR
	vkGetMemoryWin32HandleNV
	vkGetMemoryWin32HandlePropertiesKHR
	vkGetPastPresentationTimingGOOGLE
	vkGetPipelineCacheData
	vkGetQueryPoolResults
	vkGetQueueCheckpointDataNV
	vkGetRandROutputDisplayEXT
	vkGetRaytracingShaderHandlesNVX
	vkGetRefreshCycleDurationGOOGLE
	vkGetRenderAreaGranularity
	vkGetSemaphoreFdKHR
	vkGetSemaphoreWin32HandleKHR
	vkGetShaderInfoAMD
	vkGetSwapchainCounterEXT
	vkGetSwapchainGrallocUsageANDROID
	vkGetSwapchainImagesKHR
	vkGetSwapchainStatusKHR
	vkGetValidationCacheDataEXT
	vkImportFenceFdKHR
	vkImportFenceWin32HandleKHR
	vkImportSemaphoreFdKHR
	vkImportSemaphoreWin32HandleKHR
	vkInvalidateMappedMemoryRanges
	vkMapMemory
	vkMergePipelineCaches
	vkMergeValidationCachesEXT
	vkQueueBeginDebugUtilsLabelEXT
	vkQueueBindSparse
	vkQueueEndDebugUtilsLabelEXT
	vkQueueInsertDebugUtilsLabelEXT
	vkQueuePresentKHR
	vkQueueSignalReleaseImageANDROID
	vkQueueSubmit
	vkQueueWaitIdle
	vkRegisterDeviceEventEXT
	vkRegisterDisplayEventEXT
	vkRegisterObjectsNVX
	vkReleaseDisplayEXT
	vkResetCommandBuffer
	vkResetCommandPool
	vkResetDescriptorPool
	vkResetEvent
	vkResetFences
	vkSetDebugUtilsObjectNameEXT
	vkSetDebugUtilsObjectTagEXT
	vkSetEvent
	vkSetHdrMetadataEXT
	vkSubmitDebugUtilsMessageEXT
	vkTrimCommandPool
	vkUnmapMemory
	vkUnregisterObjectsNVX
	vkUpdateDescriptorSetWithTemplate
	vkUpdateDescriptorSets
	vkWaitForFences

	deviceMaxPFN
)

var instanceFpNames = [...]string{
	vkCreateAndroidSurfaceKHR:                          "vkCreateAndroidSurfaceKHR",
	vkCreateDebugReportCallbackEXT:                     "vkCreateDebugReportCallbackEXT",
	vkCreateDebugUtilsMessengerEXT:                     "vkCreateDebugUtilsMessengerEXT",
	vkCreateDevice:                                     "vkCreateDevice",
	vkCreateDisplayModeKHR:                             "vkCreateDisplayModeKHR",
	vkCreateDisplayPlaneSurfaceKHR:                     "vkCreateDisplayPlaneSurfaceKHR",
	vkCreateIOSSurfaceMVK:                              "vkCreateIOSSurfaceMVK",
	vkCreateImagePipeSurfaceFUCHSIA:                    "vkCreateImagePipeSurfaceFUCHSIA",
	vkCreateMacOSSurfaceMVK:                            "vkCreateMacOSSurfaceMVK",
	vkCreateMirSurfaceKHR:                              "vkCreateMirSurfaceKHR",
	vkCreateViSurfaceNN:                                "vkCreateViSurfaceNN",
	vkCreateWaylandSurfaceKHR:                          "vkCreateWaylandSurfaceKHR",
	vkCreateWin32SurfaceKHR:                            "vkCreateWin32SurfaceKHR",
	vkCreateXcbSurfaceKHR:                              "vkCreateXcbSurfaceKHR",
	vkCreateXlibSurfaceKHR:                             "vkCreateXlibSurfaceKHR",
	vkDestroyDebugReportCallbackEXT:                    "vkDestroyDebugReportCallbackEXT",
	vkDestroyDebugUtilsMessengerEXT:                    "vkDestroyDebugUtilsMessengerEXT",
	vkDestroyInstance:                                  "vkDestroyInstance",
	vkDestroySurfaceKHR:                                "vkDestroySurfaceKHR",
	vkEnumerateDeviceExtensionProperties:               "vkEnumerateDeviceExtensionProperties",
	vkEnumerateDeviceLayerProperties:                   "vkEnumerateDeviceLayerProperties",
	vkEnumeratePhysicalDeviceGroups:                    "vkEnumeratePhysicalDeviceGroups",
	vkEnumeratePhysicalDevices:                         "vkEnumeratePhysicalDevices",
	vkGetDeviceProcAddr:                                "vkGetDeviceProcAddr",
	vkGetPhysicalDeviceCalibrateableTimeDomainsEXT:     "vkGetPhysicalDeviceCalibrateableTimeDomainsEXT",
	vkGetPhysicalDeviceDisplayPlaneProperties2KHR:      "vkGetPhysicalDeviceDisplayPlaneProperties2KHR",
	vkGetPhysicalDeviceDisplayPlanePropertiesKHR:       "vkGetPhysicalDeviceDisplayPlanePropertiesKHR",
	vkGetPhysicalDeviceDisplayProperties2KHR:           "vkGetPhysicalDeviceDisplayProperties2KHR",
	vkGetPhysicalDeviceDisplayPropertiesKHR:            "vkGetPhysicalDeviceDisplayPropertiesKHR",
	vkGetPhysicalDeviceExternalBufferProperties:        "vkGetPhysicalDeviceExternalBufferProperties",
	vkGetPhysicalDeviceExternalFenceProperties:         "vkGetPhysicalDeviceExternalFenceProperties",
	vkGetPhysicalDeviceExternalImageFormatPropertiesNV: "vkGetPhysicalDeviceExternalImageFormatPropertiesNV",
	vkGetPhysicalDeviceExternalSemaphoreProperties:     "vkGetPhysicalDeviceExternalSemaphoreProperties",
	vkGetPhysicalDeviceFeatures:                        "vkGetPhysicalDeviceFeatures",
	vkGetPhysicalDeviceFeatures2:                       "vkGetPhysicalDeviceFeatures2",
	vkGetPhysicalDeviceFormatProperties:                "vkGetPhysicalDeviceFormatProperties",
	vkGetPhysicalDeviceFormatProperties2:               "vkGetPhysicalDeviceFormatProperties2",
	vkGetPhysicalDeviceGeneratedCommandsPropertiesNVX:  "vkGetPhysicalDeviceGeneratedCommandsPropertiesNVX",
	vkGetPhysicalDeviceImageFormatProperties:           "vkGetPhysicalDeviceImageFormatProperties",
	vkGetPhysicalDeviceImageFormatProperties2:          "vkGetPhysicalDeviceImageFormatProperties2",
	vkGetPhysicalDeviceMemoryProperties:                "vkGetPhysicalDeviceMemoryProperties",
	vkGetPhysicalDeviceMemoryProperties2:               "vkGetPhysicalDeviceMemoryProperties2",
	vkGetPhysicalDeviceMirPresentationSupportKHR:       "vkGetPhysicalDeviceMirPresentationSupportKHR",
	vkGetPhysicalDeviceMultisamplePropertiesEXT:        "vkGetPhysicalDeviceMultisamplePropertiesEXT",
	vkGetPhysicalDevicePresentRectanglesKHR:            "vkGetPhysicalDevicePresentRectanglesKHR",
	vkGetPhysicalDeviceProperties:                      "vkGetPhysicalDeviceProperties",
	vkGetPhysicalDeviceProperties2:                     "vkGetPhysicalDeviceProperties2",
	vkGetPhysicalDeviceQueueFamilyProperties:           "vkGetPhysicalDeviceQueueFamilyProperties",
	vkGetPhysicalDeviceQueueFamilyProperties2:          "vkGetPhysicalDeviceQueueFamilyProperties2",
	vkGetPhysicalDeviceSparseImageFormatProperties:     "vkGetPhysicalDeviceSparseImageFormatProperties",
	vkGetPhysicalDeviceSparseImageFormatProperties2:    "vkGetPhysicalDeviceSparseImageFormatProperties2",
	vkGetPhysicalDeviceSurfaceCapabilities2EXT:         "vkGetPhysicalDeviceSurfaceCapabilities2EXT",
	vkGetPhysicalDeviceSurfaceCapabilities2KHR:         "vkGetPhysicalDeviceSurfaceCapabilities2KHR",
	vkGetPhysicalDeviceSurfaceCapabilitiesKHR:          "vkGetPhysicalDeviceSurfaceCapabilitiesKHR",
	vkGetPhysicalDeviceSurfaceFormats2KHR:              "vkGetPhysicalDeviceSurfaceFormats2KHR",
	vkGetPhysicalDeviceSurfaceFormatsKHR:               "vkGetPhysicalDeviceSurfaceFormatsKHR",
	vkGetPhysicalDeviceSurfacePresentModesKHR:          "vkGetPhysicalDeviceSurfacePresentModesKHR",
	vkGetPhysicalDeviceSurfaceSupportKHR:               "vkGetPhysicalDeviceSurfaceSupportKHR",
	vkGetPhysicalDeviceWaylandPresentationSupportKHR:   "vkGetPhysicalDeviceWaylandPresentationSupportKHR",
	vkGetPhysicalDeviceWin32PresentationSupportKHR:     "vkGetPhysicalDeviceWin32PresentationSupportKHR",
	vkGetPhysicalDeviceXcbPresentationSupportKHR:       "vkGetPhysicalDeviceXcbPresentationSupportKHR",
	vkGetPhysicalDeviceXlibPresentationSupportKHR:      "vkGetPhysicalDeviceXlibPresentationSupportKHR",
}

var deviceFpNames = [...]string{
	vkAcquireImageANDROID:                                  "vkAcquireImageANDROID",
	vkAcquireNextImage2KHR:                                 "vkAcquireNextImage2KHR",
	vkAcquireNextImageKHR:                                  "vkAcquireNextImageKHR",
	vkAcquireXlibDisplayEXT:                                "vkAcquireXlibDisplayEXT",
	vkAllocateCommandBuffers:                               "vkAllocateCommandBuffers",
	vkAllocateDescriptorSets:                               "vkAllocateDescriptorSets",
	vkAllocateMemory:                                       "vkAllocateMemory",
	vkBeginCommandBuffer:                                   "vkBeginCommandBuffer",
	vkBindAccelerationStructureMemoryNVX:                   "vkBindAccelerationStructureMemoryNVX",
	vkBindBufferMemory:                                     "vkBindBufferMemory",
	vkBindBufferMemory2:                                    "vkBindBufferMemory2",
	vkBindImageMemory:                                      "vkBindImageMemory",
	vkBindImageMemory2:                                     "vkBindImageMemory2",
	vkCmdBeginConditionalRenderingEXT:                      "vkCmdBeginConditionalRenderingEXT",
	vkCmdBeginDebugUtilsLabelEXT:                           "vkCmdBeginDebugUtilsLabelEXT",
	vkCmdBeginQuery:                                        "vkCmdBeginQuery",
	vkCmdBeginQueryIndexedEXT:                              "vkCmdBeginQueryIndexedEXT",
	vkCmdBeginRenderPass:                                   "vkCmdBeginRenderPass",
	vkCmdBeginRenderPass2KHR:                               "vkCmdBeginRenderPass2KHR",
	vkCmdBeginTransformFeedbackEXT:                         "vkCmdBeginTransformFeedbackEXT",
	vkCmdBindDescriptorSets:                                "vkCmdBindDescriptorSets",
	vkCmdBindIndexBuffer:                                   "vkCmdBindIndexBuffer",
	vkCmdBindPipeline:                                      "vkCmdBindPipeline",
	vkCmdBindShadingRateImageNV:                            "vkCmdBindShadingRateImageNV",
	vkCmdBindTransformFeedbackBuffersEXT:                   "vkCmdBindTransformFeedbackBuffersEXT",
	vkCmdBindVertexBuffers:                                 "vkCmdBindVertexBuffers",
	vkCmdBlitImage:                                         "vkCmdBlitImage",
	vkCmdBuildAccelerationStructureNVX:                     "vkCmdBuildAccelerationStructureNVX",
	vkCmdClearAttachments:                                  "vkCmdClearAttachments",
	vkCmdClearColorImage:                                   "vkCmdClearColorImage",
	vkCmdClearDepthStencilImage:                            "vkCmdClearDepthStencilImage",
	vkCmdCopyAccelerationStructureNVX:                      "vkCmdCopyAccelerationStructureNVX",
	vkCmdCopyBuffer:                                        "vkCmdCopyBuffer",
	vkCmdCopyBufferToImage:                                 "vkCmdCopyBufferToImage",
	vkCmdCopyImage:                                         "vkCmdCopyImage",
	vkCmdCopyImageToBuffer:                                 "vkCmdCopyImageToBuffer",
	vkCmdCopyQueryPoolResults:                              "vkCmdCopyQueryPoolResults",
	vkCmdDebugMarkerBeginEXT:                               "vkCmdDebugMarkerBeginEXT",
	vkCmdDebugMarkerEndEXT:                                 "vkCmdDebugMarkerEndEXT",
	vkCmdDebugMarkerInsertEXT:                              "vkCmdDebugMarkerInsertEXT",
	vkCmdDispatch:                                          "vkCmdDispatch",
	vkCmdDispatchBase:                                      "vkCmdDispatchBase",
	vkCmdDispatchIndirect:                                  "vkCmdDispatchIndirect",
	vkCmdDraw:                                              "vkCmdDraw",
	vkCmdDrawIndexed:                                       "vkCmdDrawIndexed",
	vkCmdDrawIndexedIndirect:                               "vkCmdDrawIndexedIndirect",
	vkCmdDrawIndexedIndirectCountAMD:                       "vkCmdDrawIndexedIndirectCountAMD",
	vkCmdDrawIndexedIndirectCountKHR:                       "vkCmdDrawIndexedIndirectCountKHR",
	vkCmdDrawIndirect:                                      "vkCmdDrawIndirect",
	vkCmdDrawIndirectByteCountEXT:                          "vkCmdDrawIndirectByteCountEXT",
	vkCmdDrawIndirectCountAMD:                              "vkCmdDrawIndirectCountAMD",
	vkCmdDrawIndirectCountKHR:                              "vkCmdDrawIndirectCountKHR",
	vkCmdDrawMeshTasksIndirectCountNV:                      "vkCmdDrawMeshTasksIndirectCountNV",
	vkCmdDrawMeshTasksIndirectNV:                           "vkCmdDrawMeshTasksIndirectNV",
	vkCmdDrawMeshTasksNV:                                   "vkCmdDrawMeshTasksNV",
	vkCmdEndConditionalRenderingEXT:                        "vkCmdEndConditionalRenderingEXT",
	vkCmdEndDebugUtilsLabelEXT:                             "vkCmdEndDebugUtilsLabelEXT",
	vkCmdEndQuery:                                          "vkCmdEndQuery",
	vkCmdEndQueryIndexedEXT:                                "vkCmdEndQueryIndexedEXT",
	vkCmdEndRenderPass:                                     "vkCmdEndRenderPass",
	vkCmdEndRenderPass2KHR:                                 "vkCmdEndRenderPass2KHR",
	vkCmdEndTransformFeedbackEXT:                           "vkCmdEndTransformFeedbackEXT",
	vkCmdExecuteCommands:                                   "vkCmdExecuteCommands",
	vkCmdFillBuffer:                                        "vkCmdFillBuffer",
	vkCmdInsertDebugUtilsLabelEXT:                          "vkCmdInsertDebugUtilsLabelEXT",
	vkCmdNextSubpass:                                       "vkCmdNextSubpass",
	vkCmdNextSubpass2KHR:                                   "vkCmdNextSubpass2KHR",
	vkCmdPipelineBarrier:                                   "vkCmdPipelineBarrier",
	vkCmdProcessCommandsNVX:                                "vkCmdProcessCommandsNVX",
	vkCmdPushConstants:                                     "vkCmdPushConstants",
	vkCmdPushDescriptorSetKHR:                              "vkCmdPushDescriptorSetKHR",
	vkCmdPushDescriptorSetWithTemplateKHR:                  "vkCmdPushDescriptorSetWithTemplateKHR",
	vkCmdReserveSpaceForCommandsNVX:                        "vkCmdReserveSpaceForCommandsNVX",
	vkCmdResetEvent:                                        "vkCmdResetEvent",
	vkCmdResetQueryPool:                                    "vkCmdResetQueryPool",
	vkCmdResolveImage:                                      "vkCmdResolveImage",
	vkCmdSetBlendConstants:                                 "vkCmdSetBlendConstants",
	vkCmdSetCheckpointNV:                                   "vkCmdSetCheckpointNV",
	vkCmdSetCoarseSampleOrderNV:                            "vkCmdSetCoarseSampleOrderNV",
	vkCmdSetDepthBias:                                      "vkCmdSetDepthBias",
	vkCmdSetDepthBounds:                                    "vkCmdSetDepthBounds",
	vkCmdSetDeviceMask:                                     "vkCmdSetDeviceMask",
	vkCmdSetDiscardRectangleEXT:                            "vkCmdSetDiscardRectangleEXT",
	vkCmdSetEvent:                                          "vkCmdSetEvent",
	vkCmdSetExclusiveScissorNV:                             "vkCmdSetExclusiveScissorNV",
	vkCmdSetLineWidth:                                      "vkCmdSetLineWidth",
	vkCmdSetSampleLocationsEXT:                             "vkCmdSetSampleLocationsEXT",
	vkCmdSetScissor:                                        "vkCmdSetScissor",
	vkCmdSetStencilCompareMask:                             "vkCmdSetStencilCompareMask",
	vkCmdSetStencilReference:                               "vkCmdSetStencilReference",
	vkCmdSetStencilWriteMask:                               "vkCmdSetStencilWriteMask",
	vkCmdSetViewport:                                       "vkCmdSetViewport",
	vkCmdSetViewportShadingRatePaletteNV:                   "vkCmdSetViewportShadingRatePaletteNV",
	vkCmdSetViewportWScalingNV:                             "vkCmdSetViewportWScalingNV",
	vkCmdTraceRaysNVX:                                      "vkCmdTraceRaysNVX",
	vkCmdUpdateBuffer:                                      "vkCmdUpdateBuffer",
	vkCmdWaitEvents:                                        "vkCmdWaitEvents",
	vkCmdWriteAccelerationStructurePropertiesNVX:           "vkCmdWriteAccelerationStructurePropertiesNVX",
	vkCmdWriteBufferMarkerAMD:                              "vkCmdWriteBufferMarkerAMD",
	vkCmdWriteTimestamp:                                    "vkCmdWriteTimestamp",
	vkCompileDeferredNVX:                                   "vkCompileDeferredNVX",
	vkCreateAccelerationStructureNVX:                       "vkCreateAccelerationStructureNVX",
	vkCreateBuffer:                                         "vkCreateBuffer",
	vkCreateBufferView:                                     "vkCreateBufferView",
	vkCreateCommandPool:                                    "vkCreateCommandPool",
	vkCreateComputePipelines:                               "vkCreateComputePipelines",
	vkCreateDescriptorPool:                                 "vkCreateDescriptorPool",
	vkCreateDescriptorSetLayout:                            "vkCreateDescriptorSetLayout",
	vkCreateDescriptorUpdateTemplate:                       "vkCreateDescriptorUpdateTemplate",
	vkCreateEvent:                                          "vkCreateEvent",
	vkCreateFence:                                          "vkCreateFence",
	vkCreateFramebuffer:                                    "vkCreateFramebuffer",
	vkCreateGraphicsPipelines:                              "vkCreateGraphicsPipelines",
	vkCreateImage:                                          "vkCreateImage",
	vkCreateImageView:                                      "vkCreateImageView",
	vkCreateIndirectCommandsLayoutNVX:                      "vkCreateIndirectCommandsLayoutNVX",
	vkCreateObjectTableNVX:                                 "vkCreateObjectTableNVX",
	vkCreatePipelineCache:                                  "vkCreatePipelineCache",
	vkCreatePipelineLayout:                                 "vkCreatePipelineLayout",
	vkCreateQueryPool:                                      "vkCreateQueryPool",
	vkCreateRaytracingPipelinesNVX:                         "vkCreateRaytracingPipelinesNVX",
	vkCreateRenderPass:                                     "vkCreateRenderPass",
	vkCreateRenderPass2KHR:                                 "vkCreateRenderPass2KHR",
	vkCreateSampler:                                        "vkCreateSampler",
	vkCreateSamplerYcbcrConversion:                         "vkCreateSamplerYcbcrConversion",
	vkCreateSemaphore:                                      "vkCreateSemaphore",
	vkCreateShaderModule:                                   "vkCreateShaderModule",
	vkCreateSharedSwapchainsKHR:                            "vkCreateSharedSwapchainsKHR",
	vkCreateSwapchainKHR:                                   "vkCreateSwapchainKHR",
	vkCreateValidationCacheEXT:                             "vkCreateValidationCacheEXT",
	vkDebugMarkerSetObjectNameEXT:                          "vkDebugMarkerSetObjectNameEXT",
	vkDebugMarkerSetObjectTagEXT:                           "vkDebugMarkerSetObjectTagEXT",
	vkDebugReportMessageEXT:                                "vkDebugReportMessageEXT",
	vkDestroyAccelerationStructureNVX:                      "vkDestroyAccelerationStructureNVX",
	vkDestroyBuffer:                                        "vkDestroyBuffer",
	vkDestroyBufferView:                                    "vkDestroyBufferView",
	vkDestroyCommandPool:                                   "vkDestroyCommandPool",
	vkDestroyDescriptorPool:                                "vkDestroyDescriptorPool",
	vkDestroyDescriptorSetLayout:                           "vkDestroyDescriptorSetLayout",
	vkDestroyDescriptorUpdateTemplate:                      "vkDestroyDescriptorUpdateTemplate",
	vkDestroyDevice:                                        "vkDestroyDevice",
	vkDestroyEvent:                                         "vkDestroyEvent",
	vkDestroyFence:                                         "vkDestroyFence",
	vkDestroyFramebuffer:                                   "vkDestroyFramebuffer",
	vkDestroyImage:                                         "vkDestroyImage",
	vkDestroyImageView:                                     "vkDestroyImageView",
	vkDestroyIndirectCommandsLayoutNVX:                     "vkDestroyIndirectCommandsLayoutNVX",
	vkDestroyObjectTableNVX:                                "vkDestroyObjectTableNVX",
	vkDestroyPipeline:                                      "vkDestroyPipeline",
	vkDestroyPipelineCache:                                 "vkDestroyPipelineCache",
	vkDestroyPipelineLayout:                                "vkDestroyPipelineLayout",
	vkDestroyQueryPool:                                     "vkDestroyQueryPool",
	vkDestroyRenderPass:                                    "vkDestroyRenderPass",
	vkDestroySampler:                                       "vkDestroySampler",
	vkDestroySamplerYcbcrConversion:                        "vkDestroySamplerYcbcrConversion",
	vkDestroySemaphore:                                     "vkDestroySemaphore",
	vkDestroyShaderModule:                                  "vkDestroyShaderModule",
	vkDestroySwapchainKHR:                                  "vkDestroySwapchainKHR",
	vkDestroyValidationCacheEXT:                            "vkDestroyValidationCacheEXT",
	vkDeviceWaitIdle:                                       "vkDeviceWaitIdle",
	vkDisplayPowerControlEXT:                               "vkDisplayPowerControlEXT",
	vkEndCommandBuffer:                                     "vkEndCommandBuffer",
	vkFlushMappedMemoryRanges:                              "vkFlushMappedMemoryRanges",
	vkFreeCommandBuffers:                                   "vkFreeCommandBuffers",
	vkFreeDescriptorSets:                                   "vkFreeDescriptorSets",
	vkFreeMemory:                                           "vkFreeMemory",
	vkGetAccelerationStructureHandleNVX:                    "vkGetAccelerationStructureHandleNVX",
	vkGetAccelerationStructureMemoryRequirementsNVX:        "vkGetAccelerationStructureMemoryRequirementsNVX",
	vkGetAccelerationStructureScratchMemoryRequirementsNVX: "vkGetAccelerationStructureScratchMemoryRequirementsNVX",
	vkGetAndroidHardwareBufferPropertiesANDROID:            "vkGetAndroidHardwareBufferPropertiesANDROID",
	vkGetBufferMemoryRequirements:                          "vkGetBufferMemoryRequirements",
	vkGetBufferMemoryRequirements2:                         "vkGetBufferMemoryRequirements2",
	vkGetCalibratedTimestampsEXT:                           "vkGetCalibratedTimestampsEXT",
	vkGetDescriptorSetLayoutSupport:                        "vkGetDescriptorSetLayoutSupport",
	vkGetDeviceGroupPeerMemoryFeatures:                     "vkGetDeviceGroupPeerMemoryFeatures",
	vkGetDeviceGroupPresentCapabilitiesKHR:                 "vkGetDeviceGroupPresentCapabilitiesKHR",
	vkGetDeviceGroupSurfacePresentModesKHR:                 "vkGetDeviceGroupSurfacePresentModesKHR",
	vkGetDeviceMemoryCommitment:                            "vkGetDeviceMemoryCommitment",
	vkGetDeviceQueue:                                       "vkGetDeviceQueue",
	vkGetDeviceQueue2:                                      "vkGetDeviceQueue2",
	vkGetDisplayModeProperties2KHR:                         "vkGetDisplayModeProperties2KHR",
	vkGetDisplayModePropertiesKHR:                          "vkGetDisplayModePropertiesKHR",
	vkGetDisplayPlaneCapabilities2KHR:                      "vkGetDisplayPlaneCapabilities2KHR",
	vkGetDisplayPlaneCapabilitiesKHR:                       "vkGetDisplayPlaneCapabilitiesKHR",
	vkGetDisplayPlaneSupportedDisplaysKHR:                  "vkGetDisplayPlaneSupportedDisplaysKHR",
	vkGetEventStatus:                                       "vkGetEventStatus",
	vkGetFenceFdKHR:                                        "vkGetFenceFdKHR",
	vkGetFenceStatus:                                       "vkGetFenceStatus",
	vkGetFenceWin32HandleKHR:                               "vkGetFenceWin32HandleKHR",
	vkGetImageDrmFormatModifierPropertiesEXT:               "vkGetImageDrmFormatModifierPropertiesEXT",
	vkGetImageMemoryRequirements:                           "vkGetImageMemoryRequirements",
	vkGetImageMemoryRequirements2:                          "vkGetImageMemoryRequirements2",
	vkGetImageSparseMemoryRequirements:                     "vkGetImageSparseMemoryRequirements",
	vkGetImageSparseMemoryRequirements2:                    "vkGetImageSparseMemoryRequirements2",
	vkGetImageSubresourceLayout:                            "vkGetImageSubresourceLayout",
	vkGetMemoryAndroidHardwareBufferANDROID:                "vkGetMemoryAndroidHardwareBufferANDROID",
	vkGetMemoryFdKHR:                                       "vkGetMemoryFdKHR",
	vkGetMemoryFdPropertiesKHR:                             "vkGetMemoryFdPropertiesKHR",
	vkGetMemoryHostPointerPropertiesEXT:                    "vkGetMemoryHostPointerPropertiesEXT",
	vkGetMemoryWin32HandleKHR:                              "vkGetMemoryWin32HandleKHR",
	vkGetMemoryWin32HandleNV:                               "vkGetMemoryWin32HandleNV",
	vkGetMemoryWin32HandlePropertiesKHR:                    "vkGetMemoryWin32HandlePropertiesKHR",
	vkGetPastPresentationTimingGOOGLE:                      "vkGetPastPresentationTimingGOOGLE",
	vkGetPipelineCacheData:                                 "vkGetPipelineCacheData",
	vkGetQueryPoolResults:                                  "vkGetQueryPoolResults",
	vkGetQueueCheckpointDataNV:                             "vkGetQueueCheckpointDataNV",
	vkGetRandROutputDisplayEXT:                             "vkGetRandROutputDisplayEXT",
	vkGetRaytracingShaderHandlesNVX:                        "vkGetRaytracingShaderHandlesNVX",
	vkGetRefreshCycleDurationGOOGLE:                        "vkGetRefreshCycleDurationGOOGLE",
	vkGetRenderAreaGranularity:                             "vkGetRenderAreaGranularity",
	vkGetSemaphoreFdKHR:                                    "vkGetSemaphoreFdKHR",
	vkGetSemaphoreWin32HandleKHR:                           "vkGetSemaphoreWin32HandleKHR",
	vkGetShaderInfoAMD:                                     "vkGetShaderInfoAMD",
	vkGetSwapchainCounterEXT:                               "vkGetSwapchainCounterEXT",
	vkGetSwapchainGrallocUsageANDROID:                      "vkGetSwapchainGrallocUsageANDROID",
	vkGetSwapchainImagesKHR:                                "vkGetSwapchainImagesKHR",
	vkGetSwapchainStatusKHR:                                "vkGetSwapchainStatusKHR",
	vkGetValidationCacheDataEXT:                            "vkGetValidationCacheDataEXT",
	vkImportFenceFdKHR:                                     "vkImportFenceFdKHR",
	vkImportFenceWin32HandleKHR:                            "vkImportFenceWin32HandleKHR",
	vkImportSemaphoreFdKHR:                                 "vkImportSemaphoreFdKHR",
	vkImportSemaphoreWin32HandleKHR:                        "vkImportSemaphoreWin32HandleKHR",
	vkInvalidateMappedMemoryRanges:                         "vkInvalidateMappedMemoryRanges",
	vkMapMemory:                                            "vkMapMemory",
	vkMergePipelineCaches:                                  "vkMergePipelineCaches",
	vkMergeValidationCachesEXT:                             "vkMergeValidationCachesEXT",
	vkQueueBeginDebugUtilsLabelEXT:                         "vkQueueBeginDebugUtilsLabelEXT",
	vkQueueBindSparse:                                      "vkQueueBindSparse",
	vkQueueEndDebugUtilsLabelEXT:                           "vkQueueEndDebugUtilsLabelEXT",
	vkQueueInsertDebugUtilsLabelEXT:                        "vkQueueInsertDebugUtilsLabelEXT",
	vkQueuePresentKHR:                                      "vkQueuePresentKHR",
	vkQueueSignalReleaseImageANDROID:                       "vkQueueSignalReleaseImageANDROID",
	vkQueueSubmit:                                          "vkQueueSubmit",
	vkQueueWaitIdle:                                        "vkQueueWaitIdle",
	vkRegisterDeviceEventEXT:                               "vkRegisterDeviceEventEXT",
	vkRegisterDisplayEventEXT:                              "vkRegisterDisplayEventEXT",
	vkRegisterObjectsNVX:                                   "vkRegisterObjectsNVX",
	vkReleaseDisplayEXT:                                    "vkReleaseDisplayEXT",
	vkResetCommandBuffer:                                   "vkResetCommandBuffer",
	vkResetCommandPool:                                     "vkResetCommandPool",
	vkResetDescriptorPool:                                  "vkResetDescriptorPool",
	vkResetEvent:                                           "vkResetEvent",
	vkResetFences:                                          "vkResetFences",
	vkSetDebugUtilsObjectNameEXT:                           "vkSetDebugUtilsObjectNameEXT",
	vkSetDebugUtilsObjectTagEXT:                            "vkSetDebugUtilsObjectTagEXT",
	vkSetEvent:                                             "vkSetEvent",
	vkSetHdrMetadataEXT:                                    "vkSetHdrMetadataEXT",
	vkSubmitDebugUtilsMessageEXT:                           "vkSubmitDebugUtilsMessageEXT",
	vkTrimCommandPool:                                      "vkTrimCommandPool",
	vkUnmapMemory:                                          "vkUnmapMemory",
	vkUnregisterObjectsNVX:                                 "vkUnregisterObjectsNVX",
	vkUpdateDescriptorSetWithTemplate:                      "vkUpdateDescriptorSetWithTemplate",
	vkUpdateDescriptorSets:                                 "vkUpdateDescriptorSets",
	vkWaitForFences:                                        "vkWaitForFences",
}
