// Copyright (c) 2015-2018 Khronos Group
// Copyright (c) 2018 Dominik Honnef

package vk

// #cgo noescape domVkQueuePresentKHR
// #cgo nocallback domVkQueuePresentKHR
// #include <stdlib.h>
// #include "vk.h"
import "C"

import (
	"math"
	"time"
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
	var ptr C.VkSwapchainCreateInfoKHR
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
	res := Result(C.domVkCreateSwapchainKHR(dev.fps[vkCreateSwapchainKHR], dev.hnd, &ptr, nil, &out))
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
		out[i] = Image{hnd: img}
	}
	return out, nil
}

func (chain SwapchainKHR) AcquireNextImage(timeout time.Duration, semaphore *Semaphore, fence *Fence) (uint32, error) {
	var idx C.uint32_t
	var semaphoreHnd C.VkSemaphore
	var fenceHnd C.VkFence
	if semaphore != nil {
		semaphoreHnd = semaphore.hnd
	}
	if fence != nil {
		fenceHnd = fence.hnd
	}
	res := Result(C.domVkAcquireNextImageKHR(chain.dev.fps[vkAcquireNextImageKHR], chain.dev.hnd, chain.hnd, C.uint64_t(timeout), semaphoreHnd, fenceHnd, &idx))
	if res != Success {
		return uint32(idx), res
	}
	return uint32(idx), nil
}

type PresentInfoKHR struct {
	Next           unsafe.Pointer
	WaitSemaphores []Semaphore
	Swapchains     []SwapchainKHR
	ImageIndices   []uint32
}

func (queue *Queue) Present(info *PresentInfoKHR, results []Result) error {
	a := new(allocator)
	defer a.free()

	cinfo := C.VkPresentInfoKHR{
		sType:              C.VkStructureType(StructureTypePresentInfoKHR),
		pNext:              info.Next,
		waitSemaphoreCount: C.uint32_t(len(info.WaitSemaphores)),
		pWaitSemaphores:    pinSliceAsCastedPtr[*C.VkSemaphore](a, info.WaitSemaphores),
		swapchainCount:     C.uint32_t(len(info.Swapchains)),
		pSwapchains:        allocn[C.VkSwapchainKHR](a, len(info.Swapchains)),
		pImageIndices:      pinSliceAsCastedPtr[*C.uint32_t](a, info.ImageIndices),
	}
	if len(results) != 0 {
		cinfo.pResults = allocn[C.VkResult](a, len(info.Swapchains))
	}
	arr := (*[math.MaxInt32]C.VkSwapchainKHR)(unsafe.Pointer(cinfo.pSwapchains))[:len(info.Swapchains)]
	for i := range arr {
		arr[i] = info.Swapchains[i].hnd
	}

	res := Result(C.domVkQueuePresentKHR(queue.fps[vkQueuePresentKHR], queue.hnd, &cinfo))
	// OPT we can avoid this copy, but we have to ensure that len(results) == len(info.Swapchains)
	if len(results) != 0 {
		copy(results, (*[math.MaxInt32]Result)(unsafe.Pointer(cinfo.pResults))[:len(info.Swapchains)])
	}
	if res != Success {
		return res
	}
	return nil
}
