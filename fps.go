// Copyright (c) 2018 Dominik Honnef

package vk

const (
	vkEnumeratePhysicalDevices = iota
	vkGetPhysicalDeviceProperties
	vkGetPhysicalDeviceFeatures
	vkGetPhysicalDeviceQueueFamilyProperties

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

	// VK_KHR_swapchain
	vkCreateSwapchainKHR
	vkGetSwapchainImagesKHR

	deviceMaxPFN
)

var instanceFpNames = [...]string{
	vkEnumeratePhysicalDevices:               "vkEnumeratePhysicalDevices",
	vkGetPhysicalDeviceProperties:            "vkGetPhysicalDeviceProperties",
	vkGetPhysicalDeviceFeatures:              "vkGetPhysicalDeviceFeatures",
	vkGetPhysicalDeviceQueueFamilyProperties: "vkGetPhysicalDeviceQueueFamilyProperties",

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
	vkGetDeviceQueue:         "vkGetDeviceQueue",
	vkCreateCommandPool:      "vkCreateCommandPool",
	vkTrimCommandPool:        "vkTrimCommandPool",
	vkResetCommandPool:       "vkResetCommandPool",
	vkAllocateCommandBuffers: "vkAllocateCommandBuffers",
	vkResetCommandBuffer:     "vkResetCommandBuffer",
	vkFreeCommandBuffers:     "vkFreeCommandBuffers",
	vkEndCommandBuffer:       "vkEndCommandBuffer",
	vkBeginCommandBuffer:     "vkBeginCommandBuffer",
	vkCmdSetLineWidth:        "vkCmdSetLineWidth",
	vkCmdSetDepthBias:        "vkCmdSetDepthBias",
	vkCmdSetBlendConstants:   "vkCmdSetBlendConstants",
	vkCmdDraw:                "vkCmdDraw",
	vkQueueWaitIdle:          "vkQueueWaitIdle",
	vkDeviceWaitIdle:         "vkDeviceWaitIdle",
	vkCreateSwapchainKHR:     "vkCreateSwapchainKHR",
	vkGetSwapchainImagesKHR:  "vkGetSwapchainImagesKHR",
}
