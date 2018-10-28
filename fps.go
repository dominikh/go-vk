// Copyright (c) 2018 Dominik Honnef

package vk

const (
	vkEnumeratePhysicalDevices = iota
	vkGetPhysicalDeviceProperties
	vkGetPhysicalDeviceFeatures
	vkGetPhysicalDeviceQueueFamilyProperties
	vkEnumerateDeviceExtensionProperties
	vkGetPhysicalDeviceMemoryProperties
	vkGetPhysicalDeviceProperties2

	// VK_KHR_surface
	vkGetPhysicalDeviceSurfaceSupportKHR
	vkGetPhysicalDeviceSurfaceCapabilitiesKHR
	vkGetPhysicalDeviceSurfaceFormatsKHR
	vkGetPhysicalDeviceSurfacePresentModesKHR
	vkDestroySurfaceKHR
	// WSI
	vkCreateAndroidSurfaceKHR
	vkCreateMirSurfaceKHR
	vkCreateWaylandSurfaceKHR
	vkCreateWin32SurfaceKHR
	vkCreateXlibSurfaceKHR
	vkGetPhysicalDeviceXlibPresentationSupportKHR

	vkCreateDevice
	vkGetDeviceProcAddr

	instanceMaxPFN
)

const (
	vkGetDeviceQueue = iota
	vkCreateCommandPool
	vkTrimCommandPool
	vkResetCommandPool
	vkAllocateCommandBuffers
	vkResetCommandBuffer
	vkFreeCommandBuffers
	vkEndCommandBuffer
	vkBeginCommandBuffer
	vkCmdSetLineWidth
	vkCmdSetDepthBias
	vkCmdSetBlendConstants
	vkCmdDraw
	vkQueueWaitIdle
	vkDeviceWaitIdle
	vkCreateImageView
	vkCreateShaderModule
	vkCreateGraphicsPipelines
	vkCreatePipelineLayout
	vkCreateRenderPass
	vkCreateFramebuffer
	vkCmdBeginRenderPass
	vkCmdBindPipeline
	vkCmdEndRenderPass
	vkCreateSemaphore
	vkQueueSubmit
	vkCreateFence
	vkWaitForFences
	vkResetFences
	vkCreateBuffer
	vkGetBufferMemoryRequirements
	vkAllocateMemory
	vkBindBufferMemory
	vkBindBufferMemory2
	vkMapMemory
	vkUnmapMemory
	vkFreeMemory

	// VK_KHR_swapchain
	vkCreateSwapchainKHR
	vkGetSwapchainImagesKHR
	vkAcquireNextImageKHR
	vkQueuePresentKHR

	deviceMaxPFN
)

var instanceFpNames = [...]string{
	vkEnumeratePhysicalDevices:               "vkEnumeratePhysicalDevices",
	vkGetPhysicalDeviceProperties:            "vkGetPhysicalDeviceProperties",
	vkGetPhysicalDeviceFeatures:              "vkGetPhysicalDeviceFeatures",
	vkGetPhysicalDeviceQueueFamilyProperties: "vkGetPhysicalDeviceQueueFamilyProperties",
	vkEnumerateDeviceExtensionProperties:     "vkEnumerateDeviceExtensionProperties",
	vkGetPhysicalDeviceMemoryProperties:      "vkGetPhysicalDeviceMemoryProperties",
	vkGetPhysicalDeviceProperties2:           "vkGetPhysicalDeviceProperties2",

	vkGetPhysicalDeviceSurfaceSupportKHR:          "vkGetPhysicalDeviceSurfaceSupportKHR",
	vkGetPhysicalDeviceSurfaceCapabilitiesKHR:     "vkGetPhysicalDeviceSurfaceCapabilitiesKHR",
	vkGetPhysicalDeviceSurfaceFormatsKHR:          "vkGetPhysicalDeviceSurfaceFormatsKHR",
	vkGetPhysicalDeviceSurfacePresentModesKHR:     "vkGetPhysicalDeviceSurfacePresentModesKHR",
	vkDestroySurfaceKHR:                           "vkDestroySurfaceKHR",
	vkCreateAndroidSurfaceKHR:                     "vkCreateAndroidSurfaceKHR",
	vkCreateMirSurfaceKHR:                         "vkCreateMirSurfaceKHR",
	vkCreateWaylandSurfaceKHR:                     "vkCreateWaylandSurfaceKHR",
	vkCreateWin32SurfaceKHR:                       "vkCreateWin32SurfaceKHR",
	vkCreateXlibSurfaceKHR:                        "vkCreateXlibSurfaceKHR",
	vkGetPhysicalDeviceXlibPresentationSupportKHR: "vkGetPhysicalDeviceXlibPresentationSupportKHR",

	vkCreateDevice:      "vkCreateDevice",
	vkGetDeviceProcAddr: "vkGetDeviceProcAddr",
}

var deviceFpNames = [...]string{
	vkGetDeviceQueue:              "vkGetDeviceQueue",
	vkCreateCommandPool:           "vkCreateCommandPool",
	vkTrimCommandPool:             "vkTrimCommandPool",
	vkResetCommandPool:            "vkResetCommandPool",
	vkAllocateCommandBuffers:      "vkAllocateCommandBuffers",
	vkResetCommandBuffer:          "vkResetCommandBuffer",
	vkFreeCommandBuffers:          "vkFreeCommandBuffers",
	vkEndCommandBuffer:            "vkEndCommandBuffer",
	vkBeginCommandBuffer:          "vkBeginCommandBuffer",
	vkCmdSetLineWidth:             "vkCmdSetLineWidth",
	vkCmdSetDepthBias:             "vkCmdSetDepthBias",
	vkCmdSetBlendConstants:        "vkCmdSetBlendConstants",
	vkCmdDraw:                     "vkCmdDraw",
	vkQueueWaitIdle:               "vkQueueWaitIdle",
	vkDeviceWaitIdle:              "vkDeviceWaitIdle",
	vkCreateSwapchainKHR:          "vkCreateSwapchainKHR",
	vkGetSwapchainImagesKHR:       "vkGetSwapchainImagesKHR",
	vkCreateImageView:             "vkCreateImageView",
	vkCreateShaderModule:          "vkCreateShaderModule",
	vkCreateGraphicsPipelines:     "vkCreateGraphicsPipelines",
	vkCreatePipelineLayout:        "vkCreatePipelineLayout",
	vkCreateRenderPass:            "vkCreateRenderPass",
	vkCreateFramebuffer:           "vkCreateFramebuffer",
	vkCmdBeginRenderPass:          "vkCmdBeginRenderPass",
	vkCmdBindPipeline:             "vkCmdBindPipeline",
	vkCmdEndRenderPass:            "vkCmdEndRenderPass",
	vkCreateSemaphore:             "vkCreateSemaphore",
	vkAcquireNextImageKHR:         "vkAcquireNextImageKHR",
	vkQueueSubmit:                 "vkQueueSubmit",
	vkQueuePresentKHR:             "vkQueuePresentKHR",
	vkCreateFence:                 "vkCreateFence",
	vkWaitForFences:               "vkWaitForFences",
	vkResetFences:                 "vkResetFences",
	vkCreateBuffer:                "vkCreateBuffer",
	vkGetBufferMemoryRequirements: "vkGetBufferMemoryRequirements",
	vkAllocateMemory:              "vkAllocateMemory",
	vkBindBufferMemory:            "vkBindBufferMemory",
	vkBindBufferMemory2:           "vkBindBufferMemory2",
	vkMapMemory:                   "vkMapMemory",
	vkUnmapMemory:                 "vkUnmapMemory",
	vkFreeMemory:                  "vkFreeMemory",
}
