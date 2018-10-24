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
// VkResult domVkAcquireNextImageKHR(PFN_vkAcquireNextImageKHR fp, VkDevice device, VkSwapchainKHR swapchain, uint64_t timeout, VkSemaphore semaphore, VkFence fence, uint32_t* pImageIndex);
// VkResult domVkQueuePresentKHR(PFN_vkQueuePresentKHR fp, VkQueue queue, const VkPresentInfoKHR* pPresentInfo);
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

func (chain SwapchainKHR) AcquireNextImage(timeout time.Duration, semaphore *Semaphore) (uint32, error) {
	// TODO(dh): support fences
	var idx C.uint32_t
	var sem C.VkSemaphore
	if semaphore != nil {
		sem = semaphore.hnd
	}
	res := Result(C.domVkAcquireNextImageKHR(chain.dev.fps[vkAcquireNextImageKHR], chain.dev.hnd, chain.hnd, C.uint64_t(timeout), sem, nil, &idx))
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
	size0 := uintptr(C.sizeof_VkPresentInfoKHR)
	size1 := C.sizeof_VkSemaphore * uintptr(len(info.WaitSemaphores))
	size2 := C.sizeof_VkSwapchainKHR * uintptr(len(info.Swapchains))
	size3 := C.sizeof_uint32_t * uintptr(len(info.ImageIndices))
	size4 := C.sizeof_VkResult * uintptr(len(info.Swapchains))
	size := size0 + size1 + size2 + size3 + size4
	alloc := C.calloc(1, C.size_t(size))
	defer C.free(alloc)
	cinfo := (*C.VkPresentInfoKHR)(alloc)
	*cinfo = C.VkPresentInfoKHR{
		sType:              C.VkStructureType(StructureTypePresentInfoKHR),
		pNext:              info.Next,
		waitSemaphoreCount: C.uint32_t(len(info.WaitSemaphores)),
		pWaitSemaphores:    (*C.VkSemaphore)(unsafe.Pointer(uintptr(alloc) + size0)),
		swapchainCount:     C.uint32_t(len(info.Swapchains)),
		pSwapchains:        (*C.VkSwapchainKHR)(unsafe.Pointer(uintptr(alloc) + size0 + size1)),
		pImageIndices:      (*C.uint32_t)(unsafe.Pointer(uintptr(alloc) + size0 + size1 + size2)),
	}
	if len(results) != 0 {
		cinfo.pResults = (*C.VkResult)(unsafe.Pointer(uintptr(alloc) + size3))
	}
	ucopy(unsafe.Pointer(cinfo.pWaitSemaphores), unsafe.Pointer(&info.WaitSemaphores), C.sizeof_VkSemaphore)
	ucopy(unsafe.Pointer(cinfo.pImageIndices), unsafe.Pointer(&info.ImageIndices), C.sizeof_uint32_t)
	arr := (*[math.MaxInt32]C.VkSwapchainKHR)(unsafe.Pointer(cinfo.pSwapchains))[:len(info.Swapchains)]
	for i := range arr {
		arr[i] = info.Swapchains[i].hnd
	}

	res := Result(C.domVkQueuePresentKHR(queue.fps[vkQueuePresentKHR], queue.hnd, cinfo))
	if len(results) != 0 {
		copy(results, (*[math.MaxInt32]Result)(unsafe.Pointer(cinfo.pResults))[:len(info.Swapchains)])
	}
	if res != Success {
		return res
	}
	return nil
}
