// Copyright (c) 2015-2018 Khronos Group
// Copyright (c) 2018 Dominik Honnef

package vk

// #cgo pkg-config: vulkan
// #cgo CFLAGS: -DVK_NO_PROTOTYPES
// #include <vulkan/vulkan_core.h>
// #include <stdlib.h>
//
// VkResult domVkCreateSwapchainKHR(PFN_vkCreateSwapchainKHR fp, VkDevice device, const VkSwapchainCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSwapchainKHR* pSwapchain);
// VkResult domVkGetSwapchainImagesKHR(PFN_vkGetSwapchainImagesKHR fp, VkDevice device, VkSwapchainKHR swapchain, uint32_t* pSwapchainImageCount, VkImage* pSwapchainImages);
import "C"

import (
	"unsafe"
)

type SwapchainCreateInfoKHR struct {
	Next               unsafe.Pointer
	Surface            SurfaceKHR
	MinImageCount      uint32
	ImageFormat        Format
	ImageColorSpace    ColorSpaceKHR
	ImageExtent        Extent2D
	ImageArrayLayers   uint32
	ImageUsage         ImageUsageFlags
	ImageSharingMode   SharingMode
	QueueFamilyIndices []uint32
	PreTransform       SurfaceTransformFlagsKHR
	CompositeAlpha     CompositeAlphaFlagsKHR
	PresentMode        PresentModeKHR
	Clipped            bool
	OldSwapchain       *SwapchainKHR
}

type SwapchainKHR struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSwapchainKHR)
	hnd C.VkSwapchainKHR
	dev *Device
}

func (dev *Device) CreateSwapchainKHR(info *SwapchainCreateInfoKHR) (SwapchainKHR, error) {
	// TODO(dh): support custom allocator
	ptr := (*C.VkSwapchainCreateInfoKHR)(C.calloc(1, C.sizeof_VkSwapchainCreateInfoKHR))
	defer C.free(unsafe.Pointer(ptr))
	ptr.sType = C.VkStructureType(StructureTypeSwapchainCreateInfoKHR)
	ptr.pNext = info.Next
	ptr.surface = info.Surface.hnd
	ptr.minImageCount = C.uint32_t(info.MinImageCount)
	ptr.imageFormat = C.VkFormat(info.ImageFormat)
	ptr.imageColorSpace = C.VkColorSpaceKHR(info.ImageColorSpace)
	ptr.imageExtent = C.VkExtent2D{
		width:  C.uint32_t(info.ImageExtent.Width),
		height: C.uint32_t(info.ImageExtent.Height),
	}
	ptr.imageArrayLayers = C.uint32_t(info.ImageArrayLayers)
	ptr.imageUsage = C.VkImageUsageFlags(info.ImageUsage)
	ptr.imageSharingMode = C.VkSharingMode(info.ImageSharingMode)
	if len(info.QueueFamilyIndices) > 0 {
		ptr.queueFamilyIndexCount = C.uint32_t(len(info.QueueFamilyIndices))
		ptr.pQueueFamilyIndices = (*C.uint32_t)(&info.QueueFamilyIndices[0])
	}
	ptr.preTransform = C.VkSurfaceTransformFlagBitsKHR(info.PreTransform)
	ptr.compositeAlpha = C.VkCompositeAlphaFlagBitsKHR(info.CompositeAlpha)
	ptr.presentMode = C.VkPresentModeKHR(info.PresentMode)
	ptr.clipped = vkBool(info.Clipped)
	if info.OldSwapchain != nil {
		ptr.oldSwapchain = info.OldSwapchain.hnd
	}

	var out C.VkSwapchainKHR
	res := Result(C.domVkCreateSwapchainKHR(dev.fps[vkCreateSwapchainKHR], dev.hnd, ptr, nil, &out))
	if res != Success {
		return SwapchainKHR{}, res
	}
	return SwapchainKHR{hnd: out, dev: dev}, nil
}

func (chain SwapchainKHR) Images() ([]Image, error) {
	var count C.uint32_t
	res := Result(C.domVkGetSwapchainImagesKHR(chain.dev.fps[vkGetSwapchainImagesKHR], chain.dev.hnd, chain.hnd, &count, nil))
	if res != Success {
		return nil, res
	}
	images := make([]C.VkImage, count)
	res = Result(C.domVkGetSwapchainImagesKHR(chain.dev.fps[vkGetSwapchainImagesKHR], chain.dev.hnd, chain.hnd, &count, &images[0]))
	if res != Success {
		return nil, res
	}
	out := make([]Image, count)
	for i, img := range images {
		out[i] = Image{hnd: img, dev: chain.dev}
	}
	return out, nil
}
